#!/bin/bash

# Start the database service
docker-compose up -d db

# Wait until PostgreSQL is ready
while ! docker-compose exec -T db pg_isready -U postgres; do
    echo "Waiting for PostgreSQL..."
    sleep 1
done

# Start the API service
docker-compose up -d api

echo "All services started!"
