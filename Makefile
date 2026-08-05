
.PHONY: all
all:
	make -C cmd/tec1 $@
	make -C cmd/tec1g $@
	make -C cmd/jace $@
	make -C cmd/lcdtest $@

.PHONY: clean
clean:
	make -C cmd/tec1 $@
	make -C cmd/tec1g $@
	make -C cmd/jace $@
	make -C cmd/lcdtest $@
