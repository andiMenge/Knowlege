# Manage users and groups
## Create, delete, and modify local user accounts
global config for users, passwords, groups: `/etc/login.defs`

list users: `cat /etc/passwd`

add new user: `useradd <username>`

delete user with homedir: `userdel -r <username>`

modify user: `usermod`

change username `usermod -l <new-name> <user-name>`
## Change passwords and adjust password aging for local user accounts

assign pw to user: `passwd <user-name>`

get date in future: `date -d "+90 days" +%F`

set expiration date for user `chage -E YYYY-MM-DD <user-name>`

show account aging information `chage -l <user-name>`

## Create, delete, and modify local groups and group memberships

list groups: `cat /etc/group` `groupmems -g <grp-name> -l`

add new group: `groupadd <name>`

modify groups: `groupmod`

change group name: `groupmod -n <new-name> <name>`

delete group: `groupdel <name>`

remove user from secondary group `gpasswd -d <user> <group>`

## Configure a system to use an existing authentication service for user and group information
install packages: `yum -y install authconfig-gtk.x86_64 sssd krb5-workstation.x86_64`

GUI: `Applications -> Sundry -> Authentication`

sssd: ldap auth daemon (caches auth data)