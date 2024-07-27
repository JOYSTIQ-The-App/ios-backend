# Table of Contents
1. [Ubuntu Server Setup](#ubuntu-server-setup)
2. [Application Deployment](#application-deployment)

# Ubuntu Server Setup 

## System Update and Upgrade
```sh
sudo apt update
sudo apt upgrade -y
```

## Verify and Install Docker and Git
```sh
docker --version
git --version
```

## Install Taskfile
```sh
sudo snap install task --classic
task --version
```

## Setup SSH Key Pair for Git
```sh
git config --global user.name "Your Name"
git config --global user.email "your_email@example.com"

ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
# Copy the SSH key to your GitHub account under Settings > SSH and GPG keys > New SSH key

git clone git@github.com:your_username/your_repository.git
```

## Enable Firewall and Set Up Ports
```sh
sudo apt install ufw -y

sudo ufw allow 22/tcp
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp

sudo ufw enable

sudo ufw status
```

# Application Deployment

## Running the API
- Setup a local instance of Postgres with dummy data
- Run the GO API server

### Setup Postgres
Make sure you have docker desktop open or running in the background and run this in your terminal of choice.

```
docker run --name <db name> -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
```

Then within the `Backend` directory:

```
task all
```

This command will create a "joystiq" database, add the schema, and add dummy data so it will be good to go for testing.

To make it even cleaner you can add an alias to your .zshrc file.

#### Add the Alias

Create the file in your home directory if it doesn't exist or open the file if it does.

`nano ~/.zshrc`

Add the new alias and save your changes

```
alias tt="tea task"
```

`source ~/.zshrc`

Now the command becomes

```
tt all
```

Read `Taskfile.yml` to see all the other commands available.

### Run the API server
Change directories to the `api` directory and run:

```
tea go run server.go
```
