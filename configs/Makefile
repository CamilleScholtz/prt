include ../config.mk

CONFIGS = \
	config \
	pull/6c37 \
	pull/6c37-git \
	pull/contrib \
	pull/core \
	pull/opt \
	pull/xorg 

install:
	@echo "Installing configs."
	@for config in $(CONFIGS); do \
		$(INSTALL_FILE) $$config $(DESTDIR)/etc/prtstuff/$$config; \
	done

uninstall:
	@echo "Uninstalling configs."
	$(RM) -r $(DESTDIR)/etc/prtstuff/ 
