# Manage Kernel Modules
Blacklisted and disabled modules
>`modprobe –showconfig | egrep “^(blacklist|install)”`

Find modules
>`find /lib/modules/`uname -r` -print`

Show loaded modules
>`lsmod`

Load module
>`modprobe module`

Unload module
>`modprobe -r module`

Module details
>`modinfo module`

# Blacklist Modules
> prohibit loading of specific kernel modules

create file in `/etc/modprobe.d/`

blacklist-uas.conf
```
blacklist uas
```