#!/bin/sh

RESOURCE_CTLR_PATH=${RESOURCE_CTLR_PATH:-../../resource-ctlr/deploy/crds}

for i in $RESOURCE_CTLR_PATH/*_crd.yaml; do
  kubectl apply -f $i
done

for i in ../deploy/crds/*_crd.yaml; do
  kubectl apply -f $i
done

for i in controller01 controller02 controller03 worker01; do
  private_key=$(vagrant ssh-config "$i" | grep IdentityFile | awk '{print $2}')
  ip=$(grep -A 1 $i Vagrantfile | grep vm.network | awk -F\" '{print $2}')
  kubectl create secret generic "$i" --from-file=ssh-privatekey=${private_key} --dry-run -o yaml | kubectl apply -f -
  kubectl apply -f -<<EOF
apiVersion: resources.weisnix.org/v1alpha1
kind: Host
metadata:
  name: $i
spec:
  ipaddress: $ip
  sshkeysecret: $i
  port: 22
EOF

done

kubectl apply -f -<<EOF
apiVersion: kubeadm.weisnix.org/v1alpha1
kind: Cluster
metadata:
  name: vagrant
spec:
  controllers:
    - name: controller01
      ip: 192.168.10.2
    - name: controller02
      ip: 192.168.10.3
    - name: controller03
      ip: 192.168.10.4
  workers:
    - name: worker01
      ip: 192.168.10.5
EOF