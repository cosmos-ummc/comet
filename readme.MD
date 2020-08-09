# Backend Deployment

## Initialize Firewall

ssh root@[server ip]

ufw allow OpenSSH

ufw enable

exit

## Create User
ssh root@[server ip]

adduser [user]

usermod -aG sudo [user]

rsync --archive --chown=[user]:[user] ~/.ssh /home/[user]

exit

## Install Golang Dependencies

cd ~

curl -O https://dl.google.com/go/go1.13.9.linux-amd64.tar.gz

sudo tar -xvf go1.13.9.linux-amd64.tar.gz -C /usr/local

mkdir -p $HOME/go/{bin,src}

nano ~/.profile

export GOPATH=$HOME/go

export PATH=$PATH:$GOPATH/bin

export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin

. ~/.profile

## Pull Backend Project

cd $GOPATH/src

env GIT_TERMINAL_PROMPT=1 go get github.com/cosmos-ummc/comet

## Build Backend Project

cd github.com/cosmos-ummc/comet

go build cmd/server/main.go

## Initialize Service

sudo nano /lib/systemd/system/mainbe.service

[Unit]

Description=Golang Service

ConditionPathExists=/home/wing8169/go/src/github.com/cosmos-ummc/comet/main

After=network.target

[Service]

Restart=on-failure

RestartSec=10

startLimitIntervalSec=60

WorkingDirectory=/home/wing8169/go/src/github.com/cosmos-ummc/comet

ExecStart=/home/wing8169/go/src/github.com/cosmos-ummc/comet/main

[Install]

WantedBy=multi-user.target

## Input Environment Variables

systemctl edit mainbe

[Service]

Environment="AUTH_ENABLED=true"

Environment="BACKEND_USER=ee00c51c-5ce2-461d-9526-862e2830a113"

Environment="MONGODB_URL=mongodb://localhost:27017"

Environment="ACCESS_SECRET=$2a$04$gKZuI64CEh7Jovqf.UTefO58PwT1YuGooEHG.W7oX8AYKFqWd9tQ6"

Environment="REFRESH_SECRET=$2a$04$Sy7aXxwd.auGV4diE8u3iueKpAS.vaY4XSA.Zji6pGOrrY8Ox3eHa"

Environment="EMAIL_ADDRESS=cryptobe2020@gmail.com"

Environment="EMAIL_PASSWORD=CryptoLearning2020!"

Environment="SMTP_SERVER_HOST=smtp.gmail.com"

Environment="SMTP_SERVER_PORT=587"

Environment="TEMPLATES_PATH=templates"

Environment="ADMIN_URL=http://localhost"

## Start Service

sudo service mainbe start

## Restart Service

sudo service mainbe restart

## Expose Port in Firewall

sudo ufw allow 10001/tcp

sudo ufw allow 10002/tcp

## View Service

sudo lsof -i -P -n | grep LISTEN

## .Profile Shortcut

navbe() {

        cd $HOME/go/src/github.com/cosmos-ummc/comet

}

deploybe() {

        cd $HOME/go/src/github.com/cosmos-ummc/comet

        go build cmd/server/main.go

        cd ~

        sudo service mainbe restart

}

## important tutorials

https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-18-04

https://www.digitalocean.com/community/tutorials/how-to-install-nginx-on-ubuntu-18-04

https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04
