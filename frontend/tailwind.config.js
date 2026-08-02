/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,ts,jsx,tsx,mdx}"],
  theme: {
    extend: {
      colors: {
        ink: { 950: "#101A2B", 900: "#17243A", 700: "#3D4B63" },
        steel: { 400: "#8793A5", 200: "#D5DAE1" },
        paper: { "050": "#F4F1E8", "000": "#FFFEFA" },
        white: { "000": "#FFFFFF" },
        action: { primary: "#F45B3C", dark: "#A52F1C" },
        context: { positive: "#A7D46F", "positive-dark": "#5D7D36", attention: "#F3C969", "attention-dark": "#765816" },
        feedback: { error: "#D64545", info: "#4C82C3" },
      },
      fontFamily: {
        sans: ["Inter", "system-ui", "sans-serif"],
        mono: ["JetBrains Mono", "monospace"],
      },
    },
  },
  plugins: [],
};
