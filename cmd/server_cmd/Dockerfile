FROM golang:1.24 AS golang-builder

WORKDIR /app
# set up an /app folder in the golang image and navigate to it

COPY go.mod go.sum ./
# copy go.mod and go.sum to current working directory within the image 
# note that the syntax is <source...> <destination> where the destination will only have path

RUN go mod download
# get all dependencies, but this part will be cached unless go.mod or go.sum changes

COPY . .
# after getting all dependencies, then copy all the rest of the files from the build context into the image

RUN go build -o server_cmd ./cmd/server_cmd
# this will now build out the binary, which will then be located in /app in the golang image

FROM debian:stable-slim
WORKDIR /app

COPY --from=golang-builder /app/server_cmd .
#remember - when copying from another image, will need to refer to its alias, and the concept of workdir only 
#applies to the current image that is currently being worked on

EXPOSE 8080
#expose port 8080   
CMD ["./server_cmd"]
#run the final command, running the binary