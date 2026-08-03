export const designTokens = {
  color: {
    bg: {
      base: "#FAF6F0",
      card: "#FFFFFF",
      elevated: "#FFFFFF",
      section: "#F4EFE5",
    },
    text: {
      primary: "#1A1612",
      secondary: "#3D352D",
      muted: "#6B6358",
      subtle: "#9B9286",
    },
    accent: {
      DEFAULT: "#A8553A",
      hover: "#8E4530",
      active: "#723627",
      soft: "rgba(168, 85, 58, 0.08)",
      softStrong: "rgba(168, 85, 58, 0.14)",
      fg: "#FFFFFF",
    },
    border: {
      DEFAULT: "rgba(26, 22, 18, 0.08)",
      strong: "rgba(26, 22, 18, 0.14)",
      accent: "rgba(168, 85, 58, 0.4)",
    },
    semantic: {
      success: "#5C7A4D",
      warning: "#B8804A",
      danger: "#A14B3A",
      info: "#5A7A8A",
    },
  },
  radius: {
    sm: 6,
    md: 8,
    lg: 12,
    xl: 16,
    "2xl": 20,
    "3xl": 24,
    full: 9999,
  },
  ease: {
    out: [0.22, 1, 0.36, 1] as [number, number, number, number],
    in: [0.4, 0, 1, 1] as [number, number, number, number],
    inOut: [0.4, 0, 0.2, 1] as [number, number, number, number],
    spring: [0.34, 1.56, 0.64, 1] as [number, number, number, number],
  },
  duration: {
    micro: 0.18,
    state: 0.24,
    layout: 0.36,
    reveal: 0.6,
  },
} as const;

export type DesignTokens = typeof designTokens;
