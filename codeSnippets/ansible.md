## Restart all Services (as privileged user 'root')
`ansible -i <path-to-inventory> masters -m "systemd state=restarted name=kube-apiserver" -b`

## Install rpm with environment variables
```
- name: Ensure RPM is installed
  yum:
     name: <abs-path-to-rpm>
     state: present
  environment: 
    MY_ENV: "bar"
  become: true
  when: foo is defined
  tags: [mytag]
```
