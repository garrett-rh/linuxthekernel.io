# linuxthekernel.io

This repository contains all the stuff needed to run my personal blog/website.

Right now it is pretty barebones, but I'll be incrementally adding to it over time.

## Running

Local development can occur via running `docker compose --profile dev up --build`.
This will start a docker container, exposing port 80.

## Production

Releasing into production is cloning the repo down to the production server.
To build the production containers, run `docker compose --profile prod build` from the root of the repo.

### Caveats

Currently, do not have hot-reload working in a container for frontend work.