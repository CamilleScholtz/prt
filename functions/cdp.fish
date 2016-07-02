function cdp
	set location (prtloc $argv)
	if test $status -eq 0
		cd /usr/ports/$location
	else
		cd /usr/ports
	end
end
