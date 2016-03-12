## About

Companion package to https://github.com/gologs/log.

Implementation of gologs/log/levels [Interface](https://godoc.org/github.com/gologs/log/levels#Interface)
that pipes messages to [glog](https://godoc.org/github.com/golang/glog).

## Demo

```sh
$ go run cmd/demo/demo.go -logtostderr -v=1
I0312 15:54:25.443002   21224 demo.go:31] this is debug
I0312 15:54:25.443052   21224 demo.go:32] this is info
W0312 15:54:25.443058   21224 demo.go:33] this is warn
E0312 15:54:25.443064   21224 demo.go:34] this is error
F0312 15:54:25.443069   21224 demo.go:35] this is fatal
exit status 1
```
