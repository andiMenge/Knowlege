# Configure local storage
flow: `create partition table -> create partitions -> create pv -> create vg -> add pv to vg -> create lv from vg -> create fs on lv -> mount lv`

## List, create, delete partitions on MBR and GPT disks
list disks `lsblk -a`

select disk for edit `gdisk /dev/sdb`

create partition tabel `n`

print current state `p`

write partition to disk `w`

To force the kernel to read the updated partition table `partprobe`

## Create and remove physical volumes, assign physical volumes to volume groups, and create and delete logical volumes
list disks `lsblk -a`

### Pysical Volumes
pysikal volume: `pvcreate /dev/sdX`,  `pvremove /dev/sdX`, `pvs`

### Volume Groups
volume group: `vgcreate -s <size> <name> /dev/sdX` `vgs`

add pv to volume group: `vgextend <vg-name> /dev/sdX`

remove pv from volume group: `vgreduce <vg-name> /dev/sdX`

remove vg: `vgremove <vg-name>`

### Logical Volumes
show volumes `lvs`

volume details `lvdisplay`

create volume `lvcreate --size 1G --name lv_vol <vg-name>`

remove volume `lvremove /dev/vg/lv_vol`

resize volume `lvresize -L 10G -r /dev/vg_new/lvol0`

format lv `mkfs.ext4 /dev/mapper/<lv-name>`

## Configure systems to mount file systems at boot by Universally Unique ID (UUID) or label
get disk UUID `blkid` (disk must have FS)

add entry to fstab
```
UUID=749f0100-ec71-4617-9c6e-18b2eaead929 /mnt/media  ext4  defaults  0 0
/dev/mapper/data-office--foo  /mnt/office ext4  defaults  0 0
```

## Add new partitions and logical volumes, and swap to a system non-destructively
format lv as swap `mkswap /dev/<vg-name>/<lv-name>`

show swaps `cat /proc/swaps`

add swap `swapon /dev/mapper/<lv-name>`

to persist swap add to fstab `/dev/mapper/vg-lv_swap swap swap defaults 0 0`
