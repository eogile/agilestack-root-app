Root application builder
--------------------

The main goals of this plugin are :

* to expose the REST API used by the root application (to get menu entries for example),
* to expose the REST used by other plugins to record theirs React routes and reducers,
* to compile the JavaScript application when it's necessary.

The REST APIs are documented using [Swagger](http://swagger.io).

### Working with this application

#### Loading the dependencies and compiling the Go code

```bash
make compile
```

Please notice that : 

* the dependencies will be loaded only the first time,
* if a dependency was previously loaded, then it won't be updated.

#### Launching the Go test

```bash
make test
```

#### Generating the Docker image

```bash
make 
```

If the Go code was already compiled, then it won't be recompiled.

### Swagger documentation generation

The Swagger documentation is generated directly from the Go code using the [go-swagger](https://github.com/go-swagger/go-swagger) tool.

Once go-swagger is installed, the documentation can be generated using the following command line :
```bash
go get github.com/go-swagger/go-swagger/cmd/swagger
swagger generate spec -o swagger.json
```

The documentation can also generated with the command line :
```bash
make swagger
```

### Working with the JavaScript application

#### Launching the application with Webpack 

A Webpack dev server can be starting using the following command line :
```bash
make dev-server
```

#### Installing the node dependencies and compiling the application
```bash
make js-build
```