#!/bin/bash

CERTS=config

cat <<EOF > infra/local/scripts/kind-config-with-mounts-and-ingress.yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  extraMounts:
  - hostPath: ${CERTS}
    containerPath: /config
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 5671
    hostPort: 5671
    protocol: TCP
  - containerPort: 15672
    hostPort: 15672
    protocol: TCP
EOF
