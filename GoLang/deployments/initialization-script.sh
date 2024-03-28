#!/bin/bash

# Stop execution if any command fails
set -e

# Update and install necessary packages
echo "Updating package list..."
sudo apt-get update -y

echo "Installing necessary digitalocean packages..."
wget -qO- https://repos-droplet.digitalocean.com/install.sh | sudo bash

# Check if Docker is installed, install if it's not
if ! command -v docker &> /dev/null
then
    echo "Docker is not installed. Installing Docker..."
    sudo apt-get install -y docker.io
else
    echo "Docker is already installed."
fi

# Ensure Docker service is running
sudo systemctl start docker
sudo systemctl enable docker

# Pull the Docker image
echo "Pulling the Docker image..."
sudo docker pull jake4/loadtest-test-api:latest

echo "Running the Docker container..."
sudo docker run -d -e REST_PORT=80 -p 80:80 jake4/loadtest-test-api:latest

echo "Container deployed successfully!"