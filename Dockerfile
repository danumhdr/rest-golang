# Start from golang base image
FROM golang:1.18-alpine as builder

# Set the current working directory inside the container 
RUN mkdir /build 
ADD . /build/
WORKDIR /build

# WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum /

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

#RUN go build -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

#Command to run the executable
#CMD ["go","run","."]

FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]