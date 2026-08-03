/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,ts,jsx,tsx,mdx}"],
  theme: {
    extend: {
      colors: {
        kuali: { 950: "#1a1410", 700: "#3d362e", 500: "#4a3f35" },
        rempah: { 500: "#c75b2a", 400: "#e06a35", 700: "#8b3a1a" },
        santan: { "050": "#fef9f0", "100": "#faf3e6", "200": "#f0e6d3" },
        daun: { 600: "#3d5c3a", 400: "#5a8255" },
        bambu: { 200: "#e8e0d2", 300: "#d4c9b5" },
        canvas: { DEFAULT: "#F9F7F2", alt: "#F3F0EA" },
        surface: "#FFFFFF",
        ink: { DEFAULT: "#1C1B17", muted: "#6B6860", soft: "#9B9890" },
        sage: { DEFAULT: "#5B7B5A", soft: "#E8F0E6", deep: "#3D5A3C", glow: "#A8C4A6" },
        amber: { DEFAULT: "#E89143", soft: "#FEF3EB", deep: "#C97530" },
        mist: "#E8E4DB",
        line: "#E8E4DB",
        yellow: { DEFAULT: "#F4C95D", soft: "#FDF8EA" },
        emerald: { DEFAULT: "#167A58", soft: "#E7F4EE", deep: "#0F5C41" },
        orange: { DEFAULT: "#F28C45", soft: "#FEF3EB" },
      },
      fontFamily: {
        display: ['"Instrument Serif"', "Georgia", "serif"],
        sans: ['"Inter"', "system-ui", "sans-serif"],
        body: ['"Inter"', "system-ui", "sans-serif"],
        mono: ['"JetBrains Mono"', "monospace"],
      },
      fontSize: {
        "hero": ["clamp(3.5rem, 8vw, 6rem)", { lineHeight: "0.95", letterSpacing: "-0.03em" }],
        "hero-mobile": ["clamp(2.5rem, 10vw, 3.5rem)", { lineHeight: "1.0", letterSpacing: "-0.02em" }],
        "display": ["clamp(2rem, 5vw, 3.5rem)", { lineHeight: "1.05", letterSpacing: "-0.02em" }],
        "display-sm": ["clamp(1.5rem, 3vw, 2.25rem)", { lineHeight: "1.1", letterSpacing: "-0.015em" }],
        "body-lg": ["1.125rem", { lineHeight: "1.7" }],
        "body-xl": ["1.25rem", { lineHeight: "1.7" }],
        "eyebrow": ["0.6875rem", { lineHeight: "1.4", letterSpacing: "0.15em" }],
        "data": ["0.8125rem", { lineHeight: "1.5" }],
        "data-lg": ["0.9375rem", { lineHeight: "1.5" }],
      },
      spacing: {
        "section": "clamp(6rem, 12vh, 10rem)",
        "section-sm": "clamp(4rem, 8vh, 6rem)",
      },
      borderRadius: {
        "4xl": "2rem",
        "5xl": "2.5rem",
      },
      boxShadow: {
        "product": "0 4px 60px rgba(28, 27, 23, 0.08), 0 1px 3px rgba(28, 27, 23, 0.04)",
        "product-lg": "0 8px 80px rgba(28, 27, 23, 0.12), 0 2px 8px rgba(28, 27, 23, 0.04)",
        "product-xl": "0 20px 120px rgba(28, 27, 23, 0.16), 0 4px 16px rgba(28, 27, 23, 0.04)",
        "glow-sage": "0 0 60px rgba(91, 123, 90, 0.15)",
        "glow-amber": "0 0 60px rgba(232, 145, 67, 0.12)",
        "inset-soft": "inset 0 1px 0 rgba(255,255,255,0.6)",
      },
    },
  },
  plugins: [],
};
