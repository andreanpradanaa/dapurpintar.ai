import type { Metadata } from "next";
import { PricingTable } from "@/components/marketing/pricing-table";
import { CTA } from "@/components/marketing/cta";

export const metadata: Metadata = {
  title: "Pricing",
  description: "Simple, honest pricing. Start free, upgrade when you need more.",
};

export default function PricingPage() {
  return (
    <>
      <PricingTable />
      <CTA />
    </>
  );
}
