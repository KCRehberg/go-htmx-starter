/** @type {import('tailwindcss').Config} */
module.exports = {
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

