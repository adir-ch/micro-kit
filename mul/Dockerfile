FROM ubuntu

RUN mkdir -p /bin
COPY mul/mul.service /bin
RUN ["chmod", "a+x", "/bin/mul.service"]
ENTRYPOINT  /bin/mul.service