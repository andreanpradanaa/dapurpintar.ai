"use client";

import { useState } from "react";
import { motion } from "framer-motion";
import Image from "next/image";
import { Check, Clock, Flame, Users, Star, Bookmark, Share2, ChefHat } from "lucide-react";
import { useLanguage } from "@/components/providers/language-provider";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";
import { Badge } from "@/components/ui/badge";
import { CircularProgress } from "@/components/ui/progress";
import { RECIPES } from "@/lib/mock-data/recipes";
import { photoForRecipe } from "@/lib/photo";

export function ProductPreview() {
  const { t } = useLanguage();
  const recipe = RECIPES[0];
  const [checked, setChecked] = useState<Record<number, boolean>>({});

  return (
    <section className="relative py-32 border-t border-border">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="max-w-2xl mb-14">
          <p className="text-xs font-semibold uppercase tracking-[0.12em] text-text-muted mb-4">
            {t.preview.eyebrow}
          </p>
          <h2 className="font-display text-balance text-4xl sm:text-5xl lg:text-6xl leading-[1.05] tracking-[-0.025em] text-text-primary">
            {t.preview.title}
          </h2>
          <p className="mt-5 text-lg text-text-muted leading-[1.6]">
            {t.preview.subtitle}
          </p>
        </div>

        <div className="grid lg:grid-cols-5 gap-6 lg:gap-8">
          {/* Image + meta */}
          <motion.div
            initial={{ opacity: 0, y: 16 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            transition={{ duration: 0.6 }}
            className="lg:col-span-2"
          >
            <div className="rounded-2xl overflow-hidden border border-border bg-bg-card h-full flex flex-col">
              <div className="relative aspect-[4/5] overflow-hidden bg-bg-section">
                <Image
                  src={photoForRecipe(recipe.slug)}
                  alt={recipe.title}
                  fill
                  sizes="(min-width: 1024px) 40vw, 100vw"
                  className="object-cover photo-warm"
                />
                <div className="absolute top-4 left-4 right-4 flex items-center justify-between">
                  <Badge variant="default" className="bg-bg-card/90 backdrop-blur-sm border-border text-text-primary">
                    <Star className="h-3 w-3 fill-accent text-accent" />
                    {recipe.rating}
                  </Badge>
                  <div className="flex items-center gap-2">
                    <button
                      type="button"
                      className="flex h-8 w-8 items-center justify-center rounded-md bg-bg-card/90 backdrop-blur-sm text-text-primary hover:bg-bg-card border border-border transition-colors"
                      aria-label="Bookmark"
                    >
                      <Bookmark className="h-4 w-4" strokeWidth={1.75} />
                    </button>
                    <button
                      type="button"
                      className="flex h-8 w-8 items-center justify-center rounded-md bg-bg-card/90 backdrop-blur-sm text-text-primary hover:bg-bg-card border border-border transition-colors"
                      aria-label="Share"
                    >
                      <Share2 className="h-4 w-4" strokeWidth={1.75} />
                    </button>
                  </div>
                </div>
                <div className="absolute bottom-4 left-4 right-4">
                  <p className="text-[10px] text-text-primary/90 uppercase tracking-[0.12em] font-medium drop-shadow-sm">
                    {recipe.cuisine}
                  </p>
                  <h3 className="text-2xl font-serif font-medium text-text-primary mt-1 drop-shadow-sm">
                    {recipe.title}
                  </h3>
                </div>
              </div>
              <div className="p-5 grid grid-cols-4 gap-2 border-t border-border">
                <Meta icon={<Clock className="h-3.5 w-3.5" />} value={`${recipe.prepTime + recipe.cookTime}m`} label="Total" />
                <Meta icon={<Flame className="h-3.5 w-3.5" />} value={recipe.difficulty} label="Level" />
                <Meta icon={<Users className="h-3.5 w-3.5" />} value={`${recipe.servings}`} label="Servings" />
                <Meta icon={<ChefHat className="h-3.5 w-3.5" />} value={recipe.nutrition.calories} label="Cal" />
              </div>
            </div>
          </motion.div>

          {/* Tabs panel */}
          <motion.div
            initial={{ opacity: 0, y: 16 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            transition={{ duration: 0.6, delay: 0.1 }}
            className="lg:col-span-3"
          >
            <div className="rounded-2xl border border-border bg-bg-card p-5 lg:p-6 h-full">
              <Tabs defaultValue="ingredients">
                <TabsList>
                  <TabsTrigger value="ingredients">Ingredients</TabsTrigger>
                  <TabsTrigger value="steps">Steps</TabsTrigger>
                  <TabsTrigger value="nutrition">Nutrition</TabsTrigger>
                </TabsList>
                <TabsContent value="ingredients" className="mt-5">
                  <ul className="grid sm:grid-cols-2 gap-2">
                    {recipe.ingredients.slice(0, 8).map((ing, i) => (
                      <li key={i}>
                        <button
                          type="button"
                          onClick={() => setChecked((c) => ({ ...c, [i]: !c[i] }))}
                          className="group flex items-center gap-3 w-full p-2.5 rounded-lg border border-border bg-bg-card hover:bg-bg-section hover:border-border-strong transition-colors text-left"
                        >
                          <span
                            className={
                              "flex h-4 w-4 shrink-0 items-center justify-center rounded border transition-colors " +
                              (checked[i]
                                ? "bg-accent border-accent"
                                : "border-border-strong")
                            }
                          >
                            {checked[i] && <Check className="h-3 w-3 text-accent-fg" strokeWidth={3} />}
                          </span>
                          <span
                            className={
                              "text-sm flex-1 " +
                              (checked[i]
                                ? "line-through text-text-subtle"
                                : "text-text-primary")
                            }
                          >
                            {ing.name}
                          </span>
                          <span className="text-xs text-text-muted tabular-nums">{ing.amount}</span>
                        </button>
                      </li>
                    ))}
                  </ul>
                </TabsContent>
                <TabsContent value="steps" className="mt-5">
                  <ol className="space-y-3">
                    {recipe.steps.slice(0, 5).map((step) => (
                      <li key={step.order} className="flex gap-3">
                        <span className="flex h-6 w-6 shrink-0 items-center justify-center rounded-full bg-bg-section text-text-secondary text-xs font-semibold tabular-nums">
                          {step.order}
                        </span>
                        <p className="text-sm text-text-secondary leading-[1.6] pt-0.5">
                          {step.text}
                        </p>
                      </li>
                    ))}
                  </ol>
                </TabsContent>
                <TabsContent value="nutrition" className="mt-5">
                  <div className="grid grid-cols-2 sm:grid-cols-4 gap-4">
                    {[
                      { label: "Calories", value: recipe.nutrition.calories, unit: "kcal" },
                      { label: "Protein", value: recipe.nutrition.protein, unit: "g" },
                      { label: "Carbs", value: recipe.nutrition.carbs, unit: "g" },
                      { label: "Fat", value: recipe.nutrition.fat, unit: "g" },
                    ].map((n) => (
                      <div key={n.label} className="flex flex-col items-center gap-2 p-4 rounded-xl border border-border bg-bg-card">
                        <CircularProgress value={(n.value / 100) * 100} size={64} strokeWidth={5} />
                        <div className="text-center">
                          <div className="text-base font-semibold text-text-primary tabular-nums">
                            {n.value}
                            <span className="text-xs text-text-muted ml-0.5">{n.unit}</span>
                          </div>
                          <div className="text-[10px] uppercase tracking-[0.1em] text-text-subtle mt-0.5">
                            {n.label}
                          </div>
                        </div>
                      </div>
                    ))}
                  </div>
                </TabsContent>
              </Tabs>
            </div>
          </motion.div>
        </div>
      </div>
    </section>
  );
}

function Meta({
  icon,
  value,
  label,
}: {
  icon: React.ReactNode;
  value: string | number;
  label: string;
}) {
  return (
    <div className="flex flex-col items-center gap-1 py-1">
      <div className="text-text-muted">{icon}</div>
      <div className="text-sm font-semibold text-text-primary capitalize tabular-nums">{value}</div>
      <div className="text-[10px] uppercase tracking-[0.1em] text-text-subtle">{label}</div>
    </div>
  );
}
