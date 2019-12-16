FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
EXPOSE 4040
CMD ["go", "run", "main.go", "redis"]

