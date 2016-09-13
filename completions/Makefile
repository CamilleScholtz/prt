include ../config.mk

COMPLETIONS = \
	cdp.fish \
	prt.fish

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
