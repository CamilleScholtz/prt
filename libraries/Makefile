include ../config.mk

LIBRARIES = \
	build \
	depends \
	diff \
	info \
	list \
	location \
	patch \
	provide \
	pull \
	sysup \
	util \

install:
	@echo "Installing libraries."
	@for library in $(LIBRARIES); do \
		$(INSTALL_PROG) $$library $(DESTDIR)$(PREFIX)/share/prt/$$library; \
	done

uninstall:
	@echo "Uninstalling libraries."
	@for library in $(LIBRARIES); do \
		$(RM) $(DESTDIR)$(PREFIX)/share/prt/$$library; \
	done
