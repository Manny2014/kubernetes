FROM golang:latest 
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN go build -o function 
CMD ["/app/function"]