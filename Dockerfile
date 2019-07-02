FROM golang

ADD watchdog /bin/watchdog

ENV NAME=world
ENTRYPOINT ["/bin/watchdog"]


