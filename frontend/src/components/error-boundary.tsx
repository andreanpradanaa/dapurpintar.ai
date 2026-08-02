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
        <div className="min-h-screen flex items-center justify-center bg-santan-050">
          <div className="text-center space-y-4 max-w-sm px-4">
            <h1 className="text-xl font-bold text-kuali-950">Something went wrong</h1>
            <p className="text-sm text-kuali-700">An unexpected error occurred. Please try refreshing the page.</p>
            <div className="flex gap-2 justify-center">
              <button onClick={() => this.setState({ hasError: false })} className="text-rempah-500 text-sm font-medium hover:underline">Try again</button>
              <Link href="/today" className="text-bambu-300 text-sm hover:text-kuali-700">Go to Today</Link>
            </div>
          </div>
        </div>
      );
    }
    return this.props.children;
  }
}
