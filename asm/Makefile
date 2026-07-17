

.PHONY: all
all:
	pasmo -d --hex segments.asm segments.hex > segments.lst
	pasmo -d --hex hd44780_test.asm hd44780_test.hex > hd44780_test.lst

.PHONY: clean
clean:
	-rm *.hex
	-rm *.bin
	-rm *.lst

