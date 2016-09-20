# prt

Consitent CRUX port utility written in fish, aiming to replace, or at least, be used in combination with `prt-get`, `ports`, and some `prt-utils`. These scripts still make use of `pkgmk` and `pkgadd`,
simply because it's too hard/complex to parse `Pkgfile`s (bash) with fish.

You might ask why I'm rewriting all these utils that work perfectly fine? One reason if for fun, a few other things that make `prt` interesting are:

* An inconsitency that really bugs me is how `pkgmk` only works by being in a directory with a `Pkgfile`, but `prt-get` is
  the other way around and only works by providing a port name.
  I definitely like they way `pkgmk` does it, so `prt` works this way, but unlike `pkgmk` also does dependencies and stuff.
  In combination with `cdp` it makes managing ports a breeze.

* I'm kind of a perfectionst, I want all my terminal programs to have the exact same style of output.
  all the `--help` outputs of the prtstuff utils use the same kind of spacing, identation is
  always done with a black arrow (`->`), see the `prt depends`, `prt pull` and `prt provide` output.
  All utils use the same colors, same kind of flags, etcetera.

  The command line is kinda inspired by the `go` one.

* prtstuff uses one config file that sets ordering, aliasing, colors, and more for all `prt` utils.

* None of the prtstuff utils depend on `prt`, kinda neat huh, it's slower though, but whatever.

* `prt` has nice fish integration, for example a function named `cdp` that uses `prt location` to cd to ports, for example
  `cdp mpv` cds to `/usr/port/6c37-git/mpv`. Most commands also have some form of completion.


----

## cdp

cd to port location, using `prt location`.

### Examples

cd to port locations:
```
$ cdp mpv
$ pwd
/usr/ports/6c37-git/mpv
$ cdp
$ pwd
/usr/ports
$ cdp openssl
$ pwd
/usr/ports/6c37/libressl
```


## prt depends

List dependencies recursively.

### Help

```
Usage: prt depends [arguments]

arguments:
  -a,   --all             also list installed dependencies
  -n,   --no-alias        disable aliasing
  -t,   --tree            list using tree view
  -h,   --help            print help and exit
```

### Examples

List all not-yet-installed dependencies:
```
$ prt depends
opt/mplayer
opt/qt4
opt/libmng
```

List all not-yet-installed dependencies in tree view:
```
$ prt depends -t
opt/mplayer
opt/qt4
-> opt/libmng
```

List all dependencies without aliasing them in tree view:
```
$ prt depends -tna
opt/mplayer
-> opt/expat
-> opt/freetype
-> -> core/zlib
-> -> opt/libpng
...
```


## prt diff

List outdated packages.

### Help

```
Usage: prt diff [arguments]

arguments:
  -v,   --version         list installed and available version
  -h,   --help            print help and exit
```


## locprt

Prints port location.


### Help

```
Usage: locprt [options] [ports]

options:
  -d,   --duplicate       list duplicate ports as well
  -n,   --no-alias        disable aliasing
  -h,   --help            print help and exit
```


### Examples

List the location of all installed ports:
```
$ locprt (lsprt -i | cut -d ' ' -f 1)
opt/alsa-lib
opt/alsa-plugins
opt/alsa-utils
opt/aspell
opt/aspell-en
...
```

List eventual duplicate ports in the order they are used:
```
$ locprt -d openbox mpv
punpun/openbox
-> 6c37-git/openbox
-> -> opt/openbox
6c37-git/mpv
-> contrib/mpv
```

Like most other utils, `locprt` does aliasing, however, this can be disabled with the `-n` flag:
```
$ locprt openssl
6c37/libressl
$ locprt -n openssl
core/openssl
```




## lsprt

List repos and ports.


### Help

```
Usage: lsprt [options]

options:
  -r,   --repos           list repos
  -i,   --installed       list installed ports
  -h,   --help            print help and exit
```


### Examples

List all ports in the ports tree:
```
$ lsprt
6c37/abduco
6c37/arandr
6c37/atari800
6c37/atool
6c37/audacity
...
```

