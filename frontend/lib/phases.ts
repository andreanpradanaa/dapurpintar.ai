// Phase labels for the AI generation animation. Kept in a separate
// file so the UI (generation-loader.tsx) can import it without
// pulling in the API client.

export type GenerationPhase = {
  id: number;
  label: { en: string; id: string };
};

export const GENERATION_PHASES: GenerationPhase[] = [
  {
    id: 1,
    label: { en: "Looking at what you have", id: "Melihat bahan Anda" },
  },
  {
    id: 2,
    label: { en: "Pairing flavors", id: "Memasangkan rasa" },
  },
  {
    id: 3,
    label: { en: "Composing the steps", id: "Menyusun langkah" },
  },
  {
    id: 4,
    label: { en: "Plating up", id: "Menyajikan" },
  },
];
