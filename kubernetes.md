# Misc
- Kubelet/ Kube-Proxy **Root dir**: `/var/lib/`
- emptyDir Volume location based on *kubelet* root dir. It will **fill up** your partition!!

# Proxy a Connection from local machine to a pod (localport:podPort)
`kubectl -n tools port-forward $(kubectl -n <your-namespace> get pod -l <your-label> |grep <your-deployment-name> |cut -d ' ' -f 1) 1337:8000`

# Pod Events

- `Need to kill Pod` indicates a resource limit being exceeded
