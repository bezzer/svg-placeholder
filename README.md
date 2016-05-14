# SVG Placeholder

## About
Small service writen in Go to request svg placeholder images. The Go app is statically compiled and added to an empty docker container for running.

## How to build
The file `build.sh` builds the go app as a static binary, then adds it to a scratch docker container along with any app resources in `www-dist/`.

The build steps are:

1. Build the svg-placeholder Go binary
2. Minify static resources and output to `www-dist/`
3. Build the docker container (tagged as `bezzer/svg-placeholder`) 

## License
MIT
