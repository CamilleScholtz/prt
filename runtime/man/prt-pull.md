# prt-pull 8 "2017-01-02" prt "General Commands Manual"

## NAME

prt-pull - pull in ports


## SYNOPSIS

prt pull [arguments] [repos]


## DESCRIPTION

prt pull is pretty much `ports -u`, but it has the ability to pull to a location that is
not `/usr/ports`, it is also possible to only pull in certain repos by specifying them.


## EXAMPLES

Only pull in punpun and core repos:

```
# prt pull punpun core
Pulling in repo 1/2, core.
Pulling in repo 2/2, punpun.
```


## AUTHORS

Camille Scholtz


## SEE ALSO

`cdp(1)`, `prt(8)`, `prt.toml(5)`, `prt-depends(8)`, `prt-diff(8)`, `prt-info(8)`, `prt-install(8)`, 
`prt-list(8)`, `prt-loc(8)`, `prt-prov(8)`, `prt-sysup(8)`, `prt-uninstall(8), prt-get(8)`,
`pkgmk(8)`, `pkgrm(8)`, `pkgadd(8)`, `ports(8)`, `pkginfo(8)`, `prt-utils(1)`
