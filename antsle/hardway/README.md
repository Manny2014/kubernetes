# Installs
- ```brew install https://raw.githubusercontent.com/kadwanev/bigboybrew/master/Library/Formula/sshpass.r```
- ```curl -o cfssl https://storage.googleapis.com/kubernetes-the-hard-way/cfssl/darwin/cfssl```
- ```curl -o cfssljson https://storage.googleapis.com/kubernetes-the-hard-way/cfssl/darwin/cfssljson```


# Routes

```
# DEFAULT
ip route add 10.10.1.0/24 via 192.168.1.122

# MASTERS
ip route add 10.10.1.10/32 via 192.168.1.129
ip route add 10.10.1.11/32 via 192.168.1.130
ip route add 10.10.1.12/32 via 192.168.1.131

# WORKERS
ip route add 10.10.1.20/32 via 192.168.1.132
ip route add 10.10.1.21/32 via 192.168.1.133
ip route add 10.10.1.22/32 via 192.168.1.134

# PROXY
ip route add 10.10.1.30/32 via 192.168.1.128
```