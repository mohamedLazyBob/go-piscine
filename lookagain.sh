#!bin/bash
find . -name "*.sh" | sort -r | rev | cut -d "/" -f1 | rev | sed 's/.sh//g'
# Explanations made for the veiled lady
# not working for nested file
# find . -name "*.sh" | sort -r | sed 's/.sh//g' | sed 's/.\///g'
