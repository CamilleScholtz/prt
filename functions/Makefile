include ../config.mk

FUNCTIONS = \
	cdp.fish

install:
	@echo "Installing functions."
	$(INSTALL_DIR) $(DESTDIR)$(PREFIX)/usr/share/fish/completions/
	@for function in $(FUNCTIONS); do \
		$(INSTALL_FILE) $$function $(DESTDIR)/etc/prtstuff/pull/$$function; \
	done

uninstall:
	@echo "Uninstalling functions."
	@for function in $(FUNCTIONS); do \
		$(RM) $(DESTDIR)$(PREFIX)/share/fish/completions/$$function; \
	done
