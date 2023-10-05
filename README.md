# Backend

## Dev environment
I use docker desktop and Tea

### Install Docker Desktop
https://www.docker.com/get-started/

### Install pkgx
https://docs.pkgx.sh/run-anywhere/terminals

## Running the API
- Setup a local instance of Postgres with dummy data
- Run the GO API server

### Setup Postgres
Make sure you have docker desktop open or running in the background and run this in your terminal of choice.

```
docker run --name db-alpha -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
```

Then within the `Backend` directory:

```
pkgx task all
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
