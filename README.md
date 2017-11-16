# Europa
| Team | Stack | Framework |
| ---- | ----- | --------- |
| aID  | Go    | N/A       |

> Europa is the smallest of the four Galilean moons orbiting Jupiter, and the sixth-closest to the planet.
It is also the sixth-largest moon in the Solar System. Europa was discovered in 1610 by Galileo Galilei and was
named after Europa, the legendary mother of King Minos of Crete and lover of Zeus
(the Greek equivalent of the Roman god Jupiter)

Europa is a REST service responsible for storing and retrieving custom preferences (key/value) for a user.

The initial requirement for Europa came from the General Data Protection Regulation (GDPR). Users should be able
to decide what kind of data we store, and who we can share them with. Services requiring these user preferences
can use Europa as a lookup for the necessary info. But Europa should be a general purpose service for storing
preferences for a user.

[![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

## Requirements
- Low response time (should be able to query Europa on webpage load)
- Consistency and availability
- Persistence
- Preferences
    - Get all preferences for user (using tracking_key, i.e. a_user_key)
    - Get all users and all their preferences
    - Change value of a given preference for a given user
- Preferences should have default values, which should be relatively easy to change 

### Reservations
- Possibility to include the same functionality for anonymous users (based on browser key)

## Development

### Setup
This project uses [dep](https://github.com/golang/dep#setup) for dependency management.

Install project dependencies with either
```
# Make target
make deps

# or dep command
dep ensure
```

### Running Europa
To run Europa in "dev"-mode [realize](https://github.com/tockins/realize) must be installed. Realize enables automatic
recompilation of code, auto-running of tests, etc, when the code changes (using this flow is optional; it is possible to
write code and to manually compile between changes):
```
go get github.com/tockins/realize
```

Run Redis in Docker engine, either with [aid-devbox](https://github.com/amedia/aid-devbox), or using the
provided docker-compose.yml:

```
docker-compose up -d influx redis
```

After starting influx, run: `curl -i -XPOST http://localhost:8086/query --data-urlencode "q=CREATE DATABASE aid"`
This will be fixed later, so you don't have to do this manually.


Then start Europa in development mode:

```
REDIS_HOST=localhost make dev
```

### Tests

Test are run with this command: `go test ./...` in the root folder.

### Metrics

An InfluxDB reporter is set up with docker-compose on http://localhost:8086 in dev. Overide configuration with `METRICS_INFLUXDB_HOST` and `METRICS_INFLUXDB_PORT`.

## Building

Build a Docker container for Europa with

`make release`

### Service dependencies
Europa has the following dependencies:

- Redis
- InfluxDb

## Usage



### Redis
