#!/bin/sh

air &
cd templates
nodemon -x "templ generate" -e "templ" &
tailwindcss -i input.css -o output.css --watch
