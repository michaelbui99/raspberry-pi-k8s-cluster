#!/bin/bash

read -p "Cluster master user: " user
read -p "Cluster master ip: " ip

CONNECTION_STRING="${user}@${ip}"

echo "Copying .kube from $CONNECTION_STRING"
scp -r "${CONNECTION_STRING}/home/${user}/kube ."

mv ./.kube ~/.kube

