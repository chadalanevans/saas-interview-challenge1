FROM golang:latest 
RUN mkdir /app 

#Bring in the go project
COPY bin/. /app/ 
WORKDIR /app 

