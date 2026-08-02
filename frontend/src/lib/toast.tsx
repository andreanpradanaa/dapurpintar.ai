"use client";
import { createContext, useCallback, useContext, useState, type ReactNode } from "react";

type ToastType = "success" | "error" | "info";
type Toast = { id: number; type: ToastType; message: string };

type ToastCtx = { toast: (type: ToastType, message: string) => void };

const Ctx = createContext<ToastCtx>({ toast: () => {} });
export const useToast = () => useContext(Ctx);

export function ToastProvider({ children }: { children: ReactNode }) {
  const [toasts, setToasts] = useState<Toast[]>([]);
  let nextId = 0;

  const toast = useCallback((type: ToastType, message: string) => {
    const id = Date.now() + nextId++;
    setToasts(prev => [...prev, { id, type, message }]);
    setTimeout(() => setToasts(prev => prev.filter(t => t.id !== id)), 3500);
  }, []);

  const colors: Record<ToastType, string> = {
    success: "bg-context-positive text-context-positive-dark border-context-positive-dark",
    error: "bg-feedback-error text-white-000 border-feedback-error",
    info: "bg-feedback-info text-white-000 border-feedback-info",
  };

  return (
    <Ctx.Provider value={{ toast }}>
      {children}
      <div className="fixed bottom-20 right-4 z-50 space-y-2 max-w-sm">
        {toasts.map(t => (
          <div key={t.id} className={`${colors[t.type]} border rounded-lg px-4 py-3 text-sm shadow-lg animate-slideIn`}>
            {t.message}
          </div>
        ))}
      </div>
    </Ctx.Provider>
  );
}
