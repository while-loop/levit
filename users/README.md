Levit Users
===========

Levit Users service

Dependencies (dev only)
-----------------------

- [protoc](https://github.com/google/protobuf/releases) >= 3.4.0
- [proto-gen-go](https://github.com/golang/protpbuf)
```go get github.com/golang/protobuf/protoc-gen-go```

Installation
------------

#### Go

Note: assuming `$GOPATH/bin` is in your `PATH` env variable.

```bash
$ go get -u github.com/while-loop/levit/users/cmd/...
$ usersd
```

#### Docker

**Note**: `--net=host` is needed, because usersd binds to your Docker
host WAN/Public IP, not the container IP. This IP is what is advertised
to other services and APIs

```bash
$ docker run --net=host levitgo/users
```

Changelog
---------

The format is based on [Keep a Changelog](http://keepachangelog.com/) 
and this project adheres to [Semantic Versioning](http://semver.org/).

[CHANGELOG.md](CHANGELOG.md)

License
-------
Levit is licensed under the Apache 2.0 License.
See [LICENSE](LICENSE) for details.

Author
------

Anthony Alves
