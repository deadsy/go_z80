#------------------------------------------------------------------------------
# Makefile to get external sources.

TOP = $(shell realpath ..)
DL_DIR = $(TOP)/dl

.PHONY: all
all: .stamp_z80step

.PHONY: clean
clean:
	-rm -rf $(Z80STEP_DIR)
	-rm .stamp*

#-------------------------------------------------------------------------------
# Z80 Step

Z80STEP_VER = v1.0-beta.2
Z80STEP_URL = https://github.com/SingleStepTests/z80/archive/refs/tags/$(Z80STEP_VER).tar.gz
Z80STEP_FILE = z80step-$(Z80STEP_VER).tar.gz
Z80STEP_TGZ = $(DL_DIR)/$(Z80STEP_FILE)
Z80STEP_DIR = $(TOP)/ext/z80step

$(Z80STEP_TGZ):
	mkdir -p $(DL_DIR)
	wget $(Z80STEP_URL) -O $(Z80STEP_TGZ)

.stamp_z80step: $(Z80STEP_TGZ)
	mkdir -p $(Z80STEP_DIR)
	tar -C $(Z80STEP_DIR) -zxf $(Z80STEP_TGZ) --strip-components 1
	touch $@

#-------------------------------------------------------------------------------
