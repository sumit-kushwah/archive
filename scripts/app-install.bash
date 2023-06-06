#!/bin/bash
# This script is tested and written for Ubuntu linux system
# This script excepts user to interact at some point

# update apt cache and update packages
sudo apt update
sudo apt upgrade -y
sudo apt install curl

# install google chrome
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
sudo dpkg -i google-chrome-stable_current_amd64.deb
rm ./google-chrome-stable_current_amd64.deb
echo "google chrome installed"

# install vscode
# wget -qO- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > microsoft.gpg
# sudo mv microsoft.gpg /etc/apt/trusted.gpg.d/microsoft.gpg
# echo "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main" | sudo tee /etc/apt/sources.list.d/vscode.list
# sudo apt update
# sudo apt install code
snap install vscode --classic
echo "vscode installed"


# install termius
# wget https://www.termius.com/download/linux/Termius.deb
# sudo dpkg -i Termius.deb
# rm ./Termius.deb
snap install termius-app
echo "termius installed"

# install docker
sudo apt-get remove docker docker-engine docker.io containerd runc
sudo apt update
sudo apt install apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt update
sudo apt install docker-ce docker-ce-cli containerd.io
sudo systemctl start docker
sudo usermod -aG docker $USER
sudo systemctl restart docker
echo "docker installed"

# install docker compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
echo "docker-compose installed"

# install dbeaver
snap install dbeaver-ce 

# install todoist
snap install todoist

# install spotify
snap install spotify

# install other development tools
sudo apt install git
sudo apt install vim
sudo apt install zsh
sudo apt install htop
echo "git, vim and zsh installed"

# install anaconda3
bash ./anaconda-install.sh

# install oh-my-zsh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
echo "oh my zsh installed"

# generate ssh key without passphrase
ssh-keygen -t rsa -b 4096
echo "ssh key generated"
