# Gopher translator

## Start application

> **`make start-dev`**<br>
> This command starts the application.<br>

> **`PORT=8081`**<br>
> If you want to specify a port add this to the make start-dev command. The default port is 8080.

> **`go run cmd/server/main.go`**<br>
> Alternatively if you want to start directly with golang

> **`-port 8081`**<br>
> If you want to specify a port add this to the go run command. The default port is 8080.

**NOTE:** *If you use the go command directly you need to start your redis server manualy*

## API Routes

### GET Routes
> **`/`**<br>
> Returns basic information for API

> **`/history`**<br>
> Returns all translations available in storage

### POST Routes
> **`/word`**<br>
> Returns the translation of a single word

> **`/sentence`**<br>
> Returns the translations of a sentence