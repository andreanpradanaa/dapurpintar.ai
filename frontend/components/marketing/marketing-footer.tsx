"use client";

import Link from "next/link";
import { Logo } from "@/components/ui/logo";
import { LanguageToggle } from "@/components/ui/language-toggle";
import { useLanguage } from "@/components/providers/language-provider";
import { translations } from "@/lib/i18n";

type FooterColumn = {
  titleKey: "product" | "company" | "resources" | "legal";
  links: { labelKey: keyof typeof translations.en.footer.links; href: string }[];
};

const COLUMNS: FooterColumn[] = [
  {
    titleKey: "product",
    links: [
      { labelKey: "features", href: "/#features" },
      { labelKey: "pricing", href: "/pricing" },
      { labelKey: "generator", href: "/generate" },
      { labelKey: "favorites", href: "/favorites" },
    ],
  },
  {
    titleKey: "company",
    links: [
      { labelKey: "about", href: "#" },
      { labelKey: "blog", href: "#" },
      { labelKey: "careers", href: "#" },
      { labelKey: "contact", href: "#" },
    ],
  },
  {
    titleKey: "resources",
    links: [
      { labelKey: "docs", href: "/faq" },
      { labelKey: "community", href: "#" },
      { labelKey: "changelog", href: "#" },
    ],
  },
  {
    titleKey: "legal",
    links: [
      { labelKey: "privacy", href: "#" },
      { labelKey: "terms", href: "#" },
      { labelKey: "cookies", href: "#" },
    ],
  },
];

export function MarketingFooter() {
  const { t } = useLanguage();
  return (
    <footer className="border-t border-border bg-bg-section">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 py-20">
        <div className="grid grid-cols-2 md:grid-cols-6 gap-10 mb-14">
          <div className="col-span-2">
            <Logo size="md" />
            <p className="mt-5 text-sm text-text-muted leading-relaxed max-w-xs">
              {t.footer.tagline}
            </p>
            <div className="mt-6">
              <LanguageToggle />
            </div>
          </div>
          {COLUMNS.map((col) => (
            <div key={col.titleKey}>
              <h3 className="text-[11px] font-semibold uppercase tracking-[0.12em] text-text-subtle mb-4">
                {t.footer[col.titleKey]}
              </h3>
              <ul className="space-y-3">
                {col.links.map((link) => (
                  <li key={link.labelKey}>
                    <Link
                      href={link.href}
                      className="text-sm text-text-muted hover:text-text-primary transition-colors duration-200"
                    >
                      {t.footer.links[link.labelKey]}
                    </Link>
                  </li>
                ))}
              </ul>
            </div>
          ))}
        </div>
        <div className="pt-8 border-t border-border flex flex-col sm:flex-row items-center justify-between gap-4">
          <p className="text-xs text-text-subtle">{t.footer.copyright}</p>
          <div className="flex items-center gap-2 text-xs text-text-subtle">
            <span className="inline-block h-1.5 w-1.5 rounded-full bg-success pulse-soft" />
            <span>All systems operational</span>
          </div>
        </div>
      </div>
    </footer>
  );
}
