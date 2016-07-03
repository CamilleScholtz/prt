function cdp
	source /etc/prtstuff/config

	set location (prtloc $argv)
	if test $status -eq 0
		cd $portdir/$location
	else
		cd $portdir
	end
end
