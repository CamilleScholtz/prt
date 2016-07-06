include ../config.mk

CONFIGS = \
	config
REPOS = \
	pull/6c37 \
	pull/6c37-git \
	pull/contrib \
	pull/core \
	pull/opt \
	pull/xorg 

install:
	@echo "Installing configs."
	$(INSTALL_DIR) $(DESTDIR)/etc/prtstuff/
	$(INSTALL_DIR) $(DESTDIR)/etc/prtstuff/patches
	$(INSTALL_DIR) $(DESTDIR)/etc/prtstuff/pull
	@for config in $(CONFIGS); do \
		$(INSTALL_FILE) $$config $(DESTDIR)/etc/prtstuff/pull/$$config; \
	done
	@for repo in $(REPOS); do \
		$(INSTALL_FILE) $$repo $(DESTDIR)/etc/prtstuff/pull/$$repo; \
	done

uninstall:
	@echo "Uninstalling configs."
	$(RM) -r $(DESTDIR)/etc/prtstuff/ 
