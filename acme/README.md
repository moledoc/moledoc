# acme

`acme` is a text editor written by Rob Pike for Plan9 in the early 1990s and ported to Unix by Ross Cox.
I highly recommend watching this [tour of acme](https://research.swtch.com/acme) to get started.
To get more details about the editor, I recommend checking out links in the [Resources](#Resources) section or rest of this doc.

## Installation

```sh
sudo apt install build-essential # or equivalent to the package manager
cd /usr/local
git clone --depth=1 https://github.com/9fans/plan9port.git plan9
sudo chown -R <user>:<user> plan9
cd plan9
./INSTALL
```

Add the following lines to your shell of choice

```sh
export PLAN9=$HOME/plan9
export PATH=$PATH:$PLAN9/bin
```

### lsp

To have lsp support, install `acme-lsp` and `L` from
* acme-lsp (GO111MODULE=on go install github.com/fhs/acme-lsp/cmd/acme-lsp@v0.11.0)
* L (GO111MODULE=on go install github.com/fhs/acme-lsp/cmd/L@v0.11.0)

I have an example of `acme-lsp` config file that can be found [here](../molecurrent/acme-lsp-config.toml) - that contains configuration for go with `gopls` and c/c++ with `ccls`.

## Fonts

TODO:

## Plumber

TODO:

## Sam commands


* selection (**NOTE:** right click):
	* beginning of file `:0`
	* end of file `:$`
	* range `:<lower>,<upper>` (highlights); eg
		* full file: `:,`/`:0,$`
		* from lower to end: `:<lower>,`/`:<lower>,$`
		* from beginning to upper: `:,<upper>`/`:0,<upper>`
		* partial: `:<lower>,<upper>`, where `lower in [1, upper)` and `upper in (lower, $)`
	* search `:[+-]/<regex>`, where `+`/`-` is optional (default is `+`). `+` is forwared and `-` is backwards search.
	* regex search `:/<regex_start>/,/<regex_end>/`
* to edit the file, use `Edit` - it works with selection (see above) and has the following action options (**NOTE:** middle click):
	* **NOTE:** to use tab in `Edit` literally type tab  and not `\t`
	* range `Edit <lower>,<upper>`
	* search `Edit [+-]/<regex>`
	* line number `Edit =`
	* cursor position `Edit =#`
	* replace `s` `Edit <lower>,<upper>s/<old>/<new>/[g]`
	* insert `i` `Edit <lower>,<upper>i/<text>/`
	* append `a` `Edit <lower>,<upper>a/<text>/`
	* change `c` `Edit <lower>,<upper>c/<text>/`
	* delete `d` `Edit <lower>,<upper>d`
	* for each match `x` `Edit <lower>,<upper>x/<match>/<action>/`, where action is one of the above
	* for each unmatch `y` `Edit <lower>,<upper>y/<match>/<action>/`, where action is one of the above
	* for each matching filename `X` `Edit X/<match>/<action>/
	* Braces group commands `{}`; eg
		* 
```txt
Edit ,x/Acme/ {
  i/I like 
  a/ editor
}
```
	* incremental change `Edit .+#0/<old>/c/<new>`
	* there is more, but I don't use them too often, so check Sam's reference (**NOTE:** it's not 100% match, so try out what interests you).
* Piping:
	* input from acme to external, output is not redirected to acme: `>`
	* input from external to acme, input is not coming from acme: `<`
	* input from acme to external, output is redirected to acme: `|`
	* piping works also with `Edit`; eg
		* `Edit ,x/HERE/ < date`

## Helpful Sam commands

* comment out full line: `Edit s/^(.*)$/\/\/ \1/g`
* comment in full line: `Edit s/^\/\/ (.*)$/\1/g`
* comment out selection: `Edit s/(.*)/\/\/ \1/g`
* comment in selection: `Edit s/\/\/ (.*)/\1/g`
* indent left: `Edit s/ (.*)/\1/g`
* indent right: `Edit s/(.*)/ \1/g`
* remove ansi color chars: `Edit ,s/\[[0-9]*m//g`
* re-run given command: `Watch go run $%` ([Watch](https://github.com/9fans/go/tree/main/acme/Watch))
* run current go file: `go run $%`

## Mouse/Keyboard

* Mouse:
	* left - select
	* right - search/plumb
	* middle - execute
	* left+middle - cut
	* left+right - paste
	* left+middle+right - cut+paste
	* middle+left - execute command with highlighted selection
	* middle+right - cancel execute
	* right+middle - cancel search
* Keyboard:
	* `ctrl-a` move cursor to start of the line
	* `ctrl-e` move cursor to end of the line
	* `ctrl-f` filepath autocompletion
	* `ctrl-h` delete character before the cursor
	* `ctrl-i` tab
	* `ctrl-j` enter
	* `ctrl-u` delete from cursor to start of line
	* `ctrl-w` delete word before the cursor
	* macOS had copy/paste/cut/undo/redo, so I have a patch to enable them on linux; keybindings are as follow:
		* `ctrl-b` - copy
		* `ctrl-v` - paste
		* `ctrl-x` - cut
		* `ctrl-z` - undo
		* `ctrl-y` - redo
* misc:
	* window id: `echo $winid`
	* filename: `echo $%`/`echo $samfile`
	* append to the end of current window: `echo some text | 9p write acme/$winid/body` 
	* Esc to select last text; double-Esc to delete lastly selected text
	* switch between fonts `Font <path to font>`
	* write state of acme `Dump`, written to $HOME/acme.dump
	* load state of acme `Load`

## IDE - integrated development environment

# TODO:

## Resources

* [plan9port](https://github.com/9fans/plan9port.git)
* [tour of acme](https://research.swtch.com/acme)
* [acme man page](https://9fans.github.io/plan9port/man/man1/acme.html)
* [using acme editor](https://groups.google.com/g/comp.os.plan9/c/\_YUEVbTFuME/m/tJHB8y8-0vYJ)
* [Sam editor (to read on Edit command)](https://9p.io/sys/doc/sam/sam.html)
* [fontsrv](https://9fans.github.io/plan9port/man/man4/fontsrv.html)
* [some font stuff I found](https://9fans.topicbox.com/groups/9fans/Td0ab6c3112c95493-M4005dd63b8324e8b0133f10d)
* [acme mouse chords](http://acme.cat-v.org/mouse)

## Author

Meelis Utt
