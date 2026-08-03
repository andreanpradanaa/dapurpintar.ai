"use client";

import * as React from "react";
import { motion, AnimatePresence } from "framer-motion";
import { ChevronDown } from "lucide-react";
import { cn } from "@/lib/utils";

type AccordionContextValue = {
  openValue: string | null;
  setOpen: (v: string | null) => void;
};
const AccordionContext = React.createContext<AccordionContextValue | null>(null);

export function Accordion({
  defaultValue,
  collapsible = true,
  children,
  className,
}: {
  defaultValue?: string;
  collapsible?: boolean;
  children: React.ReactNode;
  className?: string;
}) {
  const [openValue, setOpenValue] = React.useState<string | null>(defaultValue ?? null);
  const setOpen = React.useCallback(
    (v: string | null) => {
      if (collapsible && openValue === v) setOpenValue(null);
      else setOpenValue(v);
    },
    [openValue, collapsible]
  );
  return (
    <AccordionContext.Provider value={{ openValue, setOpen }}>
      <div className={cn("divide-y divide-border", className)}>{children}</div>
    </AccordionContext.Provider>
  );
}

function useAccordion() {
  const ctx = React.useContext(AccordionContext);
  if (!ctx) throw new Error("Accordion components must be used within Accordion");
  return ctx;
}

export function AccordionItem({
  value,
  children,
  className,
}: {
  value: string;
  children: React.ReactNode;
  className?: string;
}) {
  const { openValue } = useAccordion();
  const open = openValue === value;
  return (
    <div className={cn("group", className)} data-open={open}>
      {children}
    </div>
  );
}

export function AccordionTrigger({
  value,
  children,
  className,
}: {
  value: string;
  children: React.ReactNode;
  className?: string;
}) {
  const { openValue, setOpen } = useAccordion();
  const open = openValue === value;
  return (
    <button
      type="button"
      onClick={() => setOpen(open ? null : value)}
      aria-expanded={open}
      className={cn(
        "flex w-full items-center justify-between gap-4 py-5 text-left",
        "text-base font-medium text-text-primary hover:text-accent transition-colors",
        "focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent rounded-md",
        className
      )}
    >
      <span>{children}</span>
      <ChevronDown
        className={cn(
          "h-4 w-4 shrink-0 text-text-muted transition-transform duration-240",
          open && "rotate-180 text-accent"
        )}
        strokeWidth={2}
      />
    </button>
  );
}

export function AccordionContent({
  value,
  children,
  className,
}: {
  value: string;
  children: React.ReactNode;
  className?: string;
}) {
  const { openValue } = useAccordion();
  const open = openValue === value;
  return (
    <AnimatePresence initial={false}>
      {open && (
        <motion.div
          initial={{ height: 0, opacity: 0 }}
          animate={{ height: "auto", opacity: 1 }}
          exit={{ height: 0, opacity: 0 }}
          transition={{ duration: 0.24, ease: [0.22, 1, 0.36, 1] }}
          className="overflow-hidden"
        >
          <div className={cn("pb-5 pr-8 text-[15px] text-text-muted leading-relaxed", className)}>
            {children}
          </div>
        </motion.div>
      )}
    </AnimatePresence>
  );
}
