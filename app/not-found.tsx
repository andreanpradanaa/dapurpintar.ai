import Link from "next/link";
import { ChefHat, ArrowRight } from "lucide-react";
import { Logo } from "@/components/ui/logo";

export default function NotFound() {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center px-6 bg-bg-base relative overflow-hidden">
      <div className="absolute inset-0 -z-10 bg-paper opacity-50" />
      <div className="relative text-center max-w-lg">
        <Link href="/" className="inline-flex justify-center mb-12">
          <Logo size="md" />
        </Link>
        <div className="inline-flex items-center justify-center mb-8">
          <span className="font-display text-[140px] sm:text-[180px] font-medium leading-none text-text-primary tracking-tighter">
            404
          </span>
        </div>
        <h1 className="font-display text-3xl sm:text-4xl font-medium tracking-tight text-text-primary text-balance">
          Lost in the pantry.
        </h1>
        <p className="mt-4 text-text-muted leading-[1.6] max-w-md mx-auto">
          We can&apos;t find the page you were looking for. It may have been moved, or perhaps it was never cooked up in the first place.
        </p>
        <div className="mt-10 flex flex-col sm:flex-row items-center justify-center gap-3">
          <Link href="/">
            <Button className="group">
              Back home
              <ArrowRight className="h-4 w-4 transition-transform group-hover:translate-x-0.5" />
            </Button>
          </Link>
          <Link href="/generate">
            <button className="inline-flex items-center gap-1.5 h-11 px-5 rounded-lg text-sm font-medium text-text-secondary hover:text-text-primary transition-colors">
              <ChefHat className="h-4 w-4" />
              Create a recipe
            </button>
          </Link>
        </div>
      </div>
    </div>
  );
}

function Button({
  className,
  children,
}: {
  className?: string;
  children: React.ReactNode;
}) {
  return (
    <button
      type="button"
      className={
        "inline-flex items-center justify-center gap-2 h-11 px-5 rounded-lg text-sm font-medium bg-accent text-accent-fg hover:bg-accent-hover transition-colors " +
        (className ?? "")
      }
    >
      {children}
    </button>
  );
}
