document.body.addEventListener('htmx:afterRequest', function(evt) {
	console.log(evt)
});