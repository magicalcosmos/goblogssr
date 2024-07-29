FROM alpine:latest
WORKDIR /
COPY shareus /opt/shareus
COPY conf /opt/conf
COPY client/dist* /opt
CMD [ "/opt/shareus" ]

EXPOSE 8080
