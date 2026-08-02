
STACK = 0x7ffe
CODE = 0x4000
DATA = 0x6000

.PHONY: all
all:
	sdcc -mz80 --no-std-crt0 --code-loc $(CODE) --data-loc $(DATA) --stack-loc $(STACK) main.c


.PHONY: clean
clean:
	-rm *.asm
	-rm *.ihx
	-rm *.lk
	-rm *.lst
	-rm *.map
	-rm *.noi
	-rm *.rel
	-rm *.sym
