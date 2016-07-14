function cdp
	source /etc/prtstuff/config

	set location (locprt $argv ^/dev/null)

	if test -n "$location"
		cd $portdir/$location
	else
		cd $portdir
	end
end