List all installed ports:
```
$ lsprt -i
alsa-lib 1.1.0-1
alsa-plugins 1.1.1-1
alsa-utils 1.1.0-2
aspell 0.60.6.1-1
aspell-en 2016.06.26-0-1
atk 2.20.0-1
...
```

## mkdiff

Update ports that get listed by `lsdiff`.


### Usage

See `mkdep` usage.


### Help

```
Usage: mkdiff [options]

options:
  -s,   --script          toggle execution of pre- and post-install
  -v,   --verbose         enable verbose output
  -h,   --help            print help and exit
```


## mkprt

Install or update ports.


### Usage

See `mkdep` usage.


### Help

```
Usage: mkprt [options]

options:
  -s,   --script          toggle execution of pre- and post-install
  -v,   --verbose         enable verbose output
  -h,   --help            print help and exit
```


## patchprt

Patches ports.


### Usage

`patchprt` uses files in `/etc/prtstuff/patch` to get information about what ports to patch.
Here is an example of how to patch `opt/libpcre2` to add a configure flag:
first create the path in `/etc/prtstuff/patch`, in this case that will be `opt/libpcre2` (so `/etc/prtstuff/patch/opt/libpcre2`).
Secondly create a `Pkgfile.patch` file with the following content:

```diff
--- Pkgfile	2016-03-20 02:01:46.054976416 +0100
+++ new	2016-03-20 02:02:52.534979140 +0100
@@ -13,7 +13,8 @@
     cd pcre2-$version
 
     ./configure --prefix=/usr \
-                --enable-jit
+                --enable-jit \
+                --enable-pcre2-32
 
     make
     make DESTDIR=$PKG install
```

And now run `patchprt`, which will do all the patching.

Only files in the patch directory ending with a `.patch` filetype will be used by `patchprt`,
say you want to patch `.footprint` you would create a `.footprint.patch` file.


### Help

```
Usage: patchprt [ports]

options:
  -h,   --help            print help and exit
```


## printprt

Prints port information.


### Help

```
Usage: printprt [options]

options:
  -d,   --description     print description
  -u,   --url             print url
  -m,   --maintainer      print maintainer
  -v,   --version         print version
  -r,   --release         print release
  -h,   --help            print help and exit
```


### Examples

Print all port information:
```
$ printprt
Description: Mplayer frontend
URL: http://smplayer.sf.net/
Maintainer: Alan Mizrahi, alan at mizrahi dot com dot ve
Version: 15.11.0
Release: 1
```

Print only the version and release:
```
$ printprt -v -r
Version: 15.11.0
Release: 1
```


## provprt

Search ports for files they provide.


### Help

```
Usage: provprt [options] [queries]

options:
  -h,   --help            print help and exit
```

### Examples

Search multiple terms at once for files they provide:
```
$ provprt lemonbar.1 n30f
6c37-git/lemonbar-xft
-> /usr/share/man/man1/lemonbar.1.gz
6c37/lemonbar
-> /usr/share/man/man1/lemonbar.1.gz
6c37/n30f
-> /usr/bin/n30f
```


## pullprt

Pull in ports using git.


### Usage

`pullprt` uses files in `/etc/prtstuff/pull` to get information about what repositories to pull.


### Help

```
Usage: pullprt [options] [repos]

options:
  -h,   --help            print help and exit
```


### Examples

Pull in new ports for all repos:
```
# pullprt
Updating collection 1/7, 6c37.
Updating collection 2/7, 6c37-git.
-> Modifying mpv/Pkgbuild
Updating collection 3/7, contrib.
Updating collection 4/7, core.
...
```

Pull in new ports for specified repos:
```
# pullprt punpun core
Updating collection 1/2, punpun.
Updating collection 2/2, core.
```


----


## Dependencies

* fish (2.3.0+)
* getopts (https://github.com/fisherman/getopts)
* pkgutils


## Installation

Run `make install` inside the `prtstuff` directory to install the scripts.
prtstuff can be uninstalled easily using `make uninstall`.

Edit `/etc/prtstuff/config` to your liking.

If you use CRUX (you probably do) you can also install using this port: https://github.com/onodera-punpun/crux-ports-git/tree/3.2/prtstuff
