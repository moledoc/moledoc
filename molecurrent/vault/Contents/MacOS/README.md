# vault

convenience wrapper around `pw`.

## Dependencies

* [pw](https://github.com/moledoc/pw)
* fzf
* BSD `sed`/`grep` to exist in `/usr/bin`

## Intent

- have 2 files in the `/Applications/vault/Contents/MacOS` that need `sudo` access
	- pw key
	- file named `vault.contents` that contains salt, pepper, domain and extra flags in format:
		- `salt pepper domain ("extra args")`; **NOTE:** extra args are optional and surrounded by double-quotes (")
```sh
sudo chown root:wheel <file>
sudo chmod 600 <file>
```

## TODO

* rename `/Applications/vault` to `/Appliations/vault.app`

## Author

Meelis Utt
