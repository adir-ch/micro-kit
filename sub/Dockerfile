FROM ubuntu

RUN mkdir -p /bin
COPY sub/sub.service /bin
RUN ["chmod", "a+x", "/bin/sub.service"]
ENTRYPOINT  /bin/sub.service