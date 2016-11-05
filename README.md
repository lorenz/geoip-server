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

### Volumes
| Path | Description |
| ---- | ----------- |
| `/data` | Shared volume with geoip-updater for storing GeoIP-Databases. Should be on SSD or high-IOPS volume. |

### Ports
| Port | Protocol | Description |
| ---- | -------- | ----------- |
| 4000 | HTTP | HTTP Server for GeoIP requests |

## How to use
Queries are basically just HTTP requests to http://container_ip:4000/8.8.8.8 for an IP of 8.8.8.8.
The answers are JSON-encoded and follow the same format as the original GeoIP entries.

Examples:
```
$ curl http://172.17.0.23:4000/8.8.8.8
{"City":{"GeoNameID":5375480,"Names":{"de":"Mountain View","en":"Mountain View","fr":"Mountain View","ja":"マウンテンビュー","ru":"Маунтин-Вью","zh-CN":"芒廷维尤"}},"Continent":{"Code":"NA","GeoNameID":6255149,"Names":{"de":"Nordamerika","en":"North America","es":"Norteamérica","fr":"Amérique du Nord","ja":"北アメリカ","pt-BR":"América do Norte","ru":"Северная Америка","zh-CN":"北美洲"}},"Country":{"GeoNameID":6252001,"IsoCode":"US","Names":{"de":"USA","en":"United States","es":"Estados Unidos","fr":"États-Unis","ja":"アメリカ合衆国","pt-BR":"Estados Unidos","ru":"США","zh-CN":"美国"}},"Location":{"AccuracyRadius":1000,"Latitude":37.386,"Longitude":-122.0838,"MetroCode":807,"TimeZone":"America/Los_Angeles"},"Postal":{"Code":"94035"},"RegisteredCountry":{"GeoNameID":6252001,"IsoCode":"US","Names":{"de":"USA","en":"United States","es":"Estados Unidos","fr":"États-Unis","ja":"アメリカ合衆国","pt-BR":"Estados Unidos","ru":"США","zh-CN":"美国"}},"RepresentedCountry":{"GeoNameID":0,"IsoCode":"","Names":null,"Type":""},"Subdivisions":[{"GeoNameID":5332921,"IsoCode":"CA","Names":{"de":"Kalifornien","en":"California","es":"California","fr":"Californie","ja":"カリフォルニア州","pt-BR":"Califórnia","ru":"Калифорния","zh-CN":"加利福尼亚州"}}],"Traits":{"IsAnonymousProxy":false,"IsSatelliteProvider":false}}
```
