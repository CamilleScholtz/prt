include config.mk

BINARIES = \
	locprt \
	lsdep \
	lsdiff \
	lsprt \
	mkdep \
	mkdiff \
	mkprt \
	patchprt \
	printprt \
	provprt \
	pullprt

all:
	@echo Run \'make install\' to install prtstuff.

install:
	@echo "Installing binaries."
	@for binary in $(BINARIES); do \
		$(INSTALL_PROG) $$binary $(DESTDIR)$(PREFIX)/bin/$$binary; \
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
