# GeoIP Server
*A microservice for MaxMinds GeoIP2 databasese*

**IMPORTANT**: This currently only works with GeoIP2 Lite databases, as I don't have access to the paid ones.

## Building
To build this project, you need my custom Docker build extension dockpipe, which you can get here [here](https://github.com/lorenz/dockpipe).
Then you can just type
```
$ dockpipe geoip-server:dev .
```
to get an image built.

## Setup
Before you can launch this server, you need to set up [lorenz/geoip-updater](https://github.com/lorenz/geoip-updater).
Have a look at the readme of that project for setup instructions on the updater.
Once you have an instance of that service running, you can start up one or more instances of this.

## Volumes
| Path | Description |
| ---- | ----------- |
| `/data` | Shared volume with geoip-updater for storing GeoIP-Databases. Should be on SSD or high-IOPS volume. |