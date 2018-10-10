# KInspiration

The inspirational quotes for the Kontr developed in Go.

## Installation

Here are the installation instactions how to install and build from scratch

```bash
$ go get github.com/pestanko/kinspiration
```


## First Run

Here are the instructions how to run the project


### Configuration

Environment variables that can be configured

```bash
QUOTES_TOKEN=<admin_token>
QUOTES_FILE=<path_to_file>
QUOTES_HOST=<host:port>
```

- `QUOTES_TOKEN` - If it is set, the `POST` and `DELETE` methods are secured and the `Authorization` header is required.
Authorization header should be in the format: `Authorization: Bearer <QUOTES_TOKEN>`

- `QUOTES_FILE` - Location where the file with quotes are stored. If not provided, the "default path" is used.
Default path is the `quotes.json`

- `QUOTES_HOST` - Sets host and port for the router, default is the `:3000`

### Run the server
Run the server using the command

```bash
$ kinspiration
```




