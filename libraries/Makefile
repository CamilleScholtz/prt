include ../config.mk

LIBRARIES = \
	prtstuff

install:
	@echo "Installing libraries."
	$(INSTALL_DIR) $(DESTDIR)$(PREFIX)/usr/share/fish/completions/
	@for library in $(LIBRARIES); do \
		$(INSTALL_FILE) $$library $(DESTDIR)$(PREFIX)/include/$$library; \
	done

uninstall:
	@echo "Uninstalling libraries."
	@for library in $(LIBRARIES); do \
		$(RM) $(DESTDIR)$(PREFIX)/include/$$library; \
	done
