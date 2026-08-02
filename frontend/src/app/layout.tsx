import type { Metadata } from "next";
import { AuthProvider } from "@/lib/auth";
import { ToastProvider } from "@/lib/toast";
import Shell from "@/components/shell";
import "./globals.css";

export const metadata: Metadata = { title: "DapurPintar AI", description: "Decide dinner with what you have." };

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>
        <ToastProvider>
          <AuthProvider>
            <Shell>{children}</Shell>
          </AuthProvider>
        </ToastProvider>
      </body>
    </html>
  );
}
