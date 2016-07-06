include ../config.mk

FUNCTIONS = \
	cdp.fish

install:
	@echo "Installing functions."
	@for function in $(FUNCTIONS); do \
		$(INSTALL_FILE) $$function $(DESTDIR)$(PREFIX)/share/fish/functions/; \
	done

uninstall:
	@echo "Uninstalling functions."
	@for function in $(FUNCTIONS); do \
		$(RM) $(DESTDIR)$(PREFIX)/share/fish/functions/$$function; \
	done
