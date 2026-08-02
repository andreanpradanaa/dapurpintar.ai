import type { Metadata } from "next";
import { AuthProvider } from "@/lib/auth";
import { ToastProvider } from "@/lib/toast";
import { NetworkProvider } from "@/lib/network";
import Shell from "@/components/shell";
import "./globals.css";

export const metadata: Metadata = { title: "DapurPintar AI", description: "Decide dinner with what you have." };

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>
        <NetworkProvider>
          <ToastProvider>
            <AuthProvider>
              <Shell>{children}</Shell>
            </AuthProvider>
          </ToastProvider>
        </NetworkProvider>
      </body>
    </html>
  );
}
