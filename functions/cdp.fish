function cdp
	source /etc/prtstuff/config

	set location (prtloc $argv ^/dev/null)

	if test -n "$location"
		cd $portdir/$location
	else
		cd $portdir
	end
end
