# shine

![shine logo](https://github.com/alandtsang/shine/blob/main/logo/shine-logo.png)

shine is a library written in go to create api server applications

[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

# Usage

Download and install shine:

```shell
$ go get -u github.com/alandtsang/shine/cmd/shine
```

Use shine to automatically generate the project, `pkg` is the name of the module, `author` is the name in the copyright.

```shell
$ shine -pkg myproject -author Alan
2021/03/03 10:26:40 [Info] project path: /Users/alan/myproject
```

```shell
$ cd myproject
$ go mod vendor
$ make build
```

```shell
$ ./main

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.1.17
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8080
```

Execute in another terminal:

```shell
$ curl http://127.0.0.1:8080
"hello"
```

## Help Command

```shell
$ shine -h
Usage of shine:
  -author string
    	author name (default "YourName")
  -pkg string
    	package name (default "myproject")
```
