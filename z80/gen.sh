#!/bin/bash
python3 ./z80gen.py
goimports -w opcodes.go
