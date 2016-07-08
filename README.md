prtstuff
========

Consitent CRUX port utilities written in `fish`, aiming to replace, or at least, be used in combination with `prt-get`, `ports`, some `pkgutils` and maybe `pkgmk` in the future.

You might ask why I'm rewriting all these utils that work perfectly fine? One reason if for fun, a few others are:

* I'm kind of a perfectionst, I want all my terminal programs to have the exact same style of output.
  all the `--help` outputs of the `prtstuff` utils use the same kind of spacing, identation is
  always done with a black arrow (`->`), for example in `depls`, `prtpull`, `prtprovide`, etcetera.
  All utils use the same colors, same kind of flags, etcetera.

* Another inconsitency is how `pkgmk` only works in a directory with a `Pkgfile`, but `prt-get` is
  the other way around and only works by providing a port name. This has always really bugged me.
  I definitely like they way `pkgmk` does it, so almost all `prtstuff` utils work this way.
  In combination with `cdp` it makes managing ports a breeze.

* `prtstuff` uses one config file that sets ordering, aliasing, colors, and more for all `prtstuff` utils.

* `prtstuff` tries to follow the UNIX philosophy of doing one thing and doing it well. `prtpull` ONLY pulls in new ports,
  `prtls` ONLY lists repos or ports, `depls` ONLY lists dependencies, etcetera.

*  With `fish` being my main shell, and `prtutils` being written in fish, I could add a few nifty things:
   for example completions, and a function called `cdp` that uses `prtloc` to cd to ports, for example
   `cdp mpv` cds to `/usr/port/6c37-git/mpv`.
   
   


----


depls
=====

List dependencies recursively.


Help
----

```
Usage: depls [options]

options:
  -a,   --all             also list installed dependencies
  -n,   --no-alias        disable aliasing
  -t,   --tree            list using tree view
  -h,   --help            print help and exit
```


Examples
--------

List all uninstalled dependencies:
```
$ depls
opt/mplayer
opt/qt4
opt/libmng
```

List all uninstalled dependencies in tree view:
```
$ depls -t
opt/mplayer
opt/qt4
-> opt/libmng
```

List all dependencies without aliasing in tree view:
```
depls -tna
opt/mplayer
-> opt/expat
-> opt/freetype
-> -> core/zlib
-> -> opt/libpng
...
```


depmk
=====

Install dependencies recursivly.


Usage
-----

By changing `set makecommand` in `/etc/prtstuff/config` you can use a `pkgmk` alternative. I don't
know about any, but I want to rewrite `pkgmk` in fish in the future.

Set the `set readme` and `set script` to either `true` or `false` to change the default behavoir
of `depmk` in `/etc/prtstuff/config`, you can toggle these values using the `-s` and `-r` flags.


Help
----

```
Usage: depmk [options]

options:
  -s,   --script          toggle execution of pre- and post-install
  -r,   --readme          toggle opening of readmes
  -h,   --help            print help and exit
```


prtloc
======

Prints port location.


Help
----

```
Usage: prtloc [options] [ports]

options:
  -d,   --duplicate       list duplicate ports as well
  -n,   --no-alias        disable aliasing
  -h,   --help            print help and exit
```


Examples
--------

List the location all installed ports:
```
$ prtloc (prtls -i | cut -d ' ' -f 1)
opt/alsa-lib
opt/alsa-plugins
opt/alsa-utils
opt/aspell
opt/aspell-en
...
```

List duplicate ports in the order they are used:
```
$ prtloc -d openbox mpv
punpun/openbox
-> 6c37-git/openbox
-> -> opt/openbox
6c37-git/mpv
-> contrib/mpv
```

Like most other utils, `prtloc` does aliasing, however, this can be disabled with the `-n` flag:
```
$ prtloc openssl
6c37/libressl
$ prtloc -n openssl
core/openssl
```


prtls
=====

List repos and ports.


Help
----

