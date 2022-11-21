/** @type {import('tailwindcss').Config} */

module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    fontFamily: {},
    extend: {
      backgroundImage: {
        hero: "url('/brooke-lark-RrzeCGujVfU-unsplash.jpg')",
      },
      colors: {
        primary: "#453c38",
        secondary: "#6a5c55",
        accent: "#e5b737",
      },
    },
  },
  plugins: [],
};
