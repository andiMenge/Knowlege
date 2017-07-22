> Workaround for USB 3.0 Disks that do not work with centos 7 out of the box. In my case the journal showed frequend device disconnects and errors related to the uas module.
> The step `ignore UAS for the device` could be work without gobally backlisting UAS (not tested).
> UAS = `USB Attached SCSI`

# Guide

1. Backlist UAS kernel module
2. Ignore UAS for the device

# Backlist UAS kernel module

create `blacklist-uas.conf` in `/etc/modprobe.d/`

blacklist-uas.conf
```
blacklist uas
```

# Tell usb-storage module to ignore UAS for the device
create `/etc/modprobe.d/ignore_uas.conf` with this content:

```
options usb-storage quirks=Vendor_ID:Product_ID:u
```

replace `Vendor_ID` and `Product_ID` with the values from your disk.

You find them with `lsusb -v` command.

**LaCie Rugged FW USB3 example**
```
options usb-storage quirks=0x059f:0x104b:u
```

[source](https://bbs.archlinux.org/viewtopic.php?pid=1428782#p1428782)