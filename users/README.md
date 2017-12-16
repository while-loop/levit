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

```bash
$ go get -u github.com/while-loop/levit/users/cmd/...
```

Usage
-----

Note: assuming `$GOPATH/bin` is in your `PATH` env variable.

```sh
$ usersd
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
