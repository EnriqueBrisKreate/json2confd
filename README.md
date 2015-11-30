# json2confd
Stores configuration on confd backends from a json file.

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

