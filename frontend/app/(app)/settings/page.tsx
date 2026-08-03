"use client";

import { useState } from "react";
import { motion } from "framer-motion";
import {
  User,
  Sliders,
  Bell,
  CreditCard,
  AlertTriangle,
  Check,
  Camera,
} from "lucide-react";
import { Input, Textarea } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";
import { Switch } from "@/components/ui/toggle";
import { Avatar } from "@/components/ui/avatar";
import { Badge } from "@/components/ui/badge";
import { useLanguage } from "@/components/providers/language-provider";
import { useStore } from "@/lib/store";
import { useLanguage as useLang } from "@/components/providers/language-provider";
import { toast } from "sonner";

export default function SettingsPage() {
  const { t } = useLanguage();
  const { lang, setLang } = useLang();
  const user = useStore((s) => s.user);
  const updateUser = useStore((s) => s.updateUser);
  const prefs = useStore((s) => s.preferences);
  const setPreference = useStore((s) => s.setPreference);
  const setNotification = useStore((s) => s.setNotification);

  const [name, setName] = useState(user?.name ?? "");
  const [email, setEmail] = useState(user?.email ?? "");
  const [bio, setBio] = useState(user?.bio ?? "");
  const [saved, setSaved] = useState(false);

  const onSave = () => {
    updateUser({ name, email, bio });
    setSaved(true);
    toast.success("Changes saved");
    setTimeout(() => setSaved(false), 2000);
  };

  return (
    <div className="space-y-6 max-w-4xl">
      <div>
        <h1 className="font-display text-2xl sm:text-3xl font-medium tracking-tight text-text-primary">
          {t.app.settings.title}
        </h1>
        <p className="mt-1.5 text-sm text-text-muted">{t.app.settings.subtitle}</p>
      </div>

      <Tabs defaultValue="account">
        <TabsList>
          <TabsTrigger value="account">
            <User className="h-3.5 w-3.5" />
            <span className="ml-1.5">{t.app.settings.tabs.account}</span>
          </TabsTrigger>
          <TabsTrigger value="preferences">
            <Sliders className="h-3.5 w-3.5" />
            <span className="ml-1.5">{t.app.settings.tabs.preferences}</span>
          </TabsTrigger>
          <TabsTrigger value="notifications">
            <Bell className="h-3.5 w-3.5" />
            <span className="ml-1.5">{t.app.settings.tabs.notifications}</span>
          </TabsTrigger>
          <TabsTrigger value="billing">
            <CreditCard className="h-3.5 w-3.5" />
            <span className="ml-1.5">{t.app.settings.tabs.billing}</span>
          </TabsTrigger>
          <TabsTrigger value="danger">
            <AlertTriangle className="h-3.5 w-3.5" />
            <span className="ml-1.5">{t.app.settings.tabs.danger}</span>
          </TabsTrigger>
        </TabsList>

        <TabsContent value="account">
          <motion.div
            initial={{ opacity: 0, y: 8 }}
            animate={{ opacity: 1, y: 0 }}
            className="rounded-2xl border border-border bg-bg-card p-6 space-y-6"
          >
            <div className="flex items-center gap-5 pb-6 border-b border-border">
              <div className="relative">
                <Avatar name={name || "User"} size="xl" />
                <button
                  type="button"
                  className="absolute -bottom-1 -right-1 flex h-7 w-7 items-center justify-center rounded-full bg-accent text-accent-fg border-2 border-bg-card hover:bg-accent-hover"
                  aria-label="Change photo"
                >
                  <Camera className="h-3 w-3" />
                </button>
              </div>
              <div>
                <h2 className="text-lg font-serif font-medium text-text-primary">{name || "—"}</h2>
                <p className="text-sm text-text-muted">{email}</p>
                <Badge variant="accent" className="mt-1.5 capitalize">
                  {user?.plan} plan
                </Badge>
              </div>
            </div>

            <div className="grid sm:grid-cols-2 gap-4">
              <Field label={t.app.settings.account.name}>
                <Input value={name} onChange={(e) => setName(e.target.value)} />
              </Field>
              <Field label={t.app.settings.account.email}>
                <Input type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
              </Field>
              <div className="sm:col-span-2">
                <Field label={t.app.settings.account.bio}>
                  <Textarea
                    value={bio}
                    onChange={(e) => setBio(e.target.value)}
                    rows={3}
                    placeholder="Tell us a bit about your cooking style…"
                  />
                </Field>
              </div>
            </div>

            <div className="flex items-center justify-end gap-2 pt-2">
              {saved && (
                <span className="text-xs text-accent-active flex items-center gap-1">
                  <Check className="h-3 w-3" />
                  {t.app.settings.account.saved}
                </span>
              )}
              <Button onClick={onSave}>{t.app.settings.account.save}</Button>
            </div>
          </motion.div>
        </TabsContent>

        <TabsContent value="preferences">
          <motion.div
            initial={{ opacity: 0, y: 8 }}
            animate={{ opacity: 1, y: 0 }}
            className="rounded-2xl border border-border bg-bg-card divide-y divide-border"
          >
            <Row label={t.app.settings.preferences.language}>
              <div className="inline-flex items-center gap-1 p-1 rounded-full bg-bg-section border border-border">
                {[
                  { value: "en" as const, label: "English" },
                  { value: "id" as const, label: "Bahasa" },
                ].map((opt) => (
                  <button
                    key={opt.value}
                    type="button"
                    onClick={() => setLang(opt.value)}
                    className={
                      "px-3 py-1.5 text-xs font-medium rounded-full transition-colors " +
                      (lang === opt.value
                        ? "bg-bg-card text-text-primary shadow-sm"
                        : "text-text-muted hover:text-text-primary")
                    }
                  >
                    {opt.label}
                  </button>
                ))}
              </div>
            </Row>

            <Row label={t.app.settings.preferences.defaultServings}>
              <div className="flex items-center gap-2">
                <button
                  type="button"
                  onClick={() => setPreference("defaultServings", Math.max(1, prefs.defaultServings - 1))}
                  className="h-8 w-8 rounded-md border border-border bg-bg-section text-text-muted hover:text-text-primary"
                >
                  -
                </button>
                <span className="w-8 text-center text-sm font-medium text-text-primary tabular-nums">
                  {prefs.defaultServings}
                </span>
                <button
                  type="button"
                  onClick={() => setPreference("defaultServings", prefs.defaultServings + 1)}
                  className="h-8 w-8 rounded-md border border-border bg-bg-section text-text-muted hover:text-text-primary"
                >
                  +
                </button>
              </div>
            </Row>

            <Row label={t.app.settings.preferences.units}>
              <div className="inline-flex items-center gap-1 p-1 rounded-full bg-bg-section border border-border">
                {[
                  { value: "metric" as const, label: t.app.settings.preferences.unitsMetric },
                  { value: "imperial" as const, label: t.app.settings.preferences.unitsImperial },
                ].map((opt) => (
                  <button
                    key={opt.value}
                    type="button"
                    onClick={() => setPreference("units", opt.value)}
                    className={
                      "px-3 py-1.5 text-xs font-medium rounded-full transition-colors " +
                      (prefs.units === opt.value
                        ? "bg-bg-card text-text-primary shadow-sm"
                        : "text-text-muted hover:text-text-primary")
                    }
                  >
                    {opt.label}
                  </button>
                ))}
              </div>
            </Row>
          </motion.div>
        </TabsContent>

        <TabsContent value="notifications">
          <motion.div
            initial={{ opacity: 0, y: 8 }}
            animate={{ opacity: 1, y: 0 }}
            className="rounded-2xl border border-border bg-bg-card divide-y divide-border"
          >
            <Row
              label={t.app.settings.notifications.weekly}
              description="Sunday digest with new recipe ideas based on your pantry."
            >
              <Switch
                checked={prefs.notifications.weekly}
                onCheckedChange={(v) => setNotification("weekly", v)}
              />
            </Row>
            <Row
              label={t.app.settings.notifications.newFeature}
              description="Get notified when we ship something new."
            >
              <Switch
                checked={prefs.notifications.newFeature}
                onCheckedChange={(v) => setNotification("newFeature", v)}
              />
            </Row>
            <Row
              label={t.app.settings.notifications.marketing}
              description="Occasional tips, stories, and product news."
            >
              <Switch
                checked={prefs.notifications.marketing}
                onCheckedChange={(v) => setNotification("marketing", v)}
              />
            </Row>
          </motion.div>
        </TabsContent>

        <TabsContent value="billing">
          <motion.div
            initial={{ opacity: 0, y: 8 }}
            animate={{ opacity: 1, y: 0 }}
            className="rounded-2xl border border-border bg-bg-section p-6"
          >
            <div className="flex items-start justify-between gap-4 mb-4">
              <div>
                <div className="text-[10px] font-semibold text-text-muted uppercase tracking-[0.12em]">
                  {t.app.settings.billing.plan}
                </div>
                <h2 className="font-serif text-2xl font-medium text-text-primary mt-1 capitalize">
                  {user?.plan} Plan
                </h2>
                <p className="text-sm text-text-muted mt-1">
                  {t.app.settings.billing.next}: Feb 12, 2026
                </p>
              </div>
              <Badge variant="accent" className="text-xs">Active</Badge>
            </div>

            <div className="grid sm:grid-cols-3 gap-3 mb-5">
              <div className="p-3 rounded-xl border border-border bg-bg-card">
                <div className="text-[10px] uppercase tracking-[0.1em] text-text-subtle">Used</div>
                <div className="text-lg font-semibold text-text-primary tabular-nums">147 / ∞</div>
              </div>
              <div className="p-3 rounded-xl border border-border bg-bg-card">
                <div className="text-[10px] uppercase tracking-[0.1em] text-text-subtle">Favorites</div>
                <div className="text-lg font-semibold text-text-primary tabular-nums">12</div>
              </div>
              <div className="p-3 rounded-xl border border-border bg-bg-card">
                <div className="text-[10px] uppercase tracking-[0.1em] text-text-subtle">Member since</div>
                <div className="text-lg font-semibold text-text-primary">Aug 2025</div>
              </div>
            </div>

            <div className="flex flex-wrap gap-2">
              <Button variant="outline">{t.app.settings.billing.upgrade}</Button>
              <Button variant="ghost">{t.app.settings.billing.cancel}</Button>
            </div>
          </motion.div>

          <div className="mt-4 rounded-2xl border border-border bg-bg-card p-5">
            <h3 className="text-[10px] font-semibold uppercase tracking-[0.12em] text-text-muted mb-3">
              {t.app.settings.billing.invoices}
            </h3>
            <ul className="divide-y divide-border">
              {["2026-01-12", "2025-12-12", "2025-11-12", "2025-10-12"].map((d) => (
                <li key={d} className="flex items-center justify-between py-2.5 text-sm">
                  <div>
                    <div className="text-text-primary">Pro plan</div>
                    <div className="text-xs text-text-muted">{d}</div>
                  </div>
                  <div className="flex items-center gap-3">
                    <span className="text-text-primary tabular-nums">$12.00</span>
                    <button className="text-xs text-accent-active hover:text-text-primary">Download</button>
                  </div>
                </li>
              ))}
            </ul>
          </div>
        </TabsContent>

        <TabsContent value="danger">
          <motion.div
            initial={{ opacity: 0, y: 8 }}
            animate={{ opacity: 1, y: 0 }}
            className="rounded-2xl border border-danger/30 bg-danger/[0.04] p-6"
          >
            <div className="flex items-start gap-3 mb-4">
              <div className="flex h-10 w-10 items-center justify-center rounded-lg bg-danger/10 text-danger">
                <AlertTriangle className="h-5 w-5" />
              </div>
              <div>
                <h3 className="text-base font-medium text-danger">
                  {t.app.settings.danger.title}
                </h3>
                <p className="mt-1 text-sm text-text-muted">
                  {t.app.settings.danger.desc}
                </p>
              </div>
            </div>
            <Button
              variant="destructive"
              onClick={() => {
                if (window.confirm("This action is permanent. Continue?")) {
                  toast.error("Account deleted (demo)");
                }
              }}
            >
              {t.app.settings.danger.button}
            </Button>
          </motion.div>
        </TabsContent>
      </Tabs>
    </div>
  );
}

function Field({ label, children }: { label: string; children: React.ReactNode }) {
  return (
    <div>
      <label className="text-xs font-medium text-text-muted mb-1.5 block">{label}</label>
      {children}
    </div>
  );
}

function Row({
  label,
  description,
  children,
}: {
  label: string;
  description?: string;
  children: React.ReactNode;
}) {
  return (
    <div className="flex items-center justify-between gap-4 p-5">
      <div>
        <div className="text-sm font-medium text-text-primary">{label}</div>
        {description && <div className="text-xs text-text-muted mt-0.5">{description}</div>}
      </div>
      <div>{children}</div>
    </div>
  );
}
