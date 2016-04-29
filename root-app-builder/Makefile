NAME = root-app-builder
IMAGE_NAME = agilestack-$(NAME)

SWAGGER_SPEC = swagger.json
GO_FILES=*.go


############################
#          BUILD           #
############################

# Default goal: build all, up to the docker image
install : docker-build

# Generate the docker image
docker-build : js-build go-build
		docker build -t $(IMAGE_NAME) .

# Builds the JavaScript application.
#
# If there is a "build" directory, then nothing is done.
js-build: node_modules
		npm run build

# Starts a Webpack dev server
dev-server  : node_modules
		npm run start

# Installs the NodeJS dependencies.
#
# If there is already a "node_modules" directory, then this goal is skipped.
# If there is no "package.json" file, then an error is thrown.
node_modules : package.json
		npm update
		touch node_modules

# Compiling the Go code
go-build : $(NAME)

# Compiling the main Go executable
$(NAME) : $(GO_FILES)
		env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(NAME)


############################
#          SETUP           #
############################

setup: go-deps

# Go dependencies
go-deps : .
		go get -u -t $(shell go list ./... | grep -v /vendor/)


############################
#           TEST           #
############################

test : go-test

# Run the go tests
go-test:
		# in test
		go test -v -p 1 $(shell go list ./... | grep -v /vendor/)


############################
#          DEPLOY          #
############################

# Publish the docker image
docker-deploy :
		docker tag $(IMAGE_NAME) eogile/$(IMAGE_NAME):latest && docker push eogile/$(IMAGE_NAME):latest


############################
#      DOCUMENTATION       #
############################

docs: swagger

# Generates the Swagger specification file
swagger :
		# Generating the Swagger specification file
		go get -u github.com/go-swagger/go-swagger/cmd/swagger
		swagger generate spec -o $(SWAGGER_SPEC)


############################
#          CLEAN           #
############################

clean :
		$(RM) -r $(NAME) $(SWAGGER_SPEC)

.PHONY : install docker-build js-build go-build setup go-deps test go-test docker-deploy docs swagger clean
