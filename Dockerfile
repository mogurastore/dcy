FROM busybox

COPY ./src /dcy-cli

ENTRYPOINT ["./dcy-cli/dcy"]
