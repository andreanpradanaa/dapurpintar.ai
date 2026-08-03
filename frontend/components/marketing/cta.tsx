"use client";

import Link from "next/link";
import { motion } from "framer-motion";
import { ArrowRight } from "lucide-react";
import { Button } from "@/components/ui/button";
import { useLanguage } from "@/components/providers/language-provider";

export function CTA() {
  const { t } = useLanguage();
  return (
    <section className="relative py-32 border-t border-border overflow-hidden">
      <div className="absolute inset-0 -z-10">
        <div className="absolute inset-0 bg-warm-wash" />
      </div>
      <div className="mx-auto max-w-4xl px-4 sm:px-6 lg:px-8 text-center">
        <motion.h2
          initial={{ opacity: 0, y: 16 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ duration: 0.6 }}
          className="font-display text-balance text-4xl sm:text-6xl lg:text-7xl leading-[1.0] tracking-[-0.035em] text-text-primary"
        >
          {t.cta.title}
        </motion.h2>
        <motion.p
          initial={{ opacity: 0, y: 16 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ duration: 0.6, delay: 0.1 }}
          className="mt-6 text-lg text-text-muted max-w-xl mx-auto leading-[1.6]"
        >
          {t.cta.subtitle}
        </motion.p>
        <motion.div
          initial={{ opacity: 0, y: 16 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ duration: 0.6, delay: 0.2 }}
          className="mt-10 flex flex-col sm:flex-row items-center justify-center gap-3"
        >
          <Link href="/register">
            <Button size="lg" className="group">
              {t.cta.button}
              <ArrowRight className="h-4 w-4 transition-transform group-hover:translate-x-0.5" />
            </Button>
          </Link>
          <Link href="/pricing">
            <Button size="lg" variant="ghost">
              See pricing
            </Button>
          </Link>
        </motion.div>
      </div>
    </section>
  );
}
