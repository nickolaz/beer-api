FROM golang:1.17.5

RUN mkdir "app"
ADD . /app
## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

WORKDIR /app

RUN make build
EXPOSE 8080
## Launch the wait tool and then your application
CMD /wait && ./beer-api
