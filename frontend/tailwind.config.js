/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,ts,jsx,tsx,mdx}"],
  theme: {
    extend: {
      colors: {
        // Existing app tokens (don't break logged-in pages)
        kuali: { 950: "#1a1410", 700: "#3d362e", 500: "#4a3f35" },
        rempah: { 500: "#c75b2a", 400: "#e06a35", 700: "#8b3a1a" },
        santan: { "050": "#fef9f0", "100": "#faf3e6", "200": "#f0e6d3" },
        daun: { 600: "#3d5c3a", 400: "#5a8255" },
        bambu: { 200: "#e8e0d2", 300: "#d4c9b5" },
        // New landing page tokens (modern SaaS)
        canvas: { DEFAULT: "#FAFAF8", alt: "#F5F6F4" },
        surface: "#FFFFFF",
        ink: { DEFAULT: "#17211B", muted: "#66736A", soft: "#93A099" },
        emerald: { DEFAULT: "#167A58", soft: "#E7F4EE", deep: "#0F5C41" },
        orange: { DEFAULT: "#F28C45", soft: "#FEF3EB" },
        yellow: { DEFAULT: "#F4C95D", soft: "#FDF8EA" },
        line: "#E4EAE5",
      },
      fontFamily: {
        display: ['"DM Sans"', "system-ui", "sans-serif"],
        sans: ['"DM Sans"', "system-ui", "sans-serif"],
        editorial: ['"Playfair Display"', "Georgia", "serif"],
        mono: ['"JetBrains Mono"', "monospace"],
      },
    },
  },
  plugins: [],
};
