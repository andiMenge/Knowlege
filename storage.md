> IOP/s = Input or Output operations per second

> Throughput = How many MB/s can you read/write continuously

# Benchmarking
## fio
### IO Load Tests
Random Writes
`sync;fio --randrepeat=1 --ioengine=libaio --direct=1 --gtod_reduce=1 --name=test --filename=test --bs=4k --iodepth=256 --size=4G --readwrite=randwrite --ramp_time=4`

Random Reads
`sync;fio --randrepeat=1 --ioengine=libaio --direct=1 --gtod_reduce=1 --name=test --filename=test --bs=4k --iodepth=256 --size=4G --readwrite=randread --ramp_time=4`

Sequential Writes
`sync;fio --randrepeat=1 --ioengine=libaio --direct=1 --gtod_reduce=1 --name=test --filename=test --bs=4k --iodepth=256 --size=4G --readwrite=write --ramp_time=4`

Sequential Reads
`sync;fio --randrepeat=1 --ioengine=libaio --direct=1 --gtod_reduce=1 --name=test --filename=test --bs=4k --iodepth=256 --size=4G --readwrite=read --ramp_time=4`

### Throughput Load Tests
Random Writes
`sync;fio --randrepeat=1 --ioengine=libaio --direct=1 --gtod_reduce=1 --name=test --filename=test --bs=4M --iodepth=256 --size=10G --readwrite=randwrite --ramp_time=4`

Random Reads
`sync;fio --randrepeat=1 --ioengine=libaio --direct=1 --gtod_reduce=1 --name=test --filename=test --bs=4M --iodepth=256 --size=10G --readwrite=randread --ramp_time=4`

Sequential Writes
`sync;fio --randrepeat=1 --ioengine=libaio --direct=1 --gtod_reduce=1 --name=test --filename=test --bs=4M --iodepth=256 --size=10G --readwrite=write --ramp_time=4`

Sequential Reads
`sync;fio --randrepeat=1 --ioengine=libaio --direct=1 --gtod_reduce=1 --name=test --filename=test --bs=4M --iodepth=256 --size=10G --readwrite=read --ramp_time=4`

[src](https://smcleod.net/benchmarking-io/)

## ioping 
https://github.com/koct9i/ioping

# iScsi
## Get Initiator Name
`cat /etc/iscsi/initiatorname.iscsi`

## Discover targets
`iscsiadm -m discovery -t st -p <host-ip>`

## List Sessions
`iscsiadm -m session`

## Mount target
`iscsiadm -m node --targetname "<iqn>" --portal "<host-ip>" --login`

## Unmount target
`iscsiadm -m node --targetname "<iqn>" --portal "<host-ip>" --logout`
