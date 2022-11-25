FROM golang:1.19.3

RUN mkdir /app
ADD . /app
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

RUN go build -o main .
CMD ["/app/main"]