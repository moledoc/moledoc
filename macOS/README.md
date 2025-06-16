# macOS

This documents my first look at macOS.

## Settings

So first thing I did was open Settings and explore different sections and behavior.
For now I don't have anything specific to write down.
So when I receive the next mac, I should go through the Settings manually once again.

One note though: in macOS, if the user is an `Administrator` then that doesn't mean they are the root user - it actually means they are in the `sudoers` file so they could invoke `sudo` to elevate privileges.

## Installing software

So on the first look I used MacPorts and followed `Source installation` instructions from [here](https://www.macports.org/install.php).
Before I describe the steps here, it's important to mention, that I also needed to enable developer tools

```sh
xcode-select --install
xcodebuild -license
```

Now back to MacPorts.
Installation from source

```sh
wget https://github.com/macports/macports-base/releases/download/v2.10.1/MacPorts-2.10.1.tar.gz
cd MacPorts-2.10.1
./configure && make && sudo make install
cd .. && rm -rf MacPorts-2.10.1*
sudo port -v selfupdate
```

By default, MacPorts is installed to `/opt` and the binaries installed by MacPorts can be found at `/opt/local/bin`.

The interface for MacPorts is `port`, which is from the first glance similar to `apt`, for example.

## Homebrew

install with 

```sh
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```
Essentially works like `apt`.

## Applications

One thing I investigated was to figure out how to add custom apps to `Spotlight` and Applications view.
The answer is `/Appications` directory.
So essentially an app is just a directory, that contains resources (eg app icon, configuration, etc) and lastly the executable, that can be found in `/Applications` path.
For a minimal example, see [acme.app](../molecurrent/acme.app/).

## Author

Meelis Utt
