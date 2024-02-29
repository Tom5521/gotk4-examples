#!/usr/bin/bash

new_folder="Gtk$1"
new_readme=$(cat ./generator/README_example.md)

mkdir -p "$new_folder"
printf "$new_readme" "$1" >"$new_folder/README.md"
cp ./generator/main_example.go "$new_folder/main.go"
