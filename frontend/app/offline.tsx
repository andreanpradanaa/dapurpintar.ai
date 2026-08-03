"use client";

import { useEffect, useState } from "react";
import { WifiOff, RefreshCw } from "lucide-react";

export default function OfflinePage() {
  const [online, setOnline] = useState(true);

  useEffect(() => {
    setOnline(navigator.onLine);
    const onOnline = () => setOnline(true);
    const onOffline = () => setOnline(false);
    window.addEventListener("online", onOnline);
    window.addEventListener("offline", onOffline);
    return () => {
      window.removeEventListener("online", onOnline);
      window.removeEventListener("offline", onOffline);
    };
  }, []);

  if (online) return null;

  return (
    <div className="min-h-[60vh] flex flex-col items-center justify-center px-6 text-center">
      <div className="inline-flex h-14 w-14 items-center justify-center rounded-2xl bg-warning/10 text-warning mb-5">
        <WifiOff className="h-6 w-6" strokeWidth={1.5} />
      </div>
      <h1 className="font-display text-2xl font-medium tracking-tight text-text-primary">
        You&apos;re offline
      </h1>
      <p className="mt-2 text-sm text-text-muted max-w-md leading-[1.6]">
        Some features may be unavailable. Check your connection and try again.
      </p>
      <button
        type="button"
        onClick={() => window.location.reload()}
        className="mt-6 inline-flex items-center gap-1.5 h-11 px-5 rounded-lg text-sm font-medium bg-accent text-accent-fg hover:bg-accent-hover transition-colors"
      >
        <RefreshCw className="h-4 w-4" />
        Retry
      </button>
    </div>
  );
}
