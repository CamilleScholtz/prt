function cdp
    set portdir (cat /etc/prt/config.toml | string match -r 'prtdir.*' | cut -d '=' -f 2 | string trim -c '" ')
    set loc (prt loc $argv ^/dev/null)

    if test "$loc"
        cd $portdir/$loc
    else
        cd $portdir
    end
end
