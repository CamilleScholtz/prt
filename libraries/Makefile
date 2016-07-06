include ../config.mk

LIBRARIES = \
	prtstuff

install:
	@echo "Installing libraries."
	@for library in $(LIBRARIES); do \
		$(INSTALL_FILE) $$library $(DESTDIR)$(PREFIX)/include/$$library; \
	done

uninstall:
	@echo "Uninstalling libraries."
	@for library in $(LIBRARIES); do \
		$(RM) $(DESTDIR)$(PREFIX)/include/$$library; \
	done
