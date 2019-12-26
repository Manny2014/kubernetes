# Docker networking
- No connectivity

## HOST Network
- No network isolation
- Causes port conflicts when assigning a port to a container

## BRIDGE
- Internal prive network is create and containers attach to it
- When docker is installed it creates a private local network called "bridge"
- On the host, the network is called *docker0*
    - Similar to: ```ip link add docker0 type bridge```
- Docker also creates a network namespace
    - ```ip netns list```
- Container == Network Namespace
- Find network namespace from a docker container
    - ```docker inspect --format '{{.State.Pid}}' <container_name_or_Id>```
    - ```nsenter -t <contanier_pid> -n <command>```
    - See: https://stackoverflow.com/questions/31265993/docker-networking-namespace-not-visible-in-ip-netns-list
- Exposing ports from containers
    - Accomplished by using iptables
        - ```iptables -t nat -A PREROUTING --dport 80 --to-desination <continer_ip>:80 -j DNAT```

# CNI
- Standards that define how creating network namespaces and adding a container to it should look like
    - Ex: ```pseudo-bridge add <CONTAINRE> /var/run/netns/ew9thrndf```
- Programs are referred to as *plugins*

