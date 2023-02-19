/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './index.html', 
    './src/**/*.{html,vue,js,ts}'
  ],
  theme: {
    extend: {
      colors: {
        'brand-black': 'var(--brand-black)',
        'brand-red': 'var(--brand-red)',
      },
    },
  },
  plugins: [],
}

