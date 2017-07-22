# Create and configure file systems
## Create, mount, unmount, and use vfat, ext4, and xfs file systems
format lv/disk `mkfs.ext4 /dev/mapper/<lv-name>`

mount disk (non persistent) `mount <device> <mount-point>`

persistent mount edit `/etc/fstab` -> `<device> <mount-point> <fs> <options> <dump> <fsck>` -> `/dev/mapper/vg-lv_vol /mnt xfs defaults 1 2`

`fsck` options: 
- `0` means no fsck run at boot (very dangerous)
- `1` fsck is run before root filesystem 
- `2` fsck is run after the root filesystem

`dump` options:
-  `1` for real filesystems
-  `0` for swap and NFS mounted filesystems

run fsck on unmounted device: `fsck <device>`

dump fs infos: `dumpe2fs <device>`

repair xfs: `xfs_repair <device>`

xfs details: `xfs_info <device>`

## Mount and unmount CIFS and NFS network file systems
### NFS
nfs fstab: `nfsserver:/home/tools /mnt nfs4 defaults 0 0`

install client
```
yum install -y nfs-utils
systemctl enable nfs-idmap && systemctl start nfs-idmap
```

### Samba
smb fatab: `//smbserver/shared /mnt cifs rw,username=user01,password=pass 0 0`

install client
```
yum install -y cifs-utils
yum install -y samba-client
```

## Extend existing logical volumes
--extents vs --size: `--extents [+]LogicalExtentsNumber[%{VG|LV|PVS|FREE|ORIGIN}]` `--size [+]LogicalVolumeSize[bBsSkKmMgGtTpPeE]`

`-r` is for resizing the filesystem according to the volume

extend ext4 logical volume with all space in VG: `lvextend --extents +100%FREE -r /dev/vg/lv_vol`

extend ext4 logical volume by 50MB: `lvextend --size +50M -r /dev/vg/lv_vol`

reduce ext4 logical volume
```
umount /dev/vg/lv_vol
lvreduce --size -50M -r /dev/vg/lv_vol
```

extend xfs logical volume by 50MB: `lvextend --size +50M -r /dev/vg/lv_vol`

reduce xfs logical volume:  `-r` is missing because xfs cant be reduced automaticialy
```
lvextend --size +50M /dev/vg/lv_vol
xfs_growfs /mnt
```

## Create and configure set-GID directories for collaboration
set set-GID on folder: `chmod -R g+s foo/`

### umask
user wihtout login shell: `/etc/profile`

users with login shell: `/etc/bashrc`

```
max:   777 (permission)
     - 002 (umask)
     ______
     = 775 (final permission on file)
```

## Create and manage Access Control Lists (ACLs)
### How to
1. apply ACL to all **existing** files: `u:foo:rwx`
2. set ALC as default for existing files in folder: `d:u:foo:rwx`

XFS Filesystems have ACL per default activated. For ext4 file systems acls have to be enabled on mount
### Commands
set ACL: `setfacl -m u:andi:rwx,g:consultants:r foo/`

set ACL as default: `setfacl -m d:u:andi:rwx,g:consultants:r foo/`

remove acl: `setfacl -x foo/`

remove all acls from file `setfacl -b foo/`
## Diagnose and correct file permission problems