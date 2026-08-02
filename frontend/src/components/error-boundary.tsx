"use client";
import { Component, type ReactNode } from "react";
import Link from "next/link";

interface Props { children: ReactNode; }
interface State { hasError: boolean; }

export class ErrorBoundary extends Component<Props, State> {
  state: State = { hasError: false };

  static getDerivedStateFromError() { return { hasError: true }; }

  render() {
    if (this.state.hasError) {
      return (
        <div className="min-h-screen flex items-center justify-center bg-paper-050">
          <div className="text-center space-y-4 max-w-sm px-4">
            <h1 className="text-xl font-bold text-ink-900">Something went wrong</h1>
            <p className="text-sm text-ink-700">An unexpected error occurred. Please try refreshing the page.</p>
            <div className="flex gap-2 justify-center">
              <button onClick={() => this.setState({ hasError: false })} className="text-action-primary text-sm font-medium hover:underline">Try again</button>
              <Link href="/today" className="text-steel-400 text-sm hover:text-ink-700">Go to Today</Link>
            </div>
          </div>
        </div>
      );
    }
    return this.props.children;
  }
}
