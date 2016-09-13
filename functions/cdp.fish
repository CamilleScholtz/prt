function cdp
	source /etc/prt/config

	set location (prt location $argv ^/dev/null)

	if test -n "$location"
		cd $portdir/$location
	else
		cd $portdir
	end
end
