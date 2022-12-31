FROM golang AS build 
WORKDIR /out
COPY . /out/
# Building for ARM, since cluster consists of Raspberry Pi's
ENV GOOS=linux
ENV GOARCH=arm
ENV GOARM=5
RUN go build -o ./discord-alerthandler 

FROM debian 
WORKDIR /app
COPY --from=build /out/discord-alerthandler .
EXPOSE 8089

ENTRYPOINT [ "/app/discord-alerthandler" ]