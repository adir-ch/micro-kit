FROM ubuntu

RUN mkdir -p /bin
COPY calc/calc.service /bin
RUN ["chmod", "a+x", "/bin/calc.service"]
ENTRYPOINT  /bin/calc.service