FROM alpine:3.6

# set timezone
RUN apk --no-cache --update add ca-certificates

ADD guestbook /root/
RUN mkdir -p rest/swagger
ADD rest/swagger/swagger.json rest/swagger/swagger.json

ENV GUESTBOOK_SOCKET tcp://0.0.0.0:8080

EXPOSE 8080

ENTRYPOINT [ "/root/guestbook", "run" ]
