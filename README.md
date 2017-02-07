# Development

## Requirements
Docker

## Go Dev Server
```
cd into docker/dev
cp docker-compose-dev-template.yaml docker-compose.yaml
```

# Modify any environmental variables and/or customize docker-compose.yaml.

# Then to run it:
```
docker-compose up
```
# Navigate to http://localhost:3000/ping to see the "ping" endpoint

# Production
To release to production, we need to compile the Go server into a binary.
Then we build a Docker image with that binary included.
Changing this process to suit your needs should be possible by changing variables in *app_env.sh*.
You shouldn't need to change the build.sh or deploy.sh scripts.
```
cd docker

# Build the Go server
./build.sh

# Test the docker image to verify the Go server is running correctly
# First modify the docker-compose-test.yml so that the app will run, then:
./test.sh

# Deploy to production server
./deploy.sh
```
