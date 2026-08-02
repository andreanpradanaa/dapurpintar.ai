"use client";
import { createContext, useCallback, useContext, useEffect, useState, type ReactNode } from "react";
import { api, getAuthenticated, type Account, type Profile } from "./api";

type AuthState = {
  account: Account | null;
  profile: Profile | null;
  loading: boolean;
  login: (email: string, password: string) => Promise<void>;
  register: (email: string, password: string, name: string) => Promise<void>;
  logout: () => Promise<void>;
  refresh: () => Promise<void>;
};

const AuthCtx = createContext<AuthState>({ account: null, profile: null, loading: true } as AuthState);
export const useAuth = () => useContext(AuthCtx);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [account, setAccount] = useState<Account | null>(null);
  const [profile, setProfile] = useState<Profile | null>(null);
  const [loading, setLoading] = useState(true);

  const refresh = useCallback(async () => {
    try {
      const a = await getAuthenticated(() => api.me());
      setAccount(a.data);
      try {
        const p = await getAuthenticated(() => api.getProfile());
        setProfile(p.data);
      } catch { /* profile may not exist yet */ }
    } catch {
      setAccount(null);
      setProfile(null);
    }
  }, []);

  useEffect(() => { refresh().finally(() => setLoading(false)); }, [refresh]);

  const login = async (email: string, password: string) => {
    await api.login({ email, password });
    await refresh();
  };
  const register = async (email: string, password: string, name: string) => {
    await api.register({ email, password, display_name: name, timezone: "Asia/Jakarta" });
    await refresh();
  };
  const logout = async () => {
    await api.logout();
    setAccount(null);
    setProfile(null);
  };

  return (
    <AuthCtx.Provider value={{ account, profile, loading, login, register, logout, refresh }}>
      {children}
    </AuthCtx.Provider>
  );
}
