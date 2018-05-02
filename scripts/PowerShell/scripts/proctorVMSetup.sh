echo "############### Installing Azure CLI v2.0.31 ###############"
rpm --import https://packages.microsoft.com/keys/microsoft.asc
echo -e "[azure-cli]\nname=Azure CLI\nbaseurl=https://packages.microsoft.com/yumrepos/azure-cli\nenabled=1\ngpgcheck=1\ngpgkey=https://packages.microsoft.com/keys/microsoft.asc" > /etc/yum.repos.d/azure-cli.repo
yum install azure-cli-2.0.31-1.el7.x86_64 -y

echo "############### Installing Helm v2.9.0 ###############"
curl -O https://storage.googleapis.com/kubernetes-helm/helm-v2.9.0-linux-amd64.tar.gz
tar -zxvf helm-v2.9.0-linux-amd64.tar.gz
mv linux-amd64/helm /usr/local/bin/helm

echo "############### Installing Dotnet SDK v2.1.4 ###############"
rpm --import https://packages.microsoft.com/keys/microsoft.asc
echo -e "[packages-microsoft-com-prod]\nname=packages-microsoft-com-prod \nbaseurl= https://packages.microsoft.com/yumrepos/microsoft-rhel7.3-prod\nenabled=1\ngpgcheck=1\ngpgkey=https://packages.microsoft.com/keys/microsoft.asc" > /etc/yum.repos.d/dotnetdev.repo
yum update
yum install libunwind libicu -y
yum install dotnet-sdk-2.1.4 -y

echo "############### Installing Jq v1.5 ###############"
yum install -y epel-release
yum install -y jq

echo "############### Installing Docker ###############"
yum check-update
yum install -y yum-utils \
  device-mapper-persistent-data \
  lvm2
yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
yum install -y docker-ce
systemctl start docker
groupadd docker
usermod -aG docker $USER
