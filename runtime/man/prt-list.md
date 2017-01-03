# prt-list 8 "2017-01-02" prt "General Commands Manual"

## NAME

prt-list - list ports and packages


## SYNOPSIS

prt list [arguments]


## DESCRIPTION

prt list lists all ports on the system, or all installed packages.


## OPTIONS

`-i` only lists installed packages. Without this option only ports that are not yet
installed get listed. Same as `--installed`

`-r` list with repo information. Same as `--repo`

`-v` list with version information. Same as `--version`


## EXAMPLES

List all all ports:

```
$ prt list | head -n 5
abduco
acpiclient
afuse
arandr
artwiz-fonts
```

Only list installed ports, with version and repo information: 

```
$ prt list -i -r -v | head -n 5
6c37-dropin/libressl 1.1.2-1
6c37-dropin/pkgconf 1.1.1-1
6c37-git/colorpicker 1.1.2-1
6c37-git/compton 0.60.6.1-2
6c37-git/fish 2016.11.20-0-1
```


## AUTHORS

Camille Scholtz


## SEE ALSO

`cdp(1)`, `prt(8)`, `prt.toml(5)`, `prt-depends(8)`, `prt-diff(8)`, `prt-info(8)`, `prt-install(8)`, 
`prt-loc(8)`, `prt-prov(8)`, `prt-pull(8)`, `prt-sysup(8)`, `prt-uninstall(8), prt-get(8)`,
`pkgmk(8)`, `pkgrm(8)`, `pkgadd(8)`, `ports(8)`, `pkginfo(8)`, `prt-utils(1)`
