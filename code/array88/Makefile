
TARGET = out

STACK = 0x7ffe
CODE = 0x4000
DATA = 0x6000

OBJS = entry.rel \
	main.rel \
	array88.rel \
	display.rel \
	delay.rel

CFLAGS = -mz80 \
	--Werror \
	--std-c99 \
	-DHW_TEC_1G

LFLAGS = -mz80 \
	--no-std-crt0 \
	--code-loc $(CODE) \
	--data-loc $(DATA)

.PHONY: all
all: $(OBJS)
	sdcc $(LFLAGS) $(OBJS) -o $(TARGET).ihx

%.rel : %.s
	sdasz80 -o $<

%.rel : %.c
	sdcc $(CFLAGS) -c $<

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
