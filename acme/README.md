# acme

`acme` is a text editor written by Rob Pike for Plan9 in the early 1990s and ported to Unix by Ross Cox.
I highly recommend watching this [tour of acme](https://research.swtch.com/acme) to get started.
To get more details about the editor, I recommend checking out links in the [Resources](#Resources) section or rest of this doc.

## Appeal

* simple, yet powerful interface
* enables me to get better at my system 
* enables me to gain more context for systems I work in
* (blazingly fast)

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

* `acme-lsp`
```sh
GO111MODULE=on go install github.com/fhs/acme-lsp/cmd/acme-lsp@v0.11.0)
```
* `L`
```sh
GO111MODULE=on go install github.com/fhs/acme-lsp/cmd/L@v0.11.0)
```

I have an example of `acme-lsp` config file that can be found [here](../molecurrent/acme-lsp-config.toml) - that contains configuration for go with `gopls` and c/c++ with `ccls`.

To run lsp

```sh
ACME_LSP_CONFIG=$HOME/.config/acme-lsp/config.toml acme-lsp -hidediag
```

Interfacing is done through `L` command, which has quite good help.
I've seen people creating shell scripts for convenience, but I've stuck with the "non-convenient" approach.

## Fonts

Run `fontsrv` to enable font service.
To list available fonts, run

```sh
9p ls font
```

To run acme with a specific font

```sh
acme -f /mnt/font/<font_name>/<size>a/font
```

I like to use font `/mnt/font/DejaVuSansMono/14a/font`, at least on my linux machine.

Font service doesn't have to run, so no need to start it during startup.

## Plumber

The plumber is a user-level file server that receives, examines, rewrites, and dispatches messages between programs.
It allows to make the plain text editor, that acme is, into very powerful editor.
plan9port comes with sane plumbing configuration.
As such, I've had no need to dig into it too deeply, but sometime in the future I'd like to get more familiar with it.

**NOTE:** During startup, start `plumber` in a background process.

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
	* for each matching filename `X` `Edit X/<match>/<action>/`
	* Braces group commands `{}`; eg
	```txt
	Edit ,x/Acme/ {
	  i/I like 
	  a/ editor
	}
	```
	* incremental change `Edit .+#0/<old>/c/<new>`
	* there is more, but I don't use them too often, so check Sam's reference (**NOTE:** it's not 100% match, so try out what interests you)
* Piping:
	* input from acme to external, output is not redirected to acme: `>`
	* input from external to acme, input is not coming from acme: `<`
	* input from acme to external, output is redirected to acme: `|`
	* piping works also with `Edit`; eg
		* `Edit ,x/HERE/ < date`

### Helpful commands

* comment out full line: `Edit s/^(.*)$/\/\/ \1/g`
* comment in full line: `Edit s/^\/\/ (.*)$/\1/g`
* comment out selection: `Edit s/(.*)/\/\/ \1/g`
* comment in selection: `Edit s/\/\/ (.*)/\1/g`
* indent left: `Edit s/ (.*)/\1/g`
* indent right: `Edit s/(.*)/ \1/g`
* remove ansi color chars: `Edit ,s/\[[0-9]*m//g`

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
	* left click on progress bar: move up by <'page size' * 'where clicked on bar'>
	* right click on progress bar: move down by <'page size' * 'where clicked on bar'>
	* middle click on progress bar: move to that spot in file
	* double click on word or right inside quote/parenthesis/brackets/braces to highlight it
	* left click on tag-box (left-hand in tag, right above scrollbar): increase window
	* right click on tag-box (left-hand in tag, right above scrollbar): make window only window in that column, hiding others
	* middle click on tag-box (left-hand in tag, right above scrollbar): maximize window, don't hide other windows
* Keyboard:
	* `ctrl-a` move cursor to start of the line
	* `ctrl-e` move cursor to end of the line
	* `ctrl-f` filepath autocompletion
	* `ctrl-h` delete character before the cursor
	* `ctrl-i` tab
	* `ctrl-j` enter
	* `ctrl-u` delete from cursor to start of line
	* `ctrl-w` delete word before the cursor
	* macOS had copy/paste/cut/undo/redo, so I have a [patch](../molecurrent/acme_text.patch) to enable them on linux; keybindings are as follow:
		* `ctrl-b` - copy
		* `ctrl-v` - paste
		* `ctrl-x` - cut
		* `ctrl-z` - undo
		* `ctrl-y` - redo
* misc:
	* focus is on the windows where cursor is hovering
	* window id: `echo $winid`
	* filename: `echo $%`/`echo $samfile`
	* append to the end of current window: `echo some text | 9p write acme/$winid/body` 
	* highlight last written text with `Esc`
	* when text is highlighted, `Esc` to delete it
	* `Cut` - cut last highlighted text
	* `Del` - delete window
	* `Delcol` - delete column
	* `Dump` - save acme state
	* `Exit` - exit acme
	* `Font` - change font, `Font <path to font>`
		* if ran in window tag, applies only to the window
		* if in column tag, then to that column
		* if in root tag, then for all windows
	* `Get` - refresh window (eg update file/directory)
	* `Kill` - kill execution
	* `Load` - load acme state
	* `Look` - search
	* `New` - new window
	* `Newcol` - new column
	* `Paste` - paste
	* `Put` - save
	* `Putall` - save all windows
	* `Redo` - redo last action
	* `Snarf` - copy
	* `Sort` - sort windows
	* `Undo` - undo last action
	* `Zerox` - duplicate last active window

## IDE - integrated development environment

Using piping, it's easy to use external programs to extend acme.
Here are some examples:

* re-run given command: `Watch go run $%` ([Watch](https://github.com/9fans/go/tree/main/acme/Watch))
* run current go file: `go run $%`
* `:, | sort`
* the external program `win` that comes with plan9port enables terminal usage inside acme
	* one thing I really enjoy is using `dlv` from acme; some notes why:
		* defining `dlv` commands in the tag allows to have "buttons" for the debugger
			* my `dlv` "buttons": `rebuild r c n s so breakpoints ls p`
		* easily jump to files when setting through code
		* easily rerunning previous commands
	* similar thing to `dlv` applies to `gdb`
* you can access man pages by right clicking on `<man page><section>` - opens a new window with that man page. For example, `acme(4)`, `malloc(3)`. This makes navigating man pages very easy, as you can jump to other man pages easily.
* probably something more, but I'll add it when I remember.

## Resources

* [plan9port](https://github.com/9fans/plan9port.git)
* [tour of acme](https://research.swtch.com/acme)
* [acme man page](https://9fans.github.io/plan9port/man/man1/acme.html)
* [using acme editor](https://groups.google.com/g/comp.os.plan9/c/\_YUEVbTFuME/m/tJHB8y8-0vYJ)
* [Sam editor (to read on Edit command)](https://9p.io/sys/doc/sam/sam.html)
* [fontsrv](https://9fans.github.io/plan9port/man/man4/fontsrv.html)
* [some font stuff I found](https://9fans.topicbox.com/groups/9fans/Td0ab6c3112c95493-M4005dd63b8324e8b0133f10d)
* [acme mouse chords](http://acme.cat-v.org/mouse)
* [plan 9 man pages](https://man.cat-v.org/plan_9/4/)

## Author

Meelis Utt
