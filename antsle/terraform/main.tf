provider "antsle" {
  api_key = "Token eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoicm9vdCJ9.4V8GfhhvljynrxY1wRjjQ-O0wIrpPX0Np12CjiaV773bLZcgj_xUpur8wUeCNiPcUnuLxrKuvo5e-3IEARIKDQ"
}

# He stays up
resource "antsle_antlets" "proxy" {
  dname = "k8s1-proxy"
  template = "CentOS-7"
  ram = 512
  cpu = 1
  antlet_num = 30 # IP
  zpool_name = "antlets"
  compression = "lz4"
  autostart = true
}

resource "antsle_antlets" "master-0" {
  dname = "k8s1-master-0"
  template = "CentOS-7"
  ram = 1024
  cpu = 1
  antlet_num = 10 # IP
  zpool_name = "antlets"
  compression = "off"
  autostart = true
}

resource "antsle_antlets" "master-1" {
  dname = "k8s1-master-1"
  template = "CentOS-7"
  ram = 1024
  cpu = 1
  antlet_num = 11 # IP
  zpool_name = "antlets"
  compression = "off"
  autostart = true
}

resource "antsle_antlets" "master-2" {
  dname = "k8s1-master-2"
  template = "CentOS-7"
  ram = 512
  cpu = 1
  antlet_num = 12 # IP
  zpool_name = "antlets"
  compression = "off"
  autostart = true
}

resource "antsle_antlets" "worker-0" {
  dname = "k8s1-worker-0"
  template = "CentOS-7"
  ram = 4000
  cpu = 2
  antlet_num = 20 # IP
  zpool_name = "antlets"
  compression = "off"
  autostart = true
}

resource "antsle_antlets" "worker-1" {
  dname = "k8s1-worker-1"
  template = "CentOS-7"
  ram = 4000
  cpu = 2
  antlet_num = 21 # IP
  zpool_name = "antlets"
  compression = "off"
  autostart = true
}

resource "antsle_antlets" "worker-2" {
  dname = "k8s1-worker-2"
  template = "CentOS-7"
  ram = 4000
  cpu = 2
  antlet_num = 22 # IP
  zpool_name = "antlets"
  compression = "off"
  autostart = true
}