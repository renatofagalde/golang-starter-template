#!/bin/bash

echo "Parando todos os containers..."
docker stop $(docker ps -aq) 2>/dev/null

echo "Removendo todos os containers..."
docker rm -f $(docker ps -aq) 2>/dev/null

echo "Removendo todas as imagens..."
docker rmi -f $(docker images -aq) 2>/dev/null

echo "Removendo todos os volumes..."
docker volume rm $(docker volume ls -q) 2>/dev/null

echo "Removendo todas as redes customizadas..."
docker network rm $(docker network ls | grep -v "bridge\|host\|none" | awk '{ print $1 }') 2>/dev/null

echo "Limpando build cache..."
docker builder prune -af --filter "until=0h" 2>/dev/null

echo "Docker zerado com sucesso."
