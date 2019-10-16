# r2s
[![Build Status](https://travis-ci.org/Ahton89/r2s.svg?branch=master)](https://travis-ci.org/Ahton89/r2s)

Tool for copy hashes from redis to redis

## Usage

### Run

#### Options
```
Usage of ./r2s:
   -h  --help           Print help information
   -p  --production     Redis production node <host:port>
   -s  --sandbox        Redis sandbox node <host:port>
   -i  --production-db  Redis production db <num>. Default: 0
   -o  --sandbox-db     Redis sandbox db <num>. Default: 0
   -d  --debug          Enable debug mode
```

#### Example

```
./r2s_darwin_64 -p 127.0.0.1:6379 -i 3 -s 127.0.0.1:6379 -o 4
```