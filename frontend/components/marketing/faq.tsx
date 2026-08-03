"use client";

import { motion } from "framer-motion";
import { useLanguage } from "@/components/providers/language-provider";
import { Accordion, AccordionItem, AccordionTrigger, AccordionContent } from "@/components/ui/accordion";

export function FAQ() {
  const { t } = useLanguage();
  return (
    <section id="faq" className="relative py-32 border-t border-border">
      <div className="mx-auto max-w-3xl px-4 sm:px-6 lg:px-8">
        <motion.div
          initial={{ opacity: 0, y: 12 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ duration: 0.6 }}
          className="mb-12"
        >
          <p className="text-xs font-semibold uppercase tracking-[0.12em] text-text-muted mb-4">
            {t.faq.eyebrow}
          </p>
          <h2 className="font-display text-balance text-4xl sm:text-5xl lg:text-6xl leading-[1.05] tracking-[-0.025em] text-text-primary">
            {t.faq.title}
          </h2>
          <p className="mt-5 text-lg text-text-muted">{t.faq.subtitle}</p>
        </motion.div>

        <motion.div
          initial={{ opacity: 0, y: 12 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ duration: 0.6, delay: 0.1 }}
          className="rounded-2xl border border-border bg-bg-card px-2 sm:px-4"
        >
          <Accordion defaultValue="0">
            {t.faq.items.map((item, i) => (
              <AccordionItem key={i} value={String(i)}>
                <AccordionTrigger value={String(i)}>{item.q}</AccordionTrigger>
                <AccordionContent value={String(i)}>{item.a}</AccordionContent>
              </AccordionItem>
            ))}
          </Accordion>
        </motion.div>
      </div>
    </section>
  );
}
