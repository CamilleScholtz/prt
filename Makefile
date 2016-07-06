include config.mk

BINARIES = \
	depls \
	depmk \
	prtdiff \
	prtloc \
	prtls \
	prtpatch \
	prtprint \
	prtprovide \
	prtpull

all:
	@echo Run \'make install\' to install prtstuff.

install:
	@echo "Installing binaries."
	@for binary in $(BINARIES); do \
		$(INSTALL_FILE) $$binary $(DESTDIR)$(PREFIX)/bin/$$binary; \
	done
	cd configs; $(MAKE) install
	cd completions; $(MAKE) install
	cd functions; $(MAKE) install
	cd libraries; $(MAKE) install

uninstall:
	@echo "Uninstalling binaries."
	@for binary in $(BINARIES); do \
		$(RM) $(DESTDIR)$(PREFIX)/bin/$$binary; \
	done
	cd configs; $(MAKE) uninstall
	cd completions; $(MAKE) uninstall
	cd functions; $(MAKE) uninstall
	cd libraries; $(MAKE) uninstall
