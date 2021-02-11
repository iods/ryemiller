module.exports = {
    purge: [
        "./public/**/*.html",
        "./src/**/*.vue"],
    theme: {
        extend: {
            colors: {
                gunmetal: {
                    DEFAULT: "#1f2937",
                    600: "#1c2532",

                },
                cultured: {
                    light: "#f7f8f8",
                    DEFAULT: "#F2f4f3"
                },
                keppel: {
                    DEFAULT: "#15b097"
                },
                crayola: {
                    DEFAULT: "#3F88c5",
                    darker: "#2a618d"
                },
                gate: {
                    lightest: "#fbf1ef",
                    lighter: "#eec5be",
                    light: "#de8c7d",
                    DEFAULT: "#bb4430",
                    dark: "#923626",
                    darker: "#612419",
                    darkest: "#31120d"
                }
            }
        },
    },
    variants: {},
    plugins: [],
}