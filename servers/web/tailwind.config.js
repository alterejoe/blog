/** @type {import('tailwindcss').Config} */

import { iconifyPlugin } from "@iconify/tailwind";

export default {
	content: [
		"./ui/**/*.{templ,html,js}",
		"./shared/components/**/*.{templ,html,js}",
		{
			files: ["./**/*.go"],
			extract: {
				go: (content) => {
					const regex = /Class:\s*"([^"]+)"/g;
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
		extend: {},
	},
	plugins: [iconifyPlugin()],
};
