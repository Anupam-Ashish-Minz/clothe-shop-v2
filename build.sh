#!/bin/sh

# air &
# cd templates
# nodemon -x "templ generate" -e "templ" &
# tailwindcss -i input.css -o output.css --watch

concurrently --kill-others "air"\
	"cd templates && nodemon -x 'templ generate' -e templ"\
	"cd templates && tailwindcss -i input.css -o output.css --watch"
