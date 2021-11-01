# url-shortener
A simple demo written in Golang to shorten URL and redirect

```
.
├── go.mod
├── go.sum
├── handler
│   └── handlers.go
├── main.go
├── shortener
│   ├── shorturl_generator.go
│   └── shorturl_generator_test.go
└── store
    ├── store_service.go
    └── store_service_test.go
```
## background
These small utilities services are quite helpful in reducing noise when sharing links. You can input a very long URL address and it can output a much shorter version of it much easier to share around.

## setup envs
- go 1.17
- redis 3.2

## how it works
- storage layer
- short link generator
  - using SHA256 to hash the initial inputs
  - using BASE58 binary to text encoding
- forwarding
  - get the creation request body, parse it and extract the initial long url and userId
  - generate our shortened hash using `shortener.GenerateShortLink()`
  - store the mapping of output `hash/shortUrl` with the initial long url using `store.SaveUrlMapping()`

## reference
[Let's build a URL shortener in Go - EDDY WANNY M](https://www.eddywm.com/lets-build-a-url-shortener-in-go/)
