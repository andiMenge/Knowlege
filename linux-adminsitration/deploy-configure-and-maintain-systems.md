# Deploy, configure, and maintain systems
## Configure networking and hostname resolution statically or dynamically
hostname should be FQDN
## Schedule tasks using at and cron
## Start and stop services and configure services to start automatically at boot
list all failed services: `systemctl --type service --state failed`

Disable a service **globally**:
- stop service
- mask service

## Configure systems to boot into a specific target automatically
## Install Red Hat Enterprise Linux automatically using Kickstart
## Configure a physical machine to host virtual guests
## Install Red Hat Enterprise Linux systems as virtual guests
## Configure systems to launch virtual machines at boot
## Configure network services to start automatically at boot
## Configure a system to use time services
## Install and update software packages from Red Hat Network, a remote repository, or from the local file system
where does this file come from/ what do I have to do to get this file (uses only configured repos):  `yum provide`
`yum history`

`/var/log/yum.log`

`yum group list hidden -v`

install a group and a single package: `yum install @postgresql firefox`

## Update the kernel package appropriately to ensure a bootable system
## Modify the system bootloader
grub2 config: `/boot/grub2/grub.cfg`

create new grubconfig: `grub2-mkconfig > /boot/grub2/grub.cfg`

## RHEL Boot Process
1. UEFI or BIOS runs *Power on self test (POST)* and starts to initialize some hardware
2. UEFI/ BIOS searches for a bootable disk either in the UEFI firmware or looking for a *MBR* on the disks listed in the BIOS
3. UEFI/ BIOS reads a boot loader from a disk and passes control of the system to the boot loader (*grub2*)
4. Boot loader loads its configuration and presents options on screen `/boot/grub2/grub.cfg`, `/etc/grub.d`, `/etc/default/grub`
5. After Timeout or User choice the boot loader loads the configured **kernel** and the **initramfs** into the RAM.
6. boot loader hands control of the system over to the kernel with all options configured on the kernel line in the boot loader config and the memory adress of the *initramfs*.
7. The kernel initializes all harware fo which it can find a driver in the initramfs. Also it does executes `/sbin/init` from initramfs as **PID 1**
8. the systemd from the initramfs executes all units for *initrd.target* and mounts the root filesystem on **/sysroot**
9. the kernel root filesystem is switched from the initramfs root filesystem to the root filesystem that was mounted on /sysroot. Systemd the re-executes itself using the copy installed on the system.

10. Systemd looks for a target either passed from the kernel commandline or configured on the system.
It then starts the units configured for the specified target.

### Initramfs
config: `/etc/dracut.conf`

> gzip encoded cpio archive containing kernel modules for all hardware that is nessesary to boot, init scripts, a working systemd copy, a udev daemon and more.