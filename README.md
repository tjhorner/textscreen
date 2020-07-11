# TextScreen

Simple Twilio app for screening people before they text you. Make your Twilio number public-facing to get rid of (most) spammers.

## Setup

First, copy `config.example.yml` to `config.yml` and fill it in with the appropriate values.

### Storage

You can store conversation sessions either in-memory or in Redis. Check the config for instructions. Note that sessions stored in Redis will expire after 24 hours to prevent them from piling up. Sessions stored in-memory do not expire.

### Docker

There is a Docker image provided for easy setup:

#### Volumes

Place your config file at `/config.yml` in the container as a volume.

#### Ports

By default, the server listens on port `8080`, so bind that to something on your host. Or whatever. You can change this in the config.

## License

idk I will set it later but I promise I won't sue you lol