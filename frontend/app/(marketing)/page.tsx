import { Hero } from "@/components/marketing/hero";
import { TrustedBar } from "@/components/marketing/trusted-bar";
import { Features } from "@/components/marketing/features";
import { ProductPreview } from "@/components/marketing/product-preview";
import { HowItWorks } from "@/components/marketing/how-it-works";
import { PricingTable } from "@/components/marketing/pricing-table";
import { FAQ } from "@/components/marketing/faq";
import { CTA } from "@/components/marketing/cta";

export default function HomePage() {
  return (
    <>
      <Hero />
      <TrustedBar />
      <Features />
      <ProductPreview />
      <HowItWorks />
      <PricingTable />
      <FAQ />
      <CTA />
    </>
  );
}
