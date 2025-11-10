import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#f4f6f2',
          100: '#e6ebe0',
          200: '#cfd9c4',
          300: '#afcc85',  // Your lighter green
          400: '#8db069',
          500: '#5b6a45',  // Your main color
          600: '#4a5538',
          700: '#3a422c',
          800: '#2d3323',
          900: '#1f231a',
          950: '#0f110d',
        },
        // Neutral grays
        background: '#fafafa',
        surface: '#ffffff',
        'surface-secondary': '#f5f5f5',
        border: '#e5e5e5',

        // Semantic colors
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
};

export default config;
