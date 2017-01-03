# prt-prov 8 "2017-01-02" prt "General Commands Manual"

## NAME

prt-prov - search ports for files


## SYNOPSIS

prt prov [arguments] [queries]


## DESCRIPTION

prt prov will search through `/var/lib/pkg/db` or through `.footprint`s of ports for files.
It accepts multiple queries, and has full regex support.


## OPTIONS

`-i` search in installed ports only, with this flag prov will search through  `/var/lib/pkg/db`
insted of `.footprint`s of ports. Same as `--installed`


## EXAMPLES

Get all installed files that en with `.md`:

```
$ prt prov -i '\.md$'
elementary-icon-theme
-  usr/share/icons/elementary/README.md
go
-  usr/lib/go/CONTRIBUTING.md
-  usr/lib/go/README.md
-  usr/lib/go/misc/trace/README.md
-  usr/lib/go/src/runtime/HACKING.md
```

Get location of `mpv` and `fish`, list duplicate ports as well:

```
$ prt loc -d mpv fish
6c37-git/mpv
-  contrib/mpv
6c37-git/fish
-  6c37/fish
```


## AUTHORS

Camille Scholtz


## SEE ALSO

`cdp(1)`, `prt(8)`, `prt.toml(5)`, `prt-depends(8)`, `prt-diff(8)`, `prt-info(8)`, `prt-install(8)`, 
`prt-list(8)`, `prt-loc(8)`, `prt-pull(8)`, `prt-sysup(8)`, `prt-uninstall(8), prt-get(8)`,
`pkgmk(8)`, `pkgrm(8)`, `pkgadd(8)`, `ports(8)`, `pkginfo(8)`, `prt-utils(1)`
