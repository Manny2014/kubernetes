## CREATE TWO NETWORK NAMESPACES
# STEP #1: CREATE NETWORK NAMESPACE(s)
ip netns add red
ip netns add blue

# STEP #2: CREATE "virtual cable" (or pipe)
ip link add veth-red type veth peer name veth-blue

# STEP #3: Associate each end of the cable to respective namespace
ip link set veth-red netns red
ip link set veth-blue netns blue

# STEP #4: Assign IP Address and bring interfaces up
ip netns exec red ip addr add 192.168.15.1/24 dev veth-red
ip netns exec blue ip addr add 192.168.15.2/24 dev veth-blue

# STEP #5: Bring interfaces up
ip netns exec red ip link set veth-red up
ip netns exec blue ip link set veth-blue up

# STEP 4 & 5: Alternative
ip netns exec red ifconfig veth-red  192.168.15.1/24  up
ip netns exec blue ifconfig veth-blue 192.168.15.2/24 up

## CONNECT MORE THAN ONE NAMESPACE  (requires a virtual switch using linux bridge)
# STEP 6: Create virtual switch (interface on host)
ip link add v-switch-0 type bridge

# STEP 7: Bring up switch
ip link set dev v-switch-0 up

# STEP 8: Connect network namespaces to v-switch
## Delete old cables
ip netns exec red ip link delete veth-red # Other end gets deleted automatically

## New cables
ip link add veth-red type veth peer name veth-red-br
ip link add veth-blue type veth peer name veth-blue-br

# STEP 9: Assign one end of veth to namespaces
ip link set veth-red netns red 
ip link set veth-blue netns blue

# STEP 10: Assign the other end to the v-switch as master
ip link set veth-red-br master v-switch-0
ip link set veth-blue-br master v-switch-0

# STEP 11: Set IP Addresses and bring interfaces up
ip netns exec red ip addr add 192.168.15.1/24 dev veth-red
ip netns exec blue ip addr add 192.168.15.2/24 dev veth-blue

## Bring bridge end ups
ip link set veth-red-br up
ip link set veth-blue-br up

## Bring Ns interfaces up
ip netns exec red ip link set veth-red up
ip netns exec blue ip link set veth-blue up

# Establish connectivity from host -> namespaces (but NOT namespaces -> host)
## We just need to assign and IP to the v-switch-0
ip addr add 192.168.15.4/24 dev v-switch-0

# Establish connectivity from namespaces -> host -> world
## Add routing table gateway "door" which is the host
ip netns exec red ip route add 10.10.1.0/24 via 192.168.15.4
ip netns exec blue ip route add 10.10.1.0/24 via 192.168.15.4

# Add NAT Functionality
## create ip table rule to mask traffic from the NS network with its IP
iptables -t nat -A POSTROUTING -s 192.168.15.0/24 -j MASQUERADE

## Add route for any external network
ip netns exec red ip route add default via 192.168.15.4
ip netns exec blue ip route add default via 192.168.15.4

# Allow incoming traffic to network namespaces
## Create iptable rule to port foward from host -> ns
### This one is a rule specific to blue ns
iptables -t nat -A PREROUTING --dport 80 --to-desination 192.168.15.2:80 -j DNAT


iptables -A INPUT -s 10.10.1.0/24 -i virbr1 -p tcp -m state --state NEW  --dport 6443 -j ACCEPT
iptables -A INPUT -s 192.168.1.0/24 -i virbr1 -p tcp -m state --state NEW  --dport 6443 -j ACCEPT

iptables -A INPUT -s 10.10.1.0/24 -i br0 -p tcp -m state --state NEW --dport 6443 -j ACCEPT
iptables -A INPUT -s 192.168.1.0/24 -i br0 -p tcp  -m state --state NEW --dport 6443 -j ACCEPT

iptables -A INPUT -s 10.10.1.0/24 -i br1 -p tcp -m state --state NEW --dport 6443 -j ACCEPT
iptables -A INPUT -s 192.168.1.0/24 -i br1 -p tcp  -m state --state NEW --dport 6443 -j ACCEPT



# OPEN PORT 
iptables -I INPUT -s 192.168.1.0/24 -d -p tcp -m tcp --dport 6443 -j ACCEPT
iptables -I INPUT -s 192.168.1.0/24 -p udp -m udp --dport 6443 -j ACCEPT


iptables -A INPUT -i eno2 -p tcp --dport 6443 -j ACCEPT
iptables -A INPUT -i br1 -p tcp --dport 6443 -j ACCEPT

iptables -A INPUT -i virbr1 -p tcp --dport 6443 -j ACCEPT
iptables -A INPUT -i virbr0 -p tcp --dport 6443 -j ACCEPT

iptables -t nat -D PREROUTING -d 192.168.1.122 -p udp --dport 6443 -j DNAT --to 10.10.1.10:6443
iptables -D FORWARD -d 10.10.1.10/32 -p udp -m state --state NEW,ESTABLISHED,RELATED --dport 6443 -j ACCEPT
iptables -t nat -D PREROUTING -d 192.168.1.122  -p tcp --dport 6443 -j DNAT --to 10.10.1.10:6443
iptables -D FORWARD -d 10.10.1.10/32 -p tcp -m state --state NEW,ESTABLISHED,RELATED --dport 6443 -j ACCEPT

ip route add 10.10.1.0/24 via 192.168.1.122
ip route del 10.10.1.0/24 
