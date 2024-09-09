# Reimaging Windows

* Mount Windows ISO; it will be mounted as disk `D:`.
* Connect USB-drive; in this doc, it will be mounted as disk `E:`.
* Open PowerShell with Administrative privileges and run `diskpart`; Note: `clean` frees the whole USB.

```txt
DISKPART> list disk
DISKPART> select disk <nr for USB>
DISKPART> clean
DISKPART> create partition primary size=29336
DISKPART> list partition
DISKPART> list volume
DISKPART> select volume <nr for created partition>
DISKPART> format fs=fat32 quick
DISKPART> active
DISKPART> exit
```

* [This microsoft doc](https://learn.microsoft.com/en-us/windows-hardware/manufacture/desktop/install-windows-from-a-usb-flash-drive?view=windows-11) was used.
More specifically, the section [If your Windows image is larger than 4GB](https://learn.microsoft.com/en-us/windows-hardware/manufacture/desktop/install-windows-from-a-usb-flash-drive?view=windows-11#if-your-windows-image-is-larger-than-4gb).

```sh
robocopy D: E:  /s /max:3800000000 
Dism /Split-Image /ImageFile:D:\sources\install.wim /SWMFile:E:\sources\install.swm /FileSize:3800
```

* Unmount/Eject ISO and the USB-drive.
* shutdown computer
* start computer and hit escape to go to BIOS menu
* go to Boot Order and select USB drive
* enjoy reimaging and occasionally click around for some setup

## Author

Meelis Utt
