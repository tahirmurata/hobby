/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        dark: {
          ...require("daisyui/src/theming/themes")["dark"],
          ".btn-disabled": {
            "--tw-text-opacity": 1,
            color: "var(--fallback-suc,oklch(0.648 0.15 160))",
            "outline-color": "var(--fallback-su,oklch(0.648 0.15 160))",
          },
          ".btn-outline.btn-disabled": {
            "--tw-text-opacity": 1,
            color: "var(--fallback-su,oklch(0.648 0.15 160))",
          },
        },
      },
    ],
  },
};
