/** @type {import('tailwindcss').Config} */
import { iconifyPlugin } from "@iconify/tailwind";
export default {
    mode: "jit",
    content: ["./ui/html/**/*.{templ,html,js}"],
    theme: {
        extend: {
            colors: {
                primary: "var(--color-primary)",
                secondary: "var(--color-secondary)",
                tertiary: "var(--color-tertiary)",
                quaternary: "var(--color-quaternary)",
                accent: "var(--color-accent)",
                background: "var(--color-background)",
                neutral: "var(--color-neutral)",
                unique: "var(--color-unique)",
                attention: "var(--color-attention)",
            },
        },
    },
    plugins: [
        iconifyPlugin(),
        require("tailwind-scrollbar")({ nocompatible: true }),
    ],
};
