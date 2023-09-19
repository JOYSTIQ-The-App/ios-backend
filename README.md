# Backend

# Tea
The backend uses Tea which is a new package manager meant to make the lives of developers easier

## Installing Tea
`sh <(curl https://tea.xyz)`

Install tea BUT DO NOT setup magic, you'll know when you see it

## My Setup
I have docker desktop
https://www.docker.com/get-started/

I have Tea installed

## Running the API
- Setup a local instance of Postgres with dummy data
- Run the GO API server

### Setup Postgres
Make sure you have docker desktop open or running in the background

```
docker run --name db-alpha -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
```

Then within the `Backend` directory:

```
tea task all
```

This command will create a "joystiq" database, add the schema, and add dummy data so it will be good to go for testing.

### Run the API server
Change directories to the `api` directory and run:

`tea go run server.go`
