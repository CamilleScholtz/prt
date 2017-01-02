# prt-depends 8 "2017-01-02" prt "General Commands Manual"

## NAME

prt-depends - list dependencies recursively


## SYNOPSIS

prt depends [arguments]


## DESCRIPTION

prt depends lists dependencies recursively, it does so by finding the location of a port mentioned in the
`Depends on:` comment in a Pkfile. if there are multiple ports found (for example `foo/bar` and `baz/bar`)
it will choose the most "important" repo using the `prt.toml(5)` order value.


## OPTIONS

`-a` also list installed dependencies. Without this option only ports that are not yet
installed get listed. Same as `--all`

`-n` disable aliasing. Without this option ports get aliased using values found in `prt.toml(5)`. Same as `--no-alias`

`-t` list using tree view. Same as `--tree`


## EXAMPLES

List all non-installed dependencies:

```
$ prt depends
opt/mplayer
opt/qt4
opt/libmng
```

List all dependencies in tree view: 

```
$ prt depends -t -a | head -n 5
opt/mplayer
-  opt/expat
-  6c37-dropin/freetype-iu
-  -  core/zlib
-  -  opt/libpng
```

List all dependencies in tree view, do not alias ports:

```
$ prt depends -t -a -n | head -n 5
opt/mplayer
-  opt/expat
-  opt/freetype
-  -  core/zlib
-  -  opt/libpng
```


## AUTHORS

Camille Scholtz


## SEE ALSO

prt(8), prt.toml(5), prt-diff(8), prt-info(8), prt-install(8), prt-list(8), 
prt-loc(8), prt-prov(8), prt-pull(8), prt-sysup(8), prt-uninstall(8), prt-get(8),
pkgmk(8), pkgrm(8), pkgadd(8), ports(8), pkginfo(8), prt-utils(1)
