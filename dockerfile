FROM golang:1.22

ENV CGO_ENABLED 0
ENV GOOS "linux"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download;
COPY Makefile Makefile
RUN make install-swag; make install-goose;
COPY cmd cmd
COPY controllers controllers
COPY domain domain
COPY lib lib
COPY internal internal
COPY .env .env
COPY entrypoint.sh entrypoint.sh
RUN chmod +x entrypoint.sh

RUN bin/swag init -g cmd/main.go; CGO_ENABLED=$CGO_ENABLED GOOS=$GOOS go build -o /val cmd/main.go

EXPOSE 8080

# Run
ENTRYPOINT ["./entrypoint.sh"]