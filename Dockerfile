FROM golang:1.18.3
# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/loftwah/article-api
# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .
# Download all the dependencies
RUN go get -u github.com/gorilla/mux
RUN go build -o article-api
# This container exposes port 8080 to the outside world
EXPOSE 8000
# Run the executable
CMD ["./article-api"]