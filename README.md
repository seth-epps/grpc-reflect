# grpc-reflect

grpc-reflect is a simple tool used to list services using [grpc reflection](https://github.com/grpc/grpc/blob/9a12ec91e10c555cadc42893983a3e8ac30eb805/src/proto/grpc/reflection/v1/reflection.proto#L31) for specified server.

## Usage
By default the target will be `localhost:9090` however you can specify your own target as an argument.

### Local
```sh
❯ make build
CGO_ENABLED=0 go build -o dist/reflect

❯ ./dist/reflect
|------------|
|- Services -|
|------------|
grpc.reflection.v1.ServerReflection
grpc.reflection.v1alpha.ServerReflection
main.Hello
```

### Docker
```sh
❯ docker image build -t grpc-reflect .

❯ docker container run --rm grpc-reflect
usage: /reflect [target]

list grpc server services for a given target (default `localhost:9090`)

❯ docker container inspect hello-go -f '{{index .NetworkSettings.Networks.bridge.IPAddress}}'
172.17.0.2

❯ docker container run --rm grpc-reflect 172.17.0.2:9090
|------------|
|- Services -|
|------------|
grpc.reflection.v1.ServerReflection
grpc.reflection.v1alpha.ServerReflection
main.Hello
```