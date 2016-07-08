prtstuff
========

Consitent CRUX port utilities writtin in fish, aiming to replace prt-get, ports, and pkgutils.


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
  -t,   --tree            list using tree view
  -h,   --help            print help and exit
```


depmk
=====

Install dependencies recursivly.


Usage
-----

By changing `set makecommand` in `/etc/prtstuff/config` you can use an `pkgmk` alternative. I don't
know about any, but I want to rewrite `pkgmk` in fish in the future.

Set the `set readme` and `set script` to either `false` or `true` to change the default behavoir
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
------

Prints port location.


Help
----

```
Usage: prtloc [options] [ports]

options:
  -d,   --duplicate       list duplicate ports as well
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


prtls
------

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


prtpatch
--------

Patches ports.


Usage
-----

TODO


Help
----

```
Usage: prtpatch [ports]

options:
  -h,   --help            print help and exit
```


prtprint
--------

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


prtprovide
----------

Search ports for files they provide.


Help
----

```
Usage: prtprovide [options] [queries]

options:
  -h,   --help            print help and exit
```


prtpull
-------

Pull in ports using git.


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

If use use CRUX (you probably do) you can also install using this port: https://github.com/6c37/crux-ports-git/tree/3.2/prtstuff


Notes
-----

Most of the script only workig in a directory with a `Pkgfile`, just like `pkgmk`.

`prtstuff` ships with a fish function named `cdp`, which cds to a specified port directory.
It uses `prtloc`, so comes with ordering, and aliasing.
