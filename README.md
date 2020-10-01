# space-go-api
This is a basic API made in go to test the Gin Router.

## Installation
Run the following commands in a unix environment (or whatever the Windows equivalent is)
```shell script
git clone https://github.com/bradenn/space-go-api.git
cd space-go-api
npm install
```  
## Usage

Here is a basic Dockerfile if you are not familiar with Node's integration with Docker.
```shell
go run main.go
```

## Structure
```
src
├── Config
│   ├── database.go
│   └── environment.go
├── Controller
│   ├── moons.go
│   ├── planets.go
│   ├── stars.go
│   └── systems.go
├── Models
│   ├── moon.go
│   ├── planet.go
│   ├── star.go
│   └── system.go
├── Router
│   └── routes.go
└── main.go

```
## API Reference
Here are some basic requests to get you on the right path...

Get [system/star/planet/moon]
```http request
POST localhost/api/v1/planets
```
```json
[{
        "name": "Jastreb Sector PT-R b4-8",
        "id": "daa262b151",
        "x": "-32.25",
        "y": "7.0625",
        "z": "2.5",
        "starId": "404e86b7c5"
    }, ...]
```

Get [system/star/planet/moon]
```http request
POST localhost/api/v1/planets/404e86b7c5
```
```json
{
        "name": "Jastreb Sector PT-R b4-8",
        "id": "daa262b151",
        "x": "-32.25",
        "y": "7.0625",
        "z": "2.5",
        "starId": "404e86b7c5"
}
```

## Contributing
Pull requests are always welcome.
Please open an issue before making any drastic changes.

## Licence
 [MPL-2.0](https://choosealicense.com/licenses/mpl-2.0/)


