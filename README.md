# RARBG selfhosted

Currently consists of a Torznab API that can be used with the [Servarr stack](https://wiki.servarr.com/).

This is pretty basic at the moment and unfinished in places, but it basically works (tested in Prowlarr and Radarr), I might embellish it if there's any interest. If you'd like to open a PR then even better!

## Quick start

- Obtain a copy of the RARBG SQLite database, I can't help you with that but it's floating around!
- Copy the `docker-compose.example.yml` to `docker-compose.yml` and adapt to your needs. At a minimum you'll need to point the first volume to the path of your SQLite file.
- Run `docker-compose up -d`.
- The torznab endpoint is now exposed at `http://localhost:3333/torznab`.
- In [Prowlarr](https://wiki.servarr.com/prowlarr), you can now add RARBG as a Generic Torznab indexer.
