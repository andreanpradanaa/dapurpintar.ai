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
      },
      fontFamily: {
        display: ['"Playfair Display"', "Georgia", "serif"],
        sans: ['"DM Sans"', "system-ui", "sans-serif"],
        mono: ['"JetBrains Mono"', "monospace"],
      },
    },
  },
  plugins: [],
};
