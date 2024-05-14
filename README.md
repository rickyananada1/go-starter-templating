## Apa itu Go-Starter?
Tools yang mengotomatisasi pembuatan struktur project REST API dengan menggunakan bahasa pemrograman Go

## go-starter command
```sh
 go-starter help

	┌─┐┌─┐   ┌─┐┌┬┐┌─┐┬─┐┌┬┐┌─┐┬─┐
	│ ┬│ │───└─┐ │ ├─┤├┬┘ │ ├┤ ├┬┘
	└─┘└─┘   └─┘ ┴ ┴ ┴┴└─ ┴ └─┘┴└─
	 
Usage: go-starter [OPTIONS]

Options:
  -f string
    	starter configuration file (default "starter.yaml")
  help
    	show help message

Examples:
  go-starter new      Generate configuration file
  go-starter          Run the application
  go-starter -f=configuration.yaml[default: starter.yaml] Specify a custom configuration file
```
## generate starter configuration
```sh
go-starter new

	┌─┐┌─┐   ┌─┐┌┬┐┌─┐┬─┐┌┬┐┌─┐┬─┐
	│ ┬│ │───└─┐ │ ├─┤├┬┘ │ ├┤ ├┬┘
	└─┘└─┘   └─┘ ┴ ┴ ┴┴└─ ┴ └─┘┴└─
	 
```

`starter.yaml` is a go-starter configuration:
```yaml
version: "1"
package: nama-app
database: database
```
database support :
- mysql
- mongodb
- postgres

## generate structure project
```sh
 go-starter

	┌─┐┌─┐   ┌─┐┌┬┐┌─┐┬─┐┌┬┐┌─┐┬─┐
	│ ┬│ │───└─┐ │ ├─┤├┬┘ │ ├┤ ├┬┘
	└─┘└─┘   └─┘ ┴ ┴ ┴┴└─ ┴ └─┘┴└─
	 
layer 'cmd' created successfully.
layer 'api' created successfully.
layer 'db/migrations' created successfully.
layer 'internals/config' created successfully.
layer 'internals/delivery/http' created successfully.
layer 'internals/delivery/messaging' created successfully.
layer 'internals/gateway' created successfully.
layer 'internals/models' created successfully.
layer 'internals/repository' created successfully.
layer 'internals/usecase' created successfully.
layer 'internals/helpers' created successfully.
layer 'tests' created successfully.
```

## Structure
Structure reference
- [Golang standards project layout](https://github.com/golang-standards/project-layout/)
```sh
├── api
├── cmd
│   └── main.go
├── config.dev.yaml
├── db
│   └── migrations
├── docker-compose.yaml
├── go.mod
├── internals
│   ├── config
│   │   ├── mysql.go ## filename following the database
│   │   └── viper.go
│   ├── delivery
│   │   ├── http
│   │   └── messaging
│   ├── gateway
│   ├── helpers
│   │   └── error.go
│   ├── models
│   ├── repository
│   └── usecase
├── starter.yaml
└── tests

```

## About go-starter
This tool adopt \
[gorm](https://gorm.io/)\
[mongo client](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo)\
[viper](https://pkg.go.dev/github.com/dvln/viper)

Database support:
| Database              | Support |
| :---------------- | :------: |
| MySQL        |   ✅   |
| PostgreSQL           |   ✅   |
| MongoDB    |  ✅   |

## Installation
```sh
go install github.com/rickyananada1/go-starter-templating@latest
```

## Running Your App
Setelah generate template, masuk ke folder yang telah di generate oleh `go-starter`, kemudian tambahkan dependencies dengan command berikut.
```sh
go mod tidy
```
command ini akan mendownload dan menginstall dependencies yang dibutuhkan untuk project anda\
Setelah berhasil menginstall dependencies, jalankan aplikasi dengan command :
```sh
go run cmd/main.go
```
## Authors

- [@rickyananada1](https://github.com/fanchann)

