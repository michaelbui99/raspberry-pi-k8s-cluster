#!bin/bash

read -p "IP Address to assign: " ip
read -p "Default Gateway IP: " gateway
read -p "DNS (Defaults to gateway IP): " dns

if [[ ! $dns ]]; then
    dns=$gateway
fi

sudo cat <<EOT >> /etc/dhcpcd.conf
interface eth0
static ip_address=${ip}/24
static routers=${gateway}
static domain_name_servers=${dns}
EOT