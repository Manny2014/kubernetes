FROM golang:latest 
RUN mkdir /app 
ADD ./server /app/
WORKDIR /app 
RUN go build -o app 
CMD ["/app/app"]