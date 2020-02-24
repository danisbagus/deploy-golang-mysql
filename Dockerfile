FROM danisbagus/base-go:gomod

RUN mkdir -p /go

COPY . /go

EXPOSE 80

WORKDIR /go

RUN go mod init go-project

RUN go get -u github.com/joho/godotenv
RUN go get -u github.com/go-sql-driver/mysql

CMD ["go","run", "main.go"]
