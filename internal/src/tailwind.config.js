/** @type {import('tailwindcss').Config} */
module.exports = {
	mode: "jit",
	content: [
		"./templates/**/*.html",
		"./js/*.js"
	],
	theme: {
		extend: {
			colors: {
				primary: "#0000FF"
			}
		},
	},
	plugins: [],
}

