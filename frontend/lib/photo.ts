// Curated warm, editorial food photography from Unsplash
// All images get a CSS filter (photo-warm) for consistent treatment
// Replace these with brand photography in production

export const FOOD_PHOTOS = {
  hero: "https://images.unsplash.com/photo-1556909114-f6e7ad7d3136?w=1600&q=80&auto=format&fit=crop",
  heroAlt1: "https://images.unsplash.com/photo-1504674900247-0877df9cc836?w=1200&q=80&auto=format&fit=crop",
  heroAlt2: "https://images.unsplash.com/photo-1490645935967-10de6ba17061?w=1200&q=80&auto=format&fit=crop",
  // ...
} as const;

// Per-recipe hero images. Falls back to a warm gradient if missing.
export function photoForRecipe(slug: string): string {
  const map: Record<string, string> = {
    "nasi-goreng-ayam": "https://images.unsplash.com/photo-1603133872878-684f208fb84b?w=1200&q=80&auto=format&fit=crop",
    "soto-ayam": "https://images.unsplash.com/photo-1547592180-85f173990554?w=1200&q=80&auto=format&fit=crop",
    "rendang-sapi": "https://images.unsplash.com/photo-1604908176997-125f25cc6f3d?w=1200&q=80&auto=format&fit=crop",
    "gado-gado": "https://images.unsplash.com/photo-1546069901-ba9599a7e63c?w=1200&q=80&auto=format&fit=crop",
    "mie-goreng-jawa": "https://images.unsplash.com/photo-1612929633738-8fe44f7ec841?w=1200&q=80&auto=format&fit=crop",
    "ayam-bakar-madu": "https://images.unsplash.com/photo-1598103442097-8b74394b95c6?w=1200&q=80&auto=format&fit=crop",
    "sate-lilit-bali": "https://images.unsplash.com/photo-1529193591184-b1d58069ecdd?w=1200&q=80&auto=format&fit=crop",
    "cap-cay-seafood": "https://images.unsplash.com/photo-1626804475297-41608ea09aeb?w=1200&q=80&auto=format&fit=crop",
    "bakso-urat": "https://images.unsplash.com/photo-1569718212165-3a8278d5f624?w=1200&q=80&auto=format&fit=crop",
    "tempe-orek": "https://images.unsplash.com/photo-1625938145744-533e82c1d29f?w=1200&q=80&auto=format&fit=crop",
    "ayam-pop": "https://images.unsplash.com/photo-1610057099443-fde9c4c50ce2?w=1200&q=80&auto=format&fit=crop",
    "rawon-jakarta": "https://images.unsplash.com/photo-1583224944844-5a6c7d3da7c1?w=1200&q=80&auto=format&fit=crop",
    "tumis-kangkung": "https://images.unsplash.com/photo-1627662235936-94c54c2a5183?w=1200&q=80&auto=format&fit=crop",
    "opor-ayam": "https://images.unsplash.com/photo-1604908554027-1c2e0e8c8b3a?w=1200&q=80&auto=format&fit=crop",
    "pempek-palembang": "https://images.unsplash.com/photo-1563245372-f21724e3856d?w=1200&q=80&auto=format&fit=crop",
    "nasi-uduk": "https://images.unsplash.com/photo-1596797038530-2c107229654b?w=1200&q=80&auto=format&fit=crop",
    "sambal-tempe": "https://images.unsplash.com/photo-1604908554144-79f00b8e6b13?w=1200&q=80&auto=format&fit=crop",
    "es-teler": "https://images.unsplash.com/photo-1551024506-0bccd828d307?w=1200&q=80&auto=format&fit=crop",
    "tahu-gejrot": "https://images.unsplash.com/photo-1625938144755-652e08e359b7?w=1200&q=80&auto=format&fit=crop",
    "ikan-bakar-rica": "https://images.unsplash.com/photo-1532980193133-95f4c0d5f6e9?w=1200&q=80&auto=format&fit=crop",
    "martabak-manis": "https://images.unsplash.com/photo-1606756790138-261d2b21cd75?w=1200&q=80&auto=format&fit=crop",
    "bihun-goreng": "https://images.unsplash.com/photo-1569718212165-3a8278d5f624?w=1200&q=80&auto=format&fit=crop",
    "gado-gado-bali": "https://images.unsplash.com/photo-1546069901-ba9599a7e63c?w=1200&q=80&auto=format&fit=crop",
    "tumis-tahu-cah-sawi": "https://images.unsplash.com/photo-1627662235936-94c54c2a5183?w=1200&q=80&auto=format&fit=crop",
    "rendang-jamur": "https://images.unsplash.com/photo-1518977676601-b53f82aba655?w=1200&q=80&auto=format&fit=crop",
    "pempek-kapal-selam": "https://images.unsplash.com/photo-1563245372-f21724e3856d?w=1200&q=80&auto=format&fit=crop",
    "nasi-campur-bali": "https://images.unsplash.com/photo-1567337710282-00832b415979?w=1200&q=80&auto=format&fit=crop",
    "bubur-ayam": "https://images.unsplash.com/photo-1607330289024-1535c6b4e1c1?w=1200&q=80&auto=format&fit=crop",
    "kue-lapis": "https://images.unsplash.com/photo-1565958011703-44f9829ba187?w=1200&q=80&auto=format&fit=crop",
    "tongseng-kambing": "https://images.unsplash.com/photo-1547928576-b822bc410bdf?w=1200&q=80&auto=format&fit=crop",
    "pisang-goreng": "https://images.unsplash.com/photo-1606851094291-6efae152bb87?w=1200&q=80&auto=format&fit=crop",
    "cumi-saus-tiram": "https://images.unsplash.com/photo-1599487488170-d11ec9c172f0?w=1200&q=80&auto=format&fit=crop",
  };
  return map[slug] ?? FOOD_PHOTOS.heroAlt2;
}

// Curated warm pantry / lifestyle / object photos for landing
export const LIFESTYLE_PHOTOS = {
  kitchenTable: "https://images.unsplash.com/photo-1495546968767-f0573cca821e?w=1200&q=80&auto=format&fit=crop",
  ingredientsFlatlay: "https://images.unsplash.com/photo-1505935428862-770b6f24f629?w=1200&q=80&auto=format&fit=crop",
  handWithBowl: "https://images.unsplash.com/photo-1556909114-f6e7ad7d3136?w=1200&q=80&auto=format&fit=crop",
  woodenBoard: "https://images.unsplash.com/photo-1556909114-44e3e9699e2b?w=1200&q=80&auto=format&fit=crop",
  cookOnStove: "https://images.unsplash.com/photo-1556910103-1c02745aae4d?w=1200&q=80&auto=format&fit=crop",
  freshHerbs: "https://images.unsplash.com/photo-1466637574441-749b8f19452f?w=1200&q=80&auto=format&fit=crop",
} as const;
