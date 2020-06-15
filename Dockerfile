FROM alpine

COPY ./src /dcy-cli

ENTRYPOINT ["./dcy-cli/dcy"]
