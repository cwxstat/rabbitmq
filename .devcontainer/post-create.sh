go install sigs.k8s.io/kind@latest
sudo apt-get -y install emacs
mkdir -p ~/.kube
echo $PIG|base64 -d > ~/.kube/config
cat <<EOF >> ~/.bashrc
alias k='kubectl'
alias e='emacs -nw'
alias ubuntu='kubectl run --rm -i --tty ubuntu --image=ubuntu:latest --restart=Never -- bash -il'
alias dev='kubectl run --rm -i --tty dev --image=mchirico/ubuntu:latest --restart=Never --pod-running-timeout=6m0s -- bash -il'
alias udev='kubectl exec -it dev -- /bin/bash'
EOF

cat <<EOF > ~/.gitconfig
[core]
        editor = emacs
[alias]
        st = status
        co = checkout
        p = push
    pn = push origin n --force
        di = diff --staged
        ll = log --format=%B origin..HEAD
        br = branch
        cm = commit -am 'update-automated-save'
    rb = rebase master -i
        l =  log --oneline
        ls =  log --stat
        lp = log -p
        ln = log --oneline notes/commits
        la = log --oneline --graph --decorate --all
EOF

# Argo
curl -sLO https://github.com/argoproj/argo-workflows/releases/download/v3.1.9/argo-linux-amd64.gz
gunzip argo-linux-amd64.gz
chmod +x argo-linux-amd64
sudo mv ./argo-linux-amd64 /usr/local/bin/argo

# k8s
curl -sLO https://github.com/ahmetb/kubectx/releases/download/v0.9.4/kubectx -o kubectx
curl -sLO https://github.com/ahmetb/kubectx/releases/download/v0.9.4/kubens -o kubens
chmod +x kubectx
chmod +x kubens
sudo mv kubectx /usr/local/bin/kubectx
sudo mv kubens /usr/local/bin/kubens


curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
rm -rf ./aws
rm -rf awscliv2.zip

# Terraform
sudo apt-get update && sudo apt-get install -y gnupg software-properties-common curl
curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
sudo apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
sudo apt-get update && sudo apt-get install terraform
terraform init -upgrade



curl -sLO https://github.com/vmware-tanzu/octant/releases/download/v0.24.0/octant_0.24.0_Linux-64bit.tar.gz
tar -xzf octant_0.24.0_Linux-64bit.tar.gz
rm octant_0.24.0_Linux-64bit.tar.gz 
sudo cp octant_0.24.0_Linux-64bit/octant /user/local/bin/octant
rm -rf octant_0.24.0_Linux-64bit

curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash
sudo mv kustomize /usr/local/bin/kustomize


# calico
curl -o calicoctl -O -L  "https://github.com/projectcalico/calicoctl/releases/download/v3.20.0/calicoctl" 
chmod +x calicoctl
sudo mv calicoctl /usr/local/bin/calicoctl





curl https://sdk.cloud.google.com > install.sh
bash install.sh --disable-prompts

sudo apt install python3.8-venv -y


# Cobra
go get -u github.com/spf13/cobra/...

# Ginkgo
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega/...

