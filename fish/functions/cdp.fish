function cdp
	source /etc/prt/config

	set location (prt loc $argv ^/dev/null)

	if test -n "$location"
		cd $prtdir/$location
	else
		cd $prtdir
	end
end
