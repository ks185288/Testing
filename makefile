BUILD_DIR=out

# Name used in REST API generation. Used for naming controllers
REST_API_NAME=test

# Version tag used to lock in the go-swagger code generation version
SWAGGER_VERS=v0.28.0
# API Name - used for naming controllers/handlers
SWAGGER_API=$(REST_API_NAME)
# Directory to generate swagger code into
SWAGGER_DEST=./swagger/gen
# Swagger Specification file location
SWAGGER_SPEC=./api/swagger.yml
# Generated documentation
API_DOCS=docs/apispec

# Generate runs code generation tools to populate framework
generate: swagger-server

swagger-server:
	mkdir -p $(SWAGGER_DEST)
	swagger generate server -t $(SWAGGER_DEST) -f $(SWAGGER_SPEC) --exclude-main -A $(SWAGGER_API)

docs: docs/apispec/markdown docs/apispec/html
