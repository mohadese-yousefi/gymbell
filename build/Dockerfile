FROM golang:1.22.3-alpine

# create workspace
WORKDIR /app
COPY go.mod go.sum ./

# fetch dependancies
RUN go mod download && \
    go mod verify

# copy the source code as the last step
COPY . .

RUN go build -o ./gymbell ./cmd
#RUN go build -v -o /usr/local/bin/app ./cmd

# build binary
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/gymbell ./cmd
EXPOSE 8079
# run app
ENTRYPOINT ["./gymbell"]

# syntax=docker/dockerfile:1