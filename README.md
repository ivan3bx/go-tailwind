# Go + Tailwind

This is a starting point for a simple Go application with the following traits:

- Gin web application
- Stimulus JS
- Tailwind CSS
- Procfile (for running a dev server)

The Procfile relies on 'gowatch' to reload the app server, and runs both tailwind
and esbuild to rebuild assets on the fly.

## Bootstrap

```bash
# Install gowatch for app reloading
go install github.com/silenceper/gowatch@latest

# NPM dependencies for tailwind + esbuild
yarn install

# start server
foreman start
```

## TODO

- Add configuration options
- Add Makefile (for builds)
- Add Dockerfile
