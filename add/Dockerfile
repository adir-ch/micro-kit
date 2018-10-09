FROM ubuntu

RUN mkdir -p /bin
COPY add/add.service /bin
RUN ["chmod", "a+x", "/bin/add.service"]
ENTRYPOINT  /bin/add.service