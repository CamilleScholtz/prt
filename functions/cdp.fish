function cdp
	set location (prtloc $argv)
	if test $status -eq 0
		cd $location
	else
		cd /usr/ports
	end
end
