# Use the official Postgres image as a parent image
FROM postgres:13

# Set the environment variable for the database name
ENV POSTGRES_DB=joystiq

# Add our custom schema to the Docker image
COPY 01-schema.sql /docker-entrypoint-initdb.d/
COPY 03-init.sql /docker-entrypoint-initdb.d/
