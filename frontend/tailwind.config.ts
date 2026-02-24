import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        background: '#0a0e1a',
        foreground: '#e5e7eb',
        card: '#0f1420',
        border: '#1a2332',
        muted: '#1a2332',
        primary: '#00d9ff',
        accent: {
          cyan: '#00d9ff',
          amber: '#ffb020',
          red: '#ff4757',
        },
      },
      fontFamily: {
        sans: ['General Sans', 'system-ui', 'sans-serif'],
        mono: ['JetBrains Mono', 'Consolas', 'monospace'],
      },
    },
  },
};

export default config;