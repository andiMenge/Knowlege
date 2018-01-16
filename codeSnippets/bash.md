# Bash Code Snippets

Script error handling `set -euo pipefail`

## Arrays

string array `arr=("bar1" "bar2")`

loop over array:
```
for i in ${arr[@]}; do
  echo $i
done
```

## Render multi line output into a single line
`kubectl get namespace | awk  '{print $1}' |grep -v NAME |xargs`

replace whitespace delimter with something else

`kubectl get namespace | awk  '{print $1}' |grep -v NAME |xargs |sed -e 's/ /,/g'`

## Error Handling
### Traps
```
# passes the string per signal to the cleanup() function
trap 'cleanup "INT signal received"' INT
trap 'cleanup "TERM signal received"' TERM
trap 'cleanup "unknown error occurred"' ERR
```

```
cleanup(){
  printf "\n$1\n"
  exit 1
}
```

## Check if binary exists
`requirements=("curl" "jq" "foo")`

```
checkRequirements() {
  for i in ${requirements[@]}; do
    command -v $i >/dev/null 2>&1 || { printf >&2 "\nRequirement not installed: $i\nAborting!\n\n"; exit 1; }
  done
}
```

## Variable Substitution
`foo=$(date "+%Y")`
