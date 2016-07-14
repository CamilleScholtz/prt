include ../config.mk

COMPLETIONS = \
	cdp.fish \
	locprt.fish \
	lsdep.fish \
	lsdiff.fish \
	lsprt.fish \
	mkdep.fish \
	mkdiff.fish \
	mkprt.fish \
	patchprt.fish \
	printprt.fish \
	provprt.fish \
	pullprt.fish

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
