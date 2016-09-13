include config.mk

EXECUTABLES = \
	prt

all:
	@echo Run \'make install\' to install prt.

install:
	@echo "Installing executables."
	@for executable in $(EXECUTABLES); do \
		$(INSTALL_PROG) $$executable $(DESTDIR)$(PREFIX)/bin/$$executable; \
	done
	cd configs; $(MAKE) install
	cd completions; $(MAKE) install
	cd functions; $(MAKE) install
	cd libraries; $(MAKE) install

uninstall:
	@echo "Uninstalling executables."
	@for executable in $(EXECUTABLES); do \
		$(RM) $(DESTDIR)$(PREFIX)/bin/$$executable; \
	done
	cd configs; $(MAKE) uninstall
	cd completions; $(MAKE) uninstall
	cd functions; $(MAKE) uninstall
	cd libraries; $(MAKE) uninstall
