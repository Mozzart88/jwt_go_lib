# jwt_go_lib
Java Web Token implementation

## Quick start
To install dependencies
```bash
go mod download && go mod verify
```

To create some tokens whith preseted values
```bash
go run main.go
```

If you haven't install `golang` yet, try `dokcker`
```bash
docker build -t jwt:0.1 .
docker run -ti --rm jwt:0.1
```

## Additional information
In this lib is used only `SHA256` algorithm for this moment
