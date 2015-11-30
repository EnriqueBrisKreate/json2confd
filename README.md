# json2confd
Stores configuration on confd backends from a json file.

[![Build Status](https://travis-ci.org/creativedrive/json2confd.svg)](https://travis-ci.org/creativedrive/json2confd)
[![Coverage Status](https://coveralls.io/repos/creativedrive/json2confd/badge.svg?branch=master&service=github)](https://coveralls.io/github/creativedrive/json2confd?branch=master)

# Installation

### Binary Download

Currently json2confd ships binaries for OS X and Linux 32 and 64bit systems. You can download the latest release from [GitHub](https://github.com/creativedrive/json2confd/releases)

#### OS X

```
wget -O /usr/local/bin/json2confd https://github.com/creativedrive/json2confd/releases/download/0.1/json2confd_darwin_amd64
chmod +x /usr/local/bin/json2confd
```

#### Linux

```
wget -O /usr/local/bin/json2confd https://github.com/creativedrive/json2confd/releases/download/0.1/json2confd_linux_amd64
chmod +x /usr/local/bin/json2confd
```


### Basic usage
Given the following json file:
```
{
  "elasticsearch": {
    "host": "example.com",
    "port": 9200,
    "prefix": "some_prefix"
  },
  "allow_root_login": true,
  "redis_hosts": [
    "127.0.0.1",
    "somehost.com",
    "somehost2.com"
  ]
}
```
After running:
```
json2confd --file some_file.json --backend redis --node 127.0.0.1:6379
```

The redis backend will have the following keys:
```
"/elasticsearch/port" "9200"
"/elasticsearch/prefix" "some_prefix"
"/allow_root_login" "true"
"/redis_hosts" "[\"127.0.0.1\",\"somehost.com\",\"somehost2.com\"]"
"/elasticsearch/host" "example.com"
```


### Testing
To run all tests:
```
go test
```
To run a specific test:
```
go test -check.f "TestFlattenJsonStrJson1"
```

