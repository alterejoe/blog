/** @type {import('tailwindcss').Config} */

import { iconifyPlugin } from "@iconify/tailwind";

export default {
    content: [
        "./ui/**/*.{templ,html,js}",
        "../../shared/components/**/*.{templ,html,js}",
        {
            files: ["./**/*_classes.go"],
            extract: {
                go: (content) => {
                    const regex = /"([^"]+)"/g;
                    const matches = [];
                    let match;
                    while ((match = regex.exec(content)) !== null) {
                        matches.push(...match[1].split(/\s+/));
                    }
                    return matches;
                },
            },
        },
    ],

    theme: {
        extend: {
            colors: {
                primary: "var(--color-primary)",
                secondary: "var(--color-secondary)",
                tertiary: "var(--color-tertiary)",
                accent: "var(--color-accent)",
                background: "var(--color-background)",
                neutral: "var(--color-neutral)",
            },
        },
    },
    plugins: [iconifyPlugin()],
};
