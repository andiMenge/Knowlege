# iScsi

## Discover targets
`iscsiadm -m discovery -t st -p <host-ip>`

## Mount target
`iscsiadm -m node --targetname "<iqn>" --portal "<host-ip>" --login`

## Unmount target
`iscsiadm -m node --targetname "<iqn>" --portal "<host-ip>" --logout`
