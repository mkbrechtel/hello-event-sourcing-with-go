FROM debian:trixie as env

RUN apt-get update && apt-get install -y golang ca-certificates fish && apt-get clean

#WORKDIR /opt/hello-event-sourcing-with-go

#COPY go.mod go.sum ./
#RUN go mod download


FROM env as devenv

# Install claude-code for development
RUN apt-get update && apt-get install -y nodejs npm
RUN npm install -g @anthropic-ai/claude-code


#FROM env as build

