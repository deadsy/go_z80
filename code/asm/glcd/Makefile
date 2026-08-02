
ASM = pasmo

%.hex:%.z80
	$(ASM) -d --hex $< $@ > $*.lst 

.PHONY: all
all: mad demo

.PHONY: mad
mad: lcd_mad_program.hex

.PHONY: demo
demo: lcd_3d_demo.hex

.PHONY: clean
clean:
	-rm *.hex
	-rm *.bin
	-rm *.lst
