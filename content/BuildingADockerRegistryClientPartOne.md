---
id: "BuildingADockerRegistryClientPartOne"
title: "Building A Docker Registry Client Part One"
date: "2025-02-07"
summary: "Building a Docker Registry Client from Scratch"
tags: ["Docker Registry", "distribution", "HTTP Client", "Container", "Registry"]
---

# Building a Docker Registry v2 Client from Scratch Part One

I'm going to be building a docker registry client from scratch because it seems like a fun project. It'll also be great 
to learn some more about how containers are stored and how to manipulate their OverlayFS.

Most of the functionality will be somewhat similar to [regclient](https://github.com/regclient/regclient). If you're looking 
for an off the shelf, production ready version of what I'm building you should go try out regclient!

If you're looking to track what I'm doing in real time or if you run into any issues feel free to check out the [GitHub repository 
associated with these blog posts](https://github.com/garrett-rh/sonar). My goals with this are to get to a point where I can pull down information 
on images, delete images, and upload images.

## Getting Started

Our endpoints are defined inside of [the API spec](https://distribution.github.io/distribution/spec/api/). If it seems like 
I'm pulling endpoints and request types out of nowhere it's because I'm looking at that. 

To start, I'll just be running all of my requests against an actual docker registry that I've deployed on my local machine. 
If you're following along you can run this:

```shell
docker run -d -p 80:5000 registry:2 && \
    docker pull nginx && \
    docker tag nginx localhost/nginx && \
    docker push localhost/nginx
```

Is it the best way to do this? Nope. But It's what I'm doing for now until I start to get some things setup!

<img alt="Get'er Done" src="https://media0.giphy.com/media/v1.Y2lkPTc5MGI3NjExYmt3dHE0dGJzOXZiamw1MnBjaTZvN3R4M2R4ZnMzajVsM2h3M2p1aiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/Vh2AWuLGA1TX2MPGkn/giphy.gif" class="img-fluid">

### Making the Client

Now that I have a running registry, I'm going to make my HTTP client. I'll be starting with the following project structure.

```
├── go.mod
├── client.go
cmd
└── sonar.go
```

First, I'm going to set up my timeout and declare a new registry client. Although I don't have the new client creation wired up 
just yet, I know I'll be doing it soon so I'll just throw it in there.

`/cmd/sonar.go`

```go
package main

import (
	"github.com/garrett-rh/sonar"
	"time"
)

var version = "v2"

func main() {
	timeout := 5 * time.Second
	c := sonar.NewRegistryClient(fmt.Sprintf("http://localhost/%s", version), timeout)
}
```

Now that we have the base for our executable, we can start to create the client.

`/client.go`

```go
package sonar

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type RegistryClient struct {
	client *http.Client
	uri    string
}

func NewRegistryClient(uri string, timeout time.Duration) *RegistryClient {
	return &RegistryClient{
		client: &http.Client{
			Timeout: timeout,
		}, uri: uri,
	}
}

func (r *RegistryClient) Get(endpoint string) *http.Response {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s%s", r.uri, endpoint), nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := r.client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	return response
}
```

A quick explanation of what we have going on here.

- New Registry Client initializes our client with the timeout we specified in `/cmd/sonar.go`. It also takes in the hardcoded `http://localhost` that is in there.
  - We will change the hardcoded value in the near future, just using this for now to get things off the ground.
- We also have a wrapper over `http.client.Do()` to abstract out some of the boilerplate associated with making the request on our client.
  - This might change a bit in the future depending on our deserialization but for now it is good enough.

### /v2/ Connection Check

I'm going to start with the easiest endpoint. Just need to send a GET request to `http://localhost/v2/`.

The potential HTTP response codes are:
- 200
- 401
- 404

If we get 200 or 401, we have some actions to take. If we get 404 we will just exit there as this is only implementing the 
/v2/ endpoints and 404 means that this version of the registry doesn't support distribution API version 2. For now, I'll be focusing on our simplest use case 
and building out from there. To add this connection check in we just need to add the following into our client.

`/client.go`
```go
func (r *RegistryClient) Check() int {
	resp := r.Get("")

	return resp.StatusCode
}
```

and then add this in to `/cmd/sonar.go`:

*/cmd/sonar.go*
```go
log.Println(c.Check())
```

Now we should be able to run this with 

`go run ./cmd/sonar.go` and we will be presented with something like `2025/02/08 00:28:00 200`!

<img alt="Borat Great Success" src="https://media2.giphy.com/media/v1.Y2lkPTc5MGI3NjExa20wdXpwcWc4aXZlanBscmN6YzNxODRsdmdrNGZwbTRxOTE0MG14aSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/Od0QRnzwRBYmDU3eEO/giphy.gif" class="img-fluid">

## Part Two

In Part Two, I'll be tackling pulling down the Repositories, Repository Tags, and Manifests.