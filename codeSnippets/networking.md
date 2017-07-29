# Networking

## nmap

Find all hosts in network
`nmap -sP -n <network/subnet>`

Find host details without ICMP
`sudo nmap -sS -n <ip>`

## Test JumboFrame Connection
`ping -M do -s 8972 <ip>`

8972 + tcp-overhead = 9000 MTU

## Remote Network Traffic Inspection
`ssh <srv> sudo tcpdump -i <interface> -U -s0 -w - 'not port 22' | wireshark -k -i -`
