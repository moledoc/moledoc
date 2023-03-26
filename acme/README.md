Acme
====

is a text editor Rob Pike wrote and Russ Cox ported to unix. It has
really cool ideas and I wanted to explore them. Russ Cox has made a
really good tour of [acme](https://research.swtch.com/acme) to get
started. This readme is to document some helpful commands, cool stuff
and some general notes/tips/help. It will follow rather freeflow
formatting.

Notes
-----

* line number: `Edit =`
* cursor position: `Edit =#`
* replace in the whole file: `Edit ,s/<from>/<to>/g`
* replace between lines: `Edit 10,11s/<from>/<to>/g`
* comment out full line: `Edit s/^(.*)$/\/\/ \1/g`
* comment in full line: `Edit s/^\/\/ (.*)$/\1/g`
* comment out selection: `Edit s/(.*)/\/\/ \1/g`
* comment in selection: `Edit s/\/\/ (.*)/\1/g`
* indent left: `Edit s/ (.*)/\1/g`
* indent right: `Edit s/(.*)/ \1/g`
* incremental change: `Edit .+#0/old/c/new`
* run current go file: `go run $%`
* re-run given command: `Watch go run $%` (Watch:https://github.com/9fans/go/tree/main/acme/Watch or https://github.com/eaburns/Watch)
* have goimports and gofmt support in acme:	run `acmego` from https://github.com/9fans/go/acme/acmego
* right click to go-to line: `:n`
* to use tab in: `Edit literally type tab (not \t)`

------------------------------------------------------------------------

-   index code - tool by Ross Cox

``` {.sh}
go install github.com/google/codesearch/cmd/...@latest
```

``` {.sh}
cindex <list of go dirs to index>
```

------------------------------------------------------------------------

-   Scripts using `csearch` and grep

``` {.sh}
#!/bin/sh

9 grep -i -n '^func (\([^)]+\) )?'$1'\(' *.go /dev/null
```

``` {.sh}
#!/bin/sh

csearch -n -f '\.go$' '^func (\([^)]+\) )?'$1'\('
```

``` {.sh}
#!/bin/sh

9 grep -i -n '^type '$1' ' *.go /dev/null
```

``` {.sh}
#!/bin/sh

csearch -n -f '\.go$' '^type '$1
```

MOUNTING AND FONTS
------------------

Starting with the command that worked for me, then Russ Cox's comments I
found and then some things I tried previously.

### Running Acme

``` {.sh}
./bin/acme -f /mnt/font/GoRegular/14a/font -m /mnt/acme
```

### From Russ

```{.txt}
    plan9port works well with FUSE.
    It works less well with the 9p module.
    Assuming you have write permission on /mnt/acme
    and FUSE installed,

    acme -m /mnt/acme

    should work just fine.
```

------------------------------------------------------------------------

```{.txt}
    For what it's worth, you don't need to mount fontsrv anywhere.
    You can just use 'fontsrv -p .' to list the fonts,
    and then refer to them as /mnt/font/Name/SIZEa/font.
    For example:

    /usr/local/plan9/bin/acme \
        -f /mnt/font/LucidaGrande/12a/font \
        -F /mnt/font/SourceCodePro-Regular/12a/font

    In that form, acme will "open" /mnt/font/... by reading fontsrv -p's output.
```

### Stuff that I also found/saw

```{.txt}
    mount acme to /mnt
    9 mount `namespace`/acme /mnt/acme

    start font service:
    fontsrv -m /mnt/font

    change font inside acme
    Font <path to font>

    allow acme to mount with ease:
    chmod 777 /mnt/acme
    chmod 777 /mnt/font
```

Good resources
--------------

-   tour of acme: https://research.swtch.com/acme
-   acme man page: https://9fans.github.io/plan9port/man/man1/acme.html
-   using acme editor:
    https://groups.google.com/g/comp.os.plan9/c/\_YUEVbTFuME/m/tJHB8y8-0vYJ
-   Sam editor (to read on Edit command):
    https://9p.io/sys/doc/sam/sam.html
-   fontsrv: https://9fans.github.io/plan9port/man/man4/fontsrv.html
-   some font stuff I found:
    https://9fans.topicbox.com/groups/9fans/Td0ab6c3112c95493-M4005dd63b8324e8b0133f10d

Collected by
------------

Meelis Utt
