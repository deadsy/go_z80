
TARGET = out

STACK = 0x7ffe
CODE = 0x4000
DATA = 0x6000

OBJS = entry.rel \
	main.rel \
	array88.rel \

.PHONY: all
all: $(OBJS)
	sdcc -mz80 --no-std-crt0 --code-loc $(CODE) --data-loc $(DATA) $(OBJS) -o $(TARGET).ihx

%.rel : %.s
	sdasz80 -o $<

%.rel : %.c
	sdcc -mz80 -c $<

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
