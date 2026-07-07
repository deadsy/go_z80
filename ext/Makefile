#------------------------------------------------------------------------------
# Makefile to get external sources.

TOP = $(shell realpath ..)
DL_DIR = $(TOP)/dl

.PHONY: all
all: .stamp_z80step .stamp_za

.PHONY: clean
clean:
	-rm -rf $(Z80STEP_DIR)
	-rm -rf $(ZA_DIR)
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
# zexall

ZA_VER = 8f71d41
ZA_URL = https://github.com/agn453/ZEXALL/tarball/$(ZA_VER)
ZA_FILE = zexall-$(ZA_VER).tar.gz
ZA_TGZ = $(DL_DIR)/$(ZA_FILE)
ZA_DIR = $(TOP)/ext/zexall

$(ZA_TGZ):
	mkdir -p $(DL_DIR)
	wget $(ZA_URL) -O $(ZA_TGZ)

.stamp_za: $(ZA_TGZ)
	mkdir -p $(ZA_DIR)
	tar -C $(ZA_DIR) -zxf $(ZA_TGZ) --strip-components 1
	touch $@

#-------------------------------------------------------------------------------
