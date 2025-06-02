/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        'title': ['Poppins', 'sans-serif'],
        'subtitle': ['Montserrat', 'sans-serif'],
        'body': ['Nunito', 'Lato', 'sans-serif'],
        'alt': ['Libre Franklin', 'Raleway', 'sans-serif'],
      },
    },
  },
  plugins: [],
} 