# Emergence Engineering Go Service template 

## Module

### Init new Go module
The project uses Go modules, I suggest to delete the `go.mod` and `go.sum` files in the project
root and initiate your own project by the following command
```bash
go mod init github.com/account/project
```
### Init your project
There are a few things that has to be done to take ownership of the template
- Init your brand new Go module (described above)
- Replace the text "example-service" in the `Dockerfile` and the `Makefile`
- Adjust your `ldflags` (`LINKFLAGS`) in Makefile

## Build
We use Makefiles to build our project. The Makefile is located in the project root

The template has two commands:
```bash
make build
```
Builds the service, injects the build times constants to the binary and links the packages.
The go tool will do all the above.

```bash
make test
``` 
//TODO

### Build time variables
By default the following information is added build time: 
* `SERVICE_NAME`
* `GIT_COMMIT_ID`
* `GIT_BRANCH`
* `BUILD_TIMESTAMP`

The above variables can be used in health-checks, microservice debugging or service discover...

#### Custom build time variables
You can add more build time constants in the makefile. Make sure that you are adding
the build time constant to the correct file.
  
 
### Multi-stage build
We are using a [multistage build pattern](https://docs.docker.com/develop/develop-images/multistage-build/)
for security and for small image size as well.
Our builder image is based on `golang:1.13` and our service image is `alpine:latest`.

### Service
To run your service use:
```bash
./example-service start
```

To get help
```bash
./example-service --help 
```

#### Service CLI
We are using https://github.com/urfave/cli for the service CLI. CLI flags are configured [in the service file](/internal/service/main.go).

#### Business logic
You can add your business logic [here](/internal/service/main.go:79)

### Graceful shutdown
By default the service is listening for `SIGTERM` signals from the OS. If one occurs a graceful
shutdown is initialized. You can extend the graceful shutdown process in the 
`handleShutdownGracefully` function inside the [service main file](/internal/service/main.go).

### Logging
Logging is currently done with github.com/sirupsen/logrus
