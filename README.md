# go-dependency-injection-demo

Code for bytedance's article
https://www.toutiao.com/article/7146532072629207582/

* I put the html file in `docs/`

- Adaption
- Error fix
- Infra

### Feature

commit each stage, so user can tell difference between Before and After :)

### Setup

[repo](https://github.com/google/wire)

go install github.com/google/wire/cmd/wire@latest

[tutorial](https://github.com/google/wire/blob/main/_tutorial/README.md)

[adv guide](https://github.com/google/wire/blob/main/docs/guide.md)

### Tricks

1. Missing DB won't give you an error

BuildInjector must pass in db argument
I found err when use wire mode Goland plugin
however, not sure how to print error in wire.

Clues to fix comes from
[sample](https://github.com/hellokoding/golang/blob/master/rest-gin-gorm/wire.go)

In cli no error shows, no file generated.


### Related project

https://github.com/hellokoding/golang
