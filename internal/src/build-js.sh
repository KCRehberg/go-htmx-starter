for file in js/*.js; do
    uglifyjs "js/${file##*/}" -c -m -o "../../static/${file##*/}"
    echo minified: "${file##*/}"
done