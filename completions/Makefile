include ../config.mk

COMPLETIONS = \
	cdp.fish \
	depls.fish \
	depmk.fish \
	prtdiff.fish \
	prtloc.fish \
	prtls.fish \
	prtpatch.fish \
	prtprin.fisht \
	prtprovide.fish \
	prtpull.fish

install:
	@echo "Installing completions."
	@for completion in $(COMPLETIONS); do \
		$(INSTALL_FILE) $$completion $(DESTDIR)$(PREFIX)/share/fish/completions/; \
	done

uninstall:
	@echo "Uninstalling completions."
	@for completion in $(COMPLETIONS); do \
		$(RM) $(DESTDIR)$(PREFIX)/share/fish/completions/$$completion; \
	done
