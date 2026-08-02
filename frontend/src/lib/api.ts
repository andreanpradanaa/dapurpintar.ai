const BASE = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api/v1";

type RequestOptions = { method?: string; body?: unknown; headers?: Record<string, string> };

async function request<T>(path: string, opts: RequestOptions = {}): Promise<T> {
  const url = `${BASE}${path}`;
  const headers: Record<string, string> = { "Content-Type": "application/json", ...opts.headers };
  const res = await fetch(url, {
    method: opts.method || "GET",
    headers,
    body: opts.body ? JSON.stringify(opts.body) : undefined,
    credentials: "include",
  });
  if (!res.ok) {
    const err = await res.json().catch(() => ({ error: { message: "Network error" } }));
    throw new Error(err?.error?.message || `Request failed (${res.status})`);
  }
  if (res.status === 204) return undefined as T;
  return res.json() as Promise<T>;
}

export type Account = { id: string; email: string; status: string; timezone: string; email_verified: boolean; profile_id?: string; created_at: string };
export type Profile = { id: string; display_name: string; status: string; timezone?: string };
export type PantryItem = { id: string; ingredient_name: string; category: string; quantity: number; unit: string; expiry_date?: string; status: string };
export type PantrySummary = { total_items: number; expiring_soon_count: number; running_low_count: number };
export type Page<T> = { next_cursor?: string; has_more: boolean };
export type Collection<T> = { data: T[]; page: Page<T> };

export const api = {
  // Auth
  register: (body: { email: string; password: string; display_name: string; timezone: string }) =>
    request<{ data: Account }>("/accounts", { method: "POST", body }),
  login: (body: { email: string; password: string }) =>
    request<{ data: Account }>("/accounts/login", { method: "POST", body }),
  refresh: () => request<{ data: Account }>("/accounts/refresh", { method: "POST" }),
  logout: () => request<void>("/accounts/logout", { method: "POST" }),
  me: () => request<{ data: Account }>("/accounts/me"),

  // Profile
  getProfile: () => request<{ data: Profile }>("/profile"),
  updateProfile: (body: { display_name?: string; timezone?: string }) =>
    request<{ data: Profile }>("/profile", { method: "PATCH", body }),

  // Pantry
  pantrySummary: () => request<{ data: PantrySummary }>("/pantry"),
  pantryItems: (cursor?: string) =>
    request<Collection<PantryItem>>(`/pantry/items?cursor=${cursor || ""}&limit=50`),
  addPantryItem: (body: { ingredient_name: string; category: string; quantity: number; unit?: string; expiry_date?: string }) =>
    request<{ data: PantryItem }>("/pantry/items", { method: "POST", body }),
  expiringItems: (cursor?: string) =>
    request<Collection<PantryItem>>(`/pantry/expiry?cursor=${cursor || ""}&limit=20`),
};

export async function getAuthenticated<T>(fn: () => Promise<T>): Promise<T> {
  try { return await fn(); } catch {
    try { await api.refresh(); return await fn(); } catch { throw new Error("Authentication required"); }
  }
}
