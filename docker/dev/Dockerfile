# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM golang:1.7

# Create a directory inside the container to house our application and then make it the working directory.
RUN mkdir -p /scripts
WORKDIR /go/src

#copy a stub file in so this container will build
COPY app /go/src/app

# Download and install any required third party dependencies into the container.
RUN go get github.com/codegangsta/gin
RUN go-wrapper download app
RUN go-wrapper install app

# Set the PORT environment variable inside the container
ENV PORT=8080

#where gin will watch for code changes by default
ENV GINPATH="/go/src/app"

#Expose port 3000 to the host so we can access the gin proxy
EXPOSE 3000

#This script allows us to set compose-time env vars to pass to gin for our project
COPY entrypoint.sh /scripts/entrypoint.sh

#run gin/our code
ENTRYPOINT ["/scripts/entrypoint.sh"]