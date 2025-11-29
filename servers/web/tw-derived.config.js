module.exports = function ({ addBase, theme }) {
    const colors = theme("colors");

    const light = {};
    const dark = {};

    Object.keys(colors).forEach((token) => {
        const base = colors[token];
        if (typeof base !== "string") return;

        // light variants
        light[`--color-${token}-hover`] =
            `color-mix(in srgb, ${base} 90%, white)`;

        // dark variants
        dark[`--color-${token}`] = `color-mix(in srgb, ${base} 70%, black)`;
        dark[`--color-${token}-hover`] =
            `color-mix(in srgb, ${base} 60%, black)`;
    });

    addBase({
        ":root": light,
        ".dark": dark,
    });
};
