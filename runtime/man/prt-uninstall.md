# prt-uninstall 8 "2017-01-02" prt "General Commands Manual"

## NAME

prt-uninstall - uninstall packages


## SYNOPSIS

prt uninstall [arguments] [packages]


## DESCRIPTION

prt uninstall uninstalls installed packages, it's just a wrapper for `pkgrm(8)`, but accepts
multiple packages.


## EXAMPLES

Uninstall mpv and fish:

```
# prt uninstall mpv fish
Uninstalling package 1/2, 6c37-git/mpv.
Uninstalling package 2/2, 6c37-git/fish.
```


## AUTHORS

Camille Scholtz


## SEE ALSO

`cdp(1)`, `prt(8)`, `prt.toml(5)`, `prt-depends(8)`, `prt-diff(8)`, `prt-info(8)`, `prt-install(8)`, 
`prt-list(8)`, `prt-loc(8)`, `prt-prov(8)`, `prt-pull(8)`, `prt-sysup(8), prt-get(8)`,
`pkgmk(8)`, `pkgrm(8)`, `pkgadd(8)`, `ports(8)`, `pkginfo(8)`, `prt-utils(1)`
