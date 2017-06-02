# Misc

## Docker

* Docker’s logging drivers expect logs to be sent to stdout and stderr
* Docker0 MTU only is set if there is a container to the bridge

#### Add CA Certs to CentOS 7 Images
copy ca.pem to `/etc/pki/ca-trust/source/anchors`

run `/bin/update-ca-trust` to update system CA's

## Kubernetes

- Kubelet/ Kube-Proxy **Root dir**: `/var/lib/`
- emptyDir Volume location based on *kubelet* root dir