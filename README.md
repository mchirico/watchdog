


[![Build Status](https://travis-ci.org/mchirico/watchdog.svg?branch=master)](https://travis-ci.org/mchirico/watchdog)
[![Build Status](https://mchirico.visualstudio.com/watchdog/_apis/build/status/mchirico.watchdog?branchName=master)](https://mchirico.visualstudio.com/watchdog/_build/latest?definitionId=1&branchName=master)
[![codecov](https://codecov.io/gh/mchirico/watchdog/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/watchdog)
# watchdog

## Build with vendor
```
export GO111MODULE=on
go mod init
# Below will put all packages in a vendor folder
go mod vendor



go test -v -mod=vendor ./...

# Don't forget the "." in "./cmd/script" below
go build -v -mod=vendor ./...
```


## Don't forget golint

```

golint -set_exit_status $(go list ./... | grep -v /vendor/)

```


