/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#f4f6f2',
          100: '#e6ebe0',
          200: '#cfd9c4',
          300: '#afcc85',
          400: '#8db069',
          500: '#5b6a45',
          600: '#4a5538',
          700: '#3a422c',
          800: '#2d3323',
          900: '#1f231a',
          950: '#0f110d',
        },
        background: '#fafafa',
        surface: '#ffffff',
        'surface-secondary': '#f5f5f5',
        border: '#e5e5e5',
        success: {
          light: '#86efac',
          DEFAULT: '#22c55e',
          dark: '#16a34a',
        },
        error: {
          light: '#fca5a5',
          DEFAULT: '#ef4444',
          dark: '#dc2626',
        },
        warning: {
          light: '#fcd34d',
          DEFAULT: '#f59e0b',
          dark: '#d97706',
        },
        info: {
          light: '#93c5fd',
          DEFAULT: '#3b82f6',
          dark: '#2563eb',
        },
      },
    },
  },
  plugins: [],
}

