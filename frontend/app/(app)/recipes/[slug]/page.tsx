import { RecipeDetail } from "@/components/app/recipe-detail";

export default async function RecipePage({
  params,
}: {
  params: Promise<{ slug: string }>;
}) {
  const { slug } = await params;
  return <RecipeDetail slug={slug} />;
}

export async function generateMetadata({ params }: { params: Promise<{ slug: string }> }) {
  const { slug } = await params;
  // Static metadata — fetched in component
  return { title: slug.replace(/-/g, " ") };
}
