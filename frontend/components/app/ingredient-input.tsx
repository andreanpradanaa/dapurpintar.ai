"use client";

import { useState, useRef } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { Sparkles } from "lucide-react";
import { Chip } from "@/components/ui/chip";
import { INGREDIENTS } from "@/lib/mock-data/ingredients";
import { useLanguage } from "@/components/providers/language-provider";
import { cn } from "@/lib/utils";

export function IngredientInput({
  value,
  onChange,
  className,
}: {
  value: string[];
  onChange: (next: string[]) => void;
  className?: string;
}) {
  const { t, lang } = useLanguage();
  const [input, setInput] = useState("");
  const [showSuggestions, setShowSuggestions] = useState(false);
  const inputRef = useRef<HTMLInputElement>(null);

  const filteredSuggestions = INGREDIENTS.filter(
    (ing) =>
      !value.includes(ing.name) &&
      (ing.name.toLowerCase().includes(input.toLowerCase()) ||
        ing.nameId.toLowerCase().includes(input.toLowerCase()) ||
        input === "")
  ).slice(0, 8);

  const add = (name: string) => {
    if (!name.trim() || value.includes(name.trim())) return;
    onChange([...value, name.trim()]);
    setInput("");
    setShowSuggestions(true);
    inputRef.current?.focus();
  };

  const remove = (name: string) => {
    onChange(value.filter((v) => v !== name));
  };

  const onKey = (e: React.KeyboardEvent) => {
    if (e.key === "Enter" || e.key === ",") {
      e.preventDefault();
      add(input);
    } else if (e.key === "Backspace" && !input && value.length > 0) {
      remove(value[value.length - 1]);
    }
  };

  return (
    <div className={cn("space-y-3", className)}>
      <div
        className={cn(
          "flex flex-wrap items-center gap-2 p-3 rounded-xl border border-border bg-bg-card transition-colors",
          "focus-within:border-border-strong"
        )}
        onClick={() => inputRef.current?.focus()}
      >
        <AnimatePresence>
          {value.map((v) => (
            <Chip
              key={v}
              variant="accent"
              onRemove={() => remove(v)}
              icon={<Sparkles className="h-3 w-3" />}
            >
              {v}
            </Chip>
          ))}
        </AnimatePresence>
        <input
          ref={inputRef}
          type="text"
          value={input}
          onChange={(e) => {
            setInput(e.target.value);
            setShowSuggestions(true);
          }}
          onFocus={() => setShowSuggestions(true)}
          onBlur={() => setTimeout(() => setShowSuggestions(false), 150)}
          onKeyDown={onKey}
          placeholder={value.length === 0 ? t.app.generate.placeholder : ""}
          className="flex-1 min-w-[120px] bg-transparent px-1.5 py-1 text-sm text-text-primary placeholder:text-text-subtle focus:outline-none"
        />
      </div>

      <AnimatePresence>
        {showSuggestions && filteredSuggestions.length > 0 && (
          <motion.div
            initial={{ opacity: 0, y: -4 }}
            animate={{ opacity: 1, y: 0 }}
            exit={{ opacity: 0, y: -4 }}
            className="rounded-xl border border-border bg-bg-card p-2"
          >
            <div className="text-[10px] uppercase tracking-[0.1em] text-text-subtle px-2 py-1 mb-1">
              Suggestions
            </div>
            <div className="flex flex-wrap gap-1">
              {filteredSuggestions.map((s) => (
                <button
                  key={s.name}
                  type="button"
                  onClick={() => add(s.name)}
                  className="inline-flex items-center gap-1.5 text-xs text-text-secondary hover:text-text-primary px-2 py-1 rounded-md border border-transparent hover:border-border hover:bg-bg-section transition-colors"
                >
                  <span>{s.emoji}</span>
                  <span>{lang === "en" ? s.name : s.nameId}</span>
                </button>
              ))}
            </div>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  );
}
