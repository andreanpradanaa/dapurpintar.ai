"use client";

import { useEffect } from "react";
import { AlertTriangle, RefreshCw, Home } from "lucide-react";
import Link from "next/link";

export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  useEffect(() => {
    console.error(error);
  }, [error]);

  return (
    <div className="min-h-[60vh] flex flex-col items-center justify-center px-6">
      <div className="inline-flex h-14 w-14 items-center justify-center rounded-2xl bg-danger/10 text-danger mb-5">
        <AlertTriangle className="h-6 w-6" strokeWidth={1.5} />
      </div>
      <h1 className="font-display text-2xl sm:text-3xl font-medium tracking-tight text-text-primary">
        Something went wrong
      </h1>
      <p className="mt-2 text-sm text-text-muted max-w-md text-center leading-[1.6]">
        We hit a snag rendering this page. Please try again.
      </p>
      {error.digest && (
        <p className="mt-2 text-xs text-text-subtle font-mono">Error ID: {error.digest}</p>
      )}
      <div className="mt-6 flex gap-2">
        <button
          type="button"
          onClick={reset}
          className="inline-flex items-center gap-1.5 h-11 px-5 rounded-lg text-sm font-medium bg-accent text-accent-fg hover:bg-accent-hover transition-colors"
        >
          <RefreshCw className="h-4 w-4" />
          Try again
        </button>
        <Link
          href="/dashboard"
          className="inline-flex items-center gap-1.5 h-11 px-5 rounded-lg text-sm font-medium text-text-secondary hover:text-text-primary hover:bg-bg-section transition-colors"
        >
          <Home className="h-4 w-4" />
          Dashboard
        </Link>
      </div>
    </div>
  );
}
