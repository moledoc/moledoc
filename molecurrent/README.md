# molecurrent

There used to be/is a repo called [molecurrent](https://github.com/moledoc/molecurrent).
That repo has been used to automate linux setups - and it works, but is a bit specific.
Furthermore, if the environment is different, I've found myself somewhat often just copy-pasting commands from the setup script and making sure I have everything I need.
So, this document and directory is an attempt to make setting up more manual, descriptive and less specific. Also just collect programs I want/need and some configuration/installation guides.

## Programs

NOTE: some programs might have installation/setup steps defined below.

NOTE: when building from source, use `git clone --depth=1` for quicker cloning.

* package manager (install or update); eg
	* apt update && apt upgrade
	* port ([macports](https://www.macports.org))
	* brew ([homebrew](https://brew.sh))
* git
* devel packages
	* some used in debian 12 setups: build-essential libx11-dev libxext-dev libxt-dev libfontconfig1-dev libxtst-dev libxinerama-dev libxft-dev libxrandr-dev
	* macOS: `xcode-select --install` and `xcodebuild -license`
* clang/gcc
* wget
* [plan9port](https://github.com/moledoc/plan9port)
	* acme
* [go](https://go.dev/dl/)
	* acme-lsp (GO111MODULE=on go install github.com/fhs/acme-lsp/cmd/acme-lsp@latest)
	* L (GO111MODULE=on go install github.com/fhs/acme-lsp/cmd/L@latest)
	* gopls (GO111MODULE=on go install golang.org/x/tools/gopls@latest)
* [9fans/acme/Watch](https://github.com/9fans/go.git 9fans)
	* `cd ./9fans/acme/Watch; go install -buildvcs=false; cd-`
* ccls
* [pw](https://github.com/moledoc/pw/releases/tag/v0.7.1)
	* pwcli
	* pwgui
* 9wm/9menu (linux build)
	* [9wm](https://github.com/9wm/9wm.git)
	* [9menu](https://github.com/arnoldrobbins/9menu.git)
* python
	* venv: `python3 -m venv $HOME/venv`
	* enable/disable: `source $HOME/venv/bin/activate` / `deactivate`
	* lsp setup: `$HOME/venv/bin/pip install "python-lsp-server[all]"`

### Debian 12 apt installs

```sh
apt update && apt upgrade
apt install --fix-missing -y xorg xterm build-essential libx11-dev libxext-dev libxt-dev libfontconfig1-dev libxtst-dev libxinerama-dev libxft-dev xdotool libxrandr-dev xautolock xsecurelock xinput xclip parallel doas rfkill curl network-manager git chromium sxhkd fuse3 ntfs-3g alsa-utils vlc dunst spacefm-gtk3 feh flameshot okular fzf ccls
```

## Configurations/Installations

### git

```sh
git config --global init.defaultBranch "main"
git config --global pager.log false
git config --global color.ui false
git config --global core.editor "ed"
git config --global core.pager "cat"
git config --global push.autoSetupRemote "true"
# git config --global commit.gpgsign true

git config --global user.name "<username>"
git config --global user.email "<email>"
```

### plan9port

I'm maintaining a fork of plan9port (mainly for acme), where I've added some 'modern developer' friendly features.

```sh
git clone https://github.com/moledoc/plan9port.git "$HOME/plan9port"
# linux - sudo chown -R <user>:<user> plan9port
# mac - sudo chown -R <user>:<group> plan9port
cd plan9port
./INSTALL
```

Add these to `PATH`

```text
export PLAN9=/usr/local/plan9
export PATH=$PATH:$PLAN9/bin
```

To rebuild a program in plan9port (using acme as an example):

```sh
cd $PLAN9/src/cmd/acme
mk install
mk -a install # to force install all
```

### python

Setup user specific venv and install lsp.


### apt

No color apt

```sh
sudo su
printf 'Binary::apt::DPkg::Progress-Fancy "false";\nBinary::apt::APT::Color "false";' > /etc/apt/apt.conf.d/99nocolor
```

### ssh key generation

```sh
mkdir $HOME/.ssh
ssh-keygen -t rsa -b 4096 -C "<email>" -f $HOME/.ssh/<key_name> -P ""
```

### 9wm

I added alt-tab functionality and patched `9wm`.

```sh
git apply --ignore-whitespace --ignore-space-change 9wm.patch
```

### bluetooth

https://www.makeuseof.com/manage-bluetooth-linux-with-bluetoothctl/

```sh
doas systemctl status bluetooth
doas systemctl enable bluetooth
bluetoothctl scan on # to find the wanted bluetooth device
bluetoothctl discoverable on
bluetoothctl pair <device address> # for new device
bluetoothctl paired-devices # list paired devices
bluetoothctl connect <device address> # for existing device
bluetoothctl (un)trust <device address>
bluetoothctl disconnect <device address>
bluetoothctl remove <device address>
bluetoothctl block <device address>
```

### macOS Spotlight

Applications runnable using Spotlight live in `/Applications` and are essentially dirs with specific structure.
For working examples, see [acme.app](./acme.app) or [vault.app](./vault.app).

### macPorts

Get tar from [Source installationa](https://www.macports.org/install.php) and

```sh
tar xvf <tar.gz>
# chown untarred dir
cd <untarred dir>
./configure && make && sudo make install
cd ../
rm <tar>
sudo port -v selfupdate
```

## Author

Meelis Utt
