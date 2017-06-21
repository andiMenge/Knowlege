# Misc

## Docker

* Docker’s logging drivers expect logs to be sent to stdout and stderr
* Docker0 MTU only is set if there is a container to the bridge

### Docker for Mac and no_proxy env var
`ENV no_proxy=foo` in a dockerfile gets overwritten by proxy and no_proxy conf in the **docker-engine** at least for macOS

### Alpine Linux with self-signed ca-certs
ca-certs must be copied `cp` (`ln -s`) is not working to `/usr/local/share/ca-certificates/`

After `update-ca-certificates` ran the cert should be get symlinked here `/etc/ssl/certs/`.

Also you should find the content of your cert here `/etc/ssl/certs/ca-certificates.crt`. Just use `grep` of a small portion of your cert's hash.

### Dockerfile Best Practice
- Use exec/JSON form of ENTRYPOINT because the *shell* form runs the entrypoint as a subcommand of `/bin/sh -c`
- use `exec` in entrypoint scripts so that your app doesnt get spawned in a subshell and is able to recive SIGNALS
- Avoid beeing PID 1 because of process SIGNAL handling. Use `tini` or `dumb_init`
- Listen for SIGTERM or set STOPSIGNAL in the Dockerfile

### Add CA Certs to CentOS 7 Images
copy ca.pem to `/etc/pki/ca-trust/source/anchors`

run `/bin/update-ca-trust` to update system CA's

## Kubernetes

- Kubelet/ Kube-Proxy **Root dir**: `/var/lib/`
- emptyDir Volume location based on *kubelet* root dir and can fill up your partition

## Helm

- Control Whitespace Trimming with dashes `{{- if .Values.persistence }}`

- `includes` render the templates inside them `template` does not 
