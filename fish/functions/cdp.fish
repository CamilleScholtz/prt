function cdp
	set location (prt loc $argv)

	if test -n "$location"
		cd $prtdir/$location
	else
		cd $prtdir
	end
end
