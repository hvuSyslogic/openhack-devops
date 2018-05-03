echo "############### Removing requiretty ###############"
sed -i "s/^.*requiretty/#Defaults requiretty/" /etc/sudoers

echo "############### Installing Azure CLI v2.0.31 ###############"
sudo rpm --import https://packages.microsoft.com/keys/microsoft.asc
sudo sh -c 'echo -e "[azure-cli]\nname=Azure CLI\nbaseurl=https://packages.microsoft.com/yumrepos/azure-cli\nenabled=1\ngpgcheck=1\ngpgkey=https://packages.microsoft.com/keys/microsoft.asc" > /etc/yum.repos.d/azure-cli.repo'
sudo yum install -y azure-cli-2.0.31-1.el7.x86_64

echo "############### Installing Helm v2.9.0 ###############"
sudo curl -O https://storage.googleapis.com/kubernetes-helm/helm-v2.9.0-linux-amd64.tar.gz
sudo tar -zxvf helm-v2.9.0-linux-amd64.tar.gz
sudo mv linux-amd64/helm /usr/local/bin/helm

echo "############### Installing Dotnet SDK v2.1.4 ###############"
sudo rpm --import https://packages.microsoft.com/keys/microsoft.asc
sudo sh -c 'echo -e "[packages-microsoft-com-prod]\nname=packages-microsoft-com-prod \nbaseurl= https://packages.microsoft.com/yumrepos/microsoft-rhel7.3-prod\nenabled=1\ngpgcheck=1\ngpgkey=https://packages.microsoft.com/keys/microsoft.asc" > /etc/yum.repos.d/dotnetdev.repo'
sudo yum update -y
sudo yum install -y libunwind libicu
sudo yum install -y dotnet-sdk-2.1.4

echo "############### Installing Jq v1.5 ###############"
sudo yum install -y epel-release
sudo yum install -y jq

echo "############### Installing Docker ###############"
echo "sudo yum install -y yum-utils device-mapper-persistent-data lvm2"
sudo yum install -y yum-utils \
  device-mapper-persistent-data \
  lvm2
echo "sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo"
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
echo "sudo yum install -y docker-ce"
sudo yum install -y docker-ce
echo "sudo systemctl start docker"
sudo systemctl start docker
echo "sudo systemctl enable docker"
sudo systemctl enable docker
echo "sudo groupadd docker"
sudo groupadd docker
echo "sudo usermod -aG docker $USER"
sudo usermod -aG docker $USER

echo "############### Pulling Openhack-tools from Github "###############
git clone https://github.com/Azure-Samples/openhack-devops-tools.git /home/azureuser
