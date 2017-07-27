# Arrays

string array `arr=("bar1" "bar2")`

loop over array:
```
for i in ${arr[@]}; do
  echo $i
done
```

# Render multi line output into a single line
`kubectl get namespace | awk  '{print $1}' |grep -v NAME |xargs`

replace whitespace delimter with something else

`kubectl get namespace | awk  '{print $1}' |grep -v NAME |xargs |sed -e 's/ /,/g'`

# Error Handling
`set -eE` Terminate on any error (return code != 0) but execute the ERR trap before exiting

```
trap 'my-cleanup-func "INT signal received"' INT
trap 'my-cleanup-func "TERM signal received"' TERM
trap 'my-cleanup-func "unknown error occurred"' ERR
```
