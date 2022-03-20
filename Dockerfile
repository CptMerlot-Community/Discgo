#TODO: Make multistage to reduce image size
FROM golang:latest

RUN mkdir app

WORKDIR /app

COPY . .

RUN go build -o "discord_bot" .

# Run the outyet command by default when the container starts.
ENTRYPOINT /app/discord_bot

# Document that the service listens on port 8080.
EXPOSE 8080
