# Manage security

## Configure firewall settings using firewall-config, firewall-cmd, or iptables

>Zones are basically sets of rules dictating what traffic should be allowed depending on the level of trust you have in the networks your computer is connected to

---

firewalld config file: - `/etc/firewalld/firewalld.conf` `/usr/lib/firewalld/services`
new services should be added to `/etc/firewalld/firewalld.conf`. Config in `/etc/firewalld` wins over `/usr/lib/..`

reload firewalld config (no connection loss) `firewall-cmd --reload`

reload firewalld config and netfilter kernel modul (drops network connections) `firewall-cmd --complete-reload`

make changes permanent (requires reload): `firewall-cmd --permanent` 
*These changes are not effective immediately, only after service restart/reload or system reboot.* 

make runtime config permanent `firewall-cmd –runtime-to-permanent`

get zones `firewall-cmd --get-zones`

get zone info: `firewall-cmd --info-zone public`

get active zones: `firewall-cmd --get-active-zones`

get default zone: `firewall-cmd --get-default-zone`

get config of default zone: `firewall-cmd --list-all`

set new default zone: `firewall-cmd --set-default-zone=home`

add network interface to zone: `firewall-cmd --permanent --zone=internal --change-interface=eth0`

get zone of an interface: `firewall-cmd --get-zone-of-interface=eth0`

### Sources

A **zone** can be bound to a network interface (see above) and/or to a **source**

add source to a zone `firewall-cmd --permanent --zone=trusted --add-source=192.168.2.0/24`

get all sources of a zone: `firewall-cmd --permanent --zone=trusted --list-sources`

### Service Management

add service to zone:

```
firewall-cmd --permanent --zone=internal --add-service=http
firewall-cmd --reload
```

list servies: `firewall-cmd --list-services` `firewall-cmd --list-services --zone=internal`

get details on service: `firewall-cmd --info-service ssh`

create new service skeleton in `/etc/firewalld/services`: `firewall-cmd --permanent --new-service=haproxy`

### Port Management

add port to zone: `firewall-cmd --permanent --zone=public --add-port=5000/tcp` `firewall-cmd --reload` `--add-port=5000-5100/tcp`

## Configure key-based authentication for SSH

add ssh keys to current user: `ssh-keygen -b 4096 -t rsa` 

## Set enforcing and permissive modes for SELinux
SELinux Config `/etc/selinux/config` or `/etc/sysconfig/selinux`

SELinux Modes `enforcing` `permissive` `disabled`

SELinux status `sestatus`

Get SELinux Mode `getenforce`

Change SELinux Mode (non permanent) `setenforce <0|1>`

Change SELinux Mode (Permanent) `SELINUX=permissive` in `/etc/selinux/config`

## List and identify SELinux file and process context
file context `ls -Z`

process context `ps -eZ`
## Restore default file contexts
### flow
1. add/ modify entry in SELinux DB
2. apply context from DB to file/ folder

### Install SELinux Man Pages
`yum install -y selinux-policy-doc  `

`mandb`

`man -k _selinux |grep httpd`

---

add context entry in SELinux DB: `semanage fcontext -a -t httpd_sys_content_t /foo/index.html`

modify context entry in SELinux DB: `semanage fcontext -m -t httpd_sys_content_t /foo/index.html`

apply context from SELinux DB to specified files/ folders `restorecon -Rv /path`

## Use boolean settings to modify system SELinux settings
list policies `semanage boolean -l` `getsebool -a`

list policies with local modifications `semanage boolean -l -C`

list policies from nfs `semanage boolean -l | egrep "nfs|SELinux"`

set boolean (Permanent) `setsebool -P <policy> <on|off>`

set boolean (Non Permanent) `setsebool <policy> <on|off>`

## Diagnose and address routine SELinux policy violations

all SELinux Policies are stored in this package: *selinux-policy-targeted*

install semanage tool: `yum install -y setroubleshoot-server`

check policy violations `sealert -a /var/log/audit/audit.log`

get details on violation `grep 1415714880.156:29 /var/log/audit/audit.log | audit2why`