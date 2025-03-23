#!/bin/sh -e
    
podman build -t hello-event-sourcing-with-go-devenv --target devenv .

# Create a named volume for the home directory
if ! podman volume ls | grep -q "devenv-root-home"; then
    podman volume create devenv-root-home
fi

exec podman run -it --replace \
    --name hello-event-sourcing-with-go-devenv \
    --hostname hello-event-sourcing-with-go-devenv \
    -v claude-root-home:/root \
    -v ./:/mnt/hello-event-sourcing-with-go/ \
    --workdir /mnt/hello-event-sourcing-with-go \
    hello-event-sourcing-with-go-devenv \
    fish
