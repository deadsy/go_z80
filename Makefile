
.PHONY: all
all:
	make -C cmd/tec1 $@
	make -C cmd/segtest $@

.PHONY: clean
clean:
	make -C cmd/tec1 $@
	make -C cmd/segtest $@


