# Misc

## Docker

* Docker’s logging drivers expect logs to be sent to stdout and stderr
* Docker0 MTU only is set if there is a container to the bridge

### Docker for Mac and no_proxy env var
`ENV no_proxy=foo` in a dockerfile gets overwritten by proxy and no_proxy conf in the **docker-engine** at least for macOS

#### Add CA Certs to CentOS 7 Images
copy ca.pem to `/etc/pki/ca-trust/source/anchors`

run `/bin/update-ca-trust` to update system CA's

## Kubernetes

- Kubelet/ Kube-Proxy **Root dir**: `/var/lib/`
- emptyDir Volume location based on *kubelet* root dir

## Helm

- Control Whitespace Trimming with dashes `{{- if .Values.persistence }}`

- `includes` render the templates inside them `template` does not 