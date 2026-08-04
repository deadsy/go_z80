#!/bin/bash

FILES="
main.c
array88.c
array88.h
delay.c
delay.h
"

for f in $FILES; do
  indent $f -brf -linux -l10000
	rm $f~
done

