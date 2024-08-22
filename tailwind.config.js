const colors = require('tailwindcss/colors')

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [ "./**/*.html", "./**/*.templ", "./**/*.go", ],
  theme: {
    extend: {
      colors: {
        primary: colors.blue,
        secondary: colors.gray,
      }
    },
  },
  plugins: [
    require('daisyui')
  ],
  colours: {
  }
}

