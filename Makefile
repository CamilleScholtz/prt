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
	$(INSTALL_DIR) $(DESTDIR)$(PREFIX)/bin/
	@for binary in $(BINARIES); do \
		$(INSTALL_FILE) $$binary $(DESTDIR)$(PREFIX)/bin/$$binary; \
	done
	cd configs; $(MAKE) install
	cd completions; $(MAKE) install
	cd functions; $(MAKE) install
	cd libraries; $(MAKE) install


	$(INSTALL_DIR) $(DESTDIR)$(PREFIX)/share/fish/functions


uninstall:
	@echo "Uninstalling binaries."
	@for binary in $(BINARIES); do \
		$(RM) $(DESTDIR)$(PREFIX)/share/fish/completions/$$binary; \
	done
	cd configs; $(MAKE) install
	cd completions; $(MAKE) install
	cd functions; $(MAKE) uninstall
	cd libraries; $(MAKE) uninstall
