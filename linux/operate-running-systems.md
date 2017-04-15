# Operate running systems
## Boot, reboot, and shut down a system normally
`systemctl reboot`

`systemctl halt`

`poweroff`

Halting involves stopping all CPUs on the system.
Powering off involves sending an ACPI command to signal the PSU to disconnect main power.

## Boot systems into different targets manually
switch target

`systemctl isolate graphical.target`

set default target
`systemctl set-default rescue.target`

`systemctl get-default`

get all targets
`systemctl list-units --type=target --all`

```
ll /usr/lib/systemd/system/runlevel*
runlevel0.target -> poweroff.target
runlevel1.target -> rescue.target #Single User Mode
runlevel2.target -> multi-user.target #user-defined
runlevel3.target -> multi-user.target
runlevel4.target -> multi-user.target #user-defined
runlevel5.target -> graphical.target
runlevel6.target -> reboot.target
```
## Interrupt the boot process in order to gain access to a system
1. interrupt boot process (space, `e`)
2. kernel line (the line starting with linux16) and add the following statements at the end `enforcing=0 rd.break`
4. Mount /sysroot rw `mount –o remount,rw /sysroot`
5. change root directory `chroot /sysroot` create new chroot jail aka use `/sysroot` as `/`
6. change root pw `passwd`
7. re-initialize SELinux contexts on boot: `touch /.autorelabel`
7. logout

## Identify CPU/memory intensive processes, adjust process priority with renice, and kill processes
- find niceness of a process: `top` NI colum
- Niceness: -20 to +19 (integer steps) -20 highest priority
- start process with specific niceness: `nice -n <niceness> <command>`
- change process priority: `renice <niceness> -p <pid>` `renice +5 <pid>`
- send SIGkill to process: `kill -9 <pid>`
- send SIGterm to process: `kill <pid>`
- other process tools: `pkill` `pgrep`

## Locate and interpret system log files and journals
- Analyze boot process `systemd-analyze` `systemd-analyze blame`

persist journal files
```
mkdir /var/log/journal
echo "SystemMaxUse=50M" >> /etc/systemd/journald.conf
systemctl restart systemd-journald
```
## Access a virtual machine's console
- open KVM GUI `virt-manager`
- gust management on terminal `virsh`
- todo: Emergency procedure

## Start and stop virtual machines
- `virsh start vm.example.com`
- `virsh list` `virt-top`

## Start, stop, and check the status of network services
- `systemctl [start stop restart] <service>`
- check if service is active `systemctl is-active httpd`
- permanently disable a service `systemctl mask httpd` `systemctl unmask httpd`

## Securely transfer files between systems
scp stuff