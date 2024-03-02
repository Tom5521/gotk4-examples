#!/usr/bin/bash

if [ "$1" == "" ]; then
	echo The name of the new class is required as first argument
	exit 1
fi

if [ "$1" == "help" ]; then
	echo This creates the folder, the README.md, and a main.go template for an example, \
		I use it to generate new examples quickly, because before what I did was pure copy-paste.
	echo
	echo The use is simple, just run ./generate.sh [class name] \
		and the script creates the folder and the corresponding readme.
	exit 0
fi

if [ -d "$1" ]; then
	echo The folder already exists
	exit 1
fi

new_folder="$1"
new_readme=$(cat generator/README_example.md)

mkdir -p "$new_folder"
printf "$new_readme" "$1" >"$new_folder/README.md"
cp generator/main_example.go "$new_folder/main.go"
