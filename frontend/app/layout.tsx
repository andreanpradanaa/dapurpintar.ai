import type { Metadata, Viewport } from "next";
import { Inter, Fraunces, JetBrains_Mono } from "next/font/google";
import "./globals.css";
import { Toaster } from "sonner";
import { LanguageProvider } from "@/components/providers/language-provider";

const inter = Inter({
  subsets: ["latin"],
  variable: "--font-inter",
  display: "swap",
  weight: ["300", "400", "500", "600", "700"],
});

const fraunces = Fraunces({
  subsets: ["latin"],
  variable: "--font-fraunces",
  display: "swap",
  axes: ["opsz", "SOFT"],
});

const jetbrains = JetBrains_Mono({
  subsets: ["latin"],
  variable: "--font-jetbrains",
  display: "swap",
  weight: ["400", "500", "600"],
});

export const metadata: Metadata = {
  metadataBase: new URL("https://dapurpintar.ai"),
  title: {
    default: "Dapur Pintar — A quiet cooking companion",
    template: "%s | Dapur Pintar",
  },
  description:
    "Tell us what's in your kitchen. We'll help you decide what to make.",
  keywords: [
    "cooking",
    "recipe",
    "kitchen",
    "ingredient",
    "meal",
    "Indonesian",
  ],
  authors: [{ name: "Dapur Pintar" }],
  creator: "Dapur Pintar",
  openGraph: {
    type: "website",
    locale: "en_US",
    url: "https://dapurpintar.ai",
    title: "Dapur Pintar — A quiet cooking companion",
    description: "Tell us what's in your kitchen. We'll help you decide what to make.",
    siteName: "Dapur Pintar",
    images: [
      {
        url: "/og.svg",
        width: 1200,
        height: 630,
        alt: "Dapur Pintar",
      },
    ],
  },
  twitter: {
    card: "summary_large_image",
    title: "Dapur Pintar",
    description: "Tell us what's in your kitchen. We'll help you decide what to make.",
    images: ["/og.svg"],
  },
  icons: {
    icon: [{ url: "/favicon.svg", type: "image/svg+xml" }],
    apple: "/apple-touch-icon.svg",
  },
  robots: {
    index: true,
    follow: true,
  },
};

export const viewport: Viewport = {
  themeColor: "#FAF6F0",
  width: "device-width",
  initialScale: 1,
  maximumScale: 5,
  colorScheme: "light",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html
      lang="en"
      className={`${inter.variable} ${fraunces.variable} ${jetbrains.variable}`}
      suppressHydrationWarning
    >
      <body className="min-h-screen bg-bg-base text-text-primary antialiased">
        <LanguageProvider>
          {children}
          <Toaster
            position="bottom-right"
            theme="light"
            toastOptions={{
              style: {
                background: "#FFFFFF",
                border: "1px solid rgba(26, 22, 18, 0.08)",
                color: "#1A1612",
                boxShadow: "0 12px 32px rgba(26, 22, 18, 0.10)",
              },
            }}
          />
        </LanguageProvider>
      </body>
    </html>
  );
}
