#!bin/bash

skip=false
read -p "Cluster node user: " user
read -p "Cluster node ip: " ip
CONNECTION_STRING="${user}@${ip}"

if [[ -f ~/.ssh/id_rsa ]]; then
    echo "Detected existing keys"
    echo "Skipping ssh-keygen..."
    $skip=true
fi


if [[ $skip = false ]]; then
    ssh-keygen -b 4096 -N "" -f ~/ssh/id_rsa
fi

ssh-copy-id "$CONNECTION_STRING"



