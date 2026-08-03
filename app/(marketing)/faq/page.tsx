import type { Metadata } from "next";
import { FAQ } from "@/components/marketing/faq";
import { CTA } from "@/components/marketing/cta";

export const metadata: Metadata = {
  title: "FAQ",
  description: "Common questions about Dapur Pintar AI.",
};

export default function FAQPage() {
  return (
    <>
      <div className="pt-24" />
      <FAQ />
      <CTA />
    </>
  );
}