```
Usage: prtls [options]

options:
  -r,   --repos           list repos
  -i,   --installed       list installed ports
  -h,   --help            print help and exit
```


Examples
--------

List all ports in the ports tree:
```
$ prtls
6c37/abduco
6c37/arandr
6c37/atari800
6c37/atool
6c37/audacity
...
```

List all installed ports:
```
$ prtls -i
alsa-lib 1.1.0-1
alsa-plugins 1.1.1-1
alsa-utils 1.1.0-2
aspell 0.60.6.1-1
aspell-en 2016.06.26-0-1
atk 2.20.0-1
...
```


prtpatch
========

Patches ports.


Usage
-----

`prtpatch` uses files in `/etc/prtstuff/patch` to get information about what ports to path.
Here is an example of how to patch `opt/libpcre2` to add a configure flag:
First create the path in `/etc/prtstuff/patch`, in this case that will be `opt/libpcre2` (so `/etc/prtstuff/patch/opt/libpcre2`).
Secondly create the a `Pkgfile.patch` file with the following content:

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

And now run `prtpatch`, which will do all the patching itself.

All files in the patch directory ending with the `.patch` filetype will be used by `prtpatch`,
say you want to patch `.footprint` you would create a `.footprint.patch` file.


Help
----

```
Usage: prtpatch [ports]

options:
  -h,   --help            print help and exit
```


prtprint
========

Prints port information.


Help
----

```
Usage: prtprint [options]

options:
  -d,   --description     print description
  -u,   --url             print url
  -m,   --maintainer      print maintainer
  -v,   --version         print version
  -r,   --release         print release
  -h,   --help            print help and exit
```


Examples
--------

Print everything:
```
$ prtprint
Description: Mplayer frontend
URL: http://smplayer.sf.net/
Maintainer: Alan Mizrahi, alan at mizrahi dot com dot ve
Version: 15.11.0
Release: 1
```

Print only the version and release:
```
$ prtprint -v -r
Version: 15.11.0
Release: 1
```


prtprovide
==========

Search ports for files they provide.


Help
----

```
Usage: prtprovide [options] [queries]

options:
  -h,   --help            print help and exit
```

Examples
--------

Search multiple terms at once for files they provide:
```
$ prtprovide lemonbar.1 n30f
6c37-git/lemonbar-xft
-> /usr/share/man/man1/lemonbar.1.gz
6c37/lemonbar
-> /usr/share/man/man1/lemonbar.1.gz
6c37/n30f
-> /usr/bin/n30f
```


prtpull
=======

Pull in ports using git.


Usage
-----

`prtpull` uses files in `/etc/prtstuff/pull` to get information about what repositories to pull.


Help
----

```
Usage: prtpull [options] [repos]

options:
  -h,   --help            print help and exit
```


Examples
--------

Pull in new ports for all repos:
```
# prtpull
Updating collection 1/7, 6c37.
Updating collection 2/7, 6c37-git.
Updating collection 3/7, contrib.
Updating collection 4/7, core.
Updating collection 5/7, opt.
...
```

Pull in new ports for specified repos:
```
# prtpull punpun core
Updating collection 1/2, punpun.
Updating collection 2/2, core.
```


----


Dependencies
------------

* fish (2.3.0+)
* getopts (https://github.com/fisherman/getopts)


Installation
------------

Run `make install` inside the `prtstuff` directory to install the scripts.
`prtstuff` can be uninstalled easily using `make uninstall`.

Edit `/etc/prtstuff/config` to your liking.

If you use CRUX (you probably do) you can also install using this port: https://github.com/6c37/crux-ports-git/tree/3.2/prtstuff


Notes
-----

Most of the script only workig in a directory with a `Pkgfile`, just like `pkgmk`.

`prtstuff` ships with a fish function named `cdp`, which cds to a specified port directory.
It uses `prtloc`, so comes with ordering, and aliasing.
