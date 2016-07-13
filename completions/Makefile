include ../config.mk

COMPLETIONS = \
	cdp.fish \
	depls.fish \
	depmk.fish \
	diffls.fish \
	diffmk.fish \
	prtloc.fish \
	prtls.fish \
	prtmk.fish \
	prtpatch.fish \
	prtprint.fish \
	prtprov.fish \
	prtpull.fish

install:
	@echo "Installing completions."
	@for completion in $(COMPLETIONS); do \
		$(INSTALL_FILE) $$completion $(DESTDIR)$(PREFIX)/share/fish/completions/$$completion; \
	done

uninstall:
	@echo "Uninstalling completions."
	@for completion in $(COMPLETIONS); do \
		$(RM) $(DESTDIR)$(PREFIX)/share/fish/completions/$$completion; \
	done
