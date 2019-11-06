FROM golang:latest 
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN go get -u && go build -o main main.go 
CMD ["/app/main"]