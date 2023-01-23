#!bin/bash

echo Available Playbooks:
for playbook in $(ls ./playbooks); do
    echo $playbook
done
echo 

read -p "Playbook: " playbook

ansible-playbook -i ./inventory/hosts.ini ./playbooks/${playbook}
