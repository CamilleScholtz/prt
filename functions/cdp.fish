function cdp
	# Set config dir location
	set configdir /etc/prtstuff

	set location (prtloc $argv)
	if test $status -eq 0
		cd $portdir/$location
	else
		cd $portdir
	end
end
