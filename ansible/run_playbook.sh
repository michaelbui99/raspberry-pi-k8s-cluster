#!bin/bash

read -p "Playbook: " playbook
read -p "User: " user

ansible-playbook --user ${user} --ask-pass --ask-become -i ./inventory/hosts.ini ./playbooks/${playbook}
