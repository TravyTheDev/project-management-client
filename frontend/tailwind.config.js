/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        backgroundColor: "rgba(var(--backgroundColor))",
        fontColor: "rgba(var(--fontColor))",
      }
    },
    plugins: [],
  },
}

