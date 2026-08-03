"use client";

import { motion } from "framer-motion";
import { ChefHat, Heart, Award, Mail, Calendar, Edit2, Camera } from "lucide-react";
import Image from "next/image";
import Link from "next/link";
import { Avatar } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { useLanguage } from "@/components/providers/language-provider";
import { useStore } from "@/lib/store";
import { RECIPES } from "@/lib/mock-data/recipes";
import { photoForRecipe } from "@/lib/photo";
import { formatDate, relativeTime } from "@/lib/utils";

const STATS = [
  { icon: ChefHat, label: "Recipes created", key: "recipesGenerated" as const, color: "text-accent" },
  { icon: Heart, label: "Favorites", key: "favoritesCount" as const, color: "text-accent-active" },
  { icon: Award, label: "Day streak", key: "streak" as const, color: "text-success" },
];

export default function ProfilePage() {
  const { t } = useLanguage();
  const user = useStore((s) => s.user);
  const history = useStore((s) => s.history);
  const favorites = useStore((s) => s.favorites);

  const recent = history
    .slice(0, 5)
    .map((h) => ({ ...h, recipe: RECIPES.find((r) => r.id === h.recipeId) }))
    .filter((h) => h.recipe);

  return (
    <div className="space-y-8 max-w-4xl">
      <motion.div
        initial={{ opacity: 0, y: 12 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.6 }}
        className="relative rounded-2xl border border-border bg-bg-card p-6 sm:p-8 overflow-hidden"
      >
        <div className="relative flex flex-col sm:flex-row items-start sm:items-center gap-6">
          <div className="relative shrink-0">
            <Avatar name={user?.name ?? "User"} size="xl" />
            <button
              type="button"
              className="absolute -bottom-1 -right-1 flex h-8 w-8 items-center justify-center rounded-full bg-accent text-accent-fg border-2 border-bg-card hover:bg-accent-hover"
              aria-label="Change photo"
            >
              <Camera className="h-3.5 w-3.5" />
            </button>
          </div>
          <div className="flex-1 min-w-0">
            <div className="flex items-center gap-2 mb-1">
              <h1 className="font-display text-2xl sm:text-3xl font-medium tracking-tight text-text-primary">
                {user?.name}
              </h1>
              <Badge variant="accent" className="capitalize">{user?.plan}</Badge>
            </div>
            <div className="flex flex-wrap items-center gap-x-4 gap-y-1 text-sm text-text-muted">
              <span className="flex items-center gap-1.5">
                <Mail className="h-3.5 w-3.5" />
                {user?.email}
              </span>
              <span className="flex items-center gap-1.5">
                <Calendar className="h-3.5 w-3.5" />
                {t.app.profile.member} {formatDate(user?.joinedAt ?? new Date(), { month: "long", year: "numeric" })}
              </span>
            </div>
            {user?.bio && (
              <p className="mt-3 text-sm text-text-secondary max-w-xl leading-[1.6]">{user.bio}</p>
            )}
          </div>
          <Button variant="outline" size="sm">
            <Edit2 className="h-3.5 w-3.5" />
            Edit profile
          </Button>
        </div>
      </motion.div>

      <div className="grid grid-cols-3 gap-4">
        {STATS.map((s, i) => {
          const value = user?.[s.key] ?? 0;
          return (
            <motion.div
              key={s.key}
              initial={{ opacity: 0, y: 8 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: i * 0.05 }}
              className="rounded-2xl border border-border bg-bg-card p-5 hover:border-border-strong transition-colors"
            >
              <div className="flex items-center gap-2 mb-2">
                <div className={`flex h-8 w-8 items-center justify-center rounded-md bg-bg-section ${s.color}`}>
                  <s.icon className="h-4 w-4" strokeWidth={1.5} />
                </div>
              </div>
              <div className="font-serif text-2xl font-medium text-text-primary tabular-nums">
                {value}
              </div>
              <div className="text-xs text-text-muted mt-0.5">{s.label}</div>
            </motion.div>
          );
        })}
      </div>

      <div>
        <h2 className="text-base font-serif font-medium text-text-primary mb-4">
          {t.app.profile.activity}
        </h2>
        {recent.length === 0 ? (
          <div className="rounded-2xl border border-dashed border-border bg-bg-card/40 p-10 text-center text-sm text-text-muted">
            No activity yet.
          </div>
        ) : (
          <div className="rounded-2xl border border-border bg-bg-card divide-y divide-border">
            {recent.map((h) => (
              <Link
                key={h.id}
                href={`/recipes/${h.recipe!.slug}`}
                className="group flex items-center gap-4 p-3 hover:bg-bg-section transition-colors"
              >
                <div className="relative h-12 w-12 shrink-0 rounded-lg overflow-hidden bg-bg-section">
                  <Image
                    src={photoForRecipe(h.recipe!.slug)}
                    alt={h.recipe!.title}
                    fill
                    sizes="48px"
                    className="object-cover photo-warm"
                  />
                </div>
                <div className="flex-1 min-w-0">
                  <div className="text-sm font-medium text-text-primary truncate font-serif">
                    {h.recipe!.title}
                  </div>
                  <div className="text-xs text-text-muted truncate">
                    {h.ingredients.slice(0, 4).join(" · ")}
                  </div>
                </div>
                <div className="text-xs text-text-muted tabular-nums shrink-0">
                  {relativeTime(h.createdAt)}
                </div>
              </Link>
            ))}
          </div>
        )}
      </div>

      {favorites.length > 0 && (
        <div>
          <h2 className="text-base font-serif font-medium text-text-primary mb-4">Favorites</h2>
          <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-3">
            {favorites.slice(0, 4).map((id) => {
              const r = RECIPES.find((x) => x.id === id);
              if (!r) return null;
              return (
                <Link
                  key={id}
                  href={`/recipes/${r.slug}`}
                  className="group rounded-xl overflow-hidden border border-border bg-bg-card hover:border-border-strong transition-all"
                >
                  <div className="relative aspect-square bg-bg-section">
                    <Image
                      src={photoForRecipe(r.slug)}
                      alt={r.title}
                      fill
                      sizes="(min-width: 1024px) 25vw, 50vw"
                      className="object-cover photo-warm"
                    />
                  </div>
                  <div className="p-3">
                    <div className="text-sm font-medium text-text-primary truncate font-serif">
                      {r.title}
                    </div>
                    <div className="text-xs text-text-muted">{r.cuisine}</div>
                  </div>
                </Link>
              );
            })}
          </div>
        </div>
      )}
    </div>
  );
}
