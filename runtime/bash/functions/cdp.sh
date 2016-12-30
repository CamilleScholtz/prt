cdp () {
    portdir=$(grep 'portdir.*' /etc/prt/config.toml | cut -d '=' -f 2 | tr -d '" ')
    loc=$(prt loc "$argv" 2>/dev/null)

    if [ -n "$loc" ]; then
        cd "$portdir/$loc"
    else
        cd "$portdir"
    fi
}
