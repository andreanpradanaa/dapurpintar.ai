export const SITE_CONFIG = {
  name: "Dapur Pintar AI",
  shortName: "Dapur Pintar",
  tagline: "Turn Ingredients Into Delicious Meals with AI.",
  taglineId: "Ubah Bahan Masakan Menjadi Hidangan Lezat dengan AI.",
  description:
    "AI-powered recipe generator that turns the ingredients in your kitchen into delicious meals. Reduce food waste, save time, eat better.",
  url: "https://dapurpintar.ai",
  ogImage: "/og.svg",
  twitter: "@dapurpintarai",
  author: "Dapur Pintar AI",
  email: "hello@dapurpintar.ai",
} as const;

export const NAV_LINKS = [
  { href: "/#features", labelKey: "nav.features" as const },
  { href: "/pricing", labelKey: "nav.pricing" as const },
  { href: "/faq", labelKey: "nav.faq" as const },
];

export const APP_NAV = [
  { href: "/dashboard", labelKey: "app.sidebar.dashboard" as const, icon: "LayoutDashboard" },
  { href: "/generate", labelKey: "app.sidebar.generate" as const, icon: "Sparkles" },
  { href: "/history", labelKey: "app.sidebar.history" as const, icon: "History" },
  { href: "/favorites", labelKey: "app.sidebar.favorites" as const, icon: "Heart" },
  { href: "/settings", labelKey: "app.sidebar.settings" as const, icon: "Settings" },
  { href: "/profile", labelKey: "app.sidebar.profile" as const, icon: "User" },
] as const;
