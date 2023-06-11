# RARBG selfhosted

Currently consists of a Torznab API that can be used with the [Servarr stack](https://wiki.servarr.com/).

This is pretty basic at the moment and unfinished in places, but it basically works (tested in Prowlarr, Radarr and Sonarr), I might embellish it if there's any interest. If you'd like to open a PR then even better!

## Quick start

- Obtain a copy of the RARBG SQLite database, I can't help you with that but it's floating around!
- Copy the `docker-compose.example.yml` to `docker-compose.yml` and adapt to your needs. At a minimum you'll need to point the first volume to the path of your SQLite file.
- Run `docker-compose up -d`.
- The torznab endpoint is now exposed at `http://localhost:3333/torznab`.
- In [Prowlarr](https://wiki.servarr.com/prowlarr), [Radarr](https://wiki.servarr.com/radarr) or [Sonarr](https://wiki.servarr.com/sonarr), you can now add RARBG as a Generic Torznab indexer.
- If you want to `git pull` updates from this repo, you'll then need to rebuild with `docker-compose down` then `docker-compose up -d --build`.

Alternatively, install GoLang and run directly: `PATH_SQLITE_DB=/path/to/rarbg_db.sqlite go run .`
