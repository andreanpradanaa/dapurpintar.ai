"use client";
import { createContext, useCallback, useContext, useEffect, useState, type ReactNode } from "react";

type NetworkCtx = { online: boolean };

const Ctx = createContext<NetworkCtx>({ online: true });
export const useNetwork = () => useContext(Ctx);

export function NetworkProvider({ children }: { children: ReactNode }) {
  const [online, setOnline] = useState(true);

  useEffect(() => {
    setOnline(navigator.onLine);
    const go = () => setOnline(true);
    const off = () => setOnline(false);
    window.addEventListener("online", go);
    window.addEventListener("offline", off);
    return () => { window.removeEventListener("online", go); window.removeEventListener("offline", off); };
  }, []);

  return (
    <Ctx.Provider value={{ online }}>
      {!online && (
        <div className="fixed top-0 left-0 right-0 bg-context-attention text-context-attention-dark text-sm text-center py-2 font-medium z-50" role="alert">
          You are offline. Changes may not be saved.
        </div>
      )}
      {children}
    </Ctx.Provider>
  );
}
