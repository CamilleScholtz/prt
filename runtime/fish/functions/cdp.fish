function cdp
    set prtdir (cat /etc/prt/config.toml | string match -r 'prtdir.*' | string trim '" ')
    set loc (prt loc $argv ^/dev/null)

    if test "$prtdir/$loc"
        cd $prtdir/$loc
    else
        cd $prtdir
    end
end
