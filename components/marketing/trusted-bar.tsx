"use client";

import { motion } from "framer-motion";
import { useLanguage } from "@/components/providers/language-provider";
import { AvatarGroup } from "@/components/ui/avatar";

const USERS = [
  { name: "Andini Pradipta" },
  { name: "Made Wirawan" },
  { name: "Sari Indah" },
  { name: "Reza Hidayat" },
  { name: "Putu Anjani" },
  { name: "Budi Santoso" },
];

export function TrustedBar() {
  const { t } = useLanguage();
  return (
    <section className="relative py-16 border-y border-border bg-bg-section">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <motion.div
          initial={{ opacity: 0, y: 8 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ duration: 0.6 }}
          className="flex flex-col sm:flex-row items-center justify-between gap-8"
        >
          <div>
            <p className="text-sm text-text-muted mb-1 font-serif italic">
              {t.trusted.title}
            </p>
            <p className="text-2xl sm:text-3xl font-serif font-medium tracking-tight text-text-primary">
              8,200+ cooks · 24k recipes
            </p>
          </div>
          <div className="flex items-center gap-8">
            <AvatarGroup avatars={USERS} max={5} size="md" />
          </div>
        </motion.div>
      </div>
    </section>
  );
}
