function cdp
	set loc (prt loc $argv ^/dev/null)

	if test "$loc"
		cd /usr/ports/$loc
	else
		cd /usr/ports
	end
end
