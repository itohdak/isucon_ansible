#!/bin/bash

# 第一引数を取得
PARAM=$1

cd /home/isucon/isucon_ansible
ansible-playbook playbooks/deploy_repo.yaml --extra-vars "branch=$PARAM"