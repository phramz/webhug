Webhug 🤗
==========

A generic webhook dispatcher.

## Build & Run

``` 
git clone https://github.com/phramz/webhug.git
make build
./webhug

> INFO : 2020/01/25 18:25:03.520102 webhug.go:16: reading config ...
> INFO : 2020/01/25 18:25:03.521061 webhug.go:20: setting up webhook 'example' at path '/example'
> INFO : 2020/01/25 18:25:03.521134 webhug.go:37: 🤗 webhug listening on :8080 ...
```

With the default config you should now be able to try this:
``` 
curl -H 'x-auth-token: top secret' -X POST -d '{"some": ["random", "json"]}' http://localhost:8080/example

> [example] {"some": ["random", "json"]}HOSTNAME=fbc56923c4e7
> [example] SHLVL=1
> [example] HOME=/root
> [example] WEBHUG_WEBHOOK=example
> [example] WEBHUG_REQUEST_HEADER={"Accept":["*/*"],"Content-Length":["28"],"Content-Type":["application/x-www-form-urlencoded"],"User-Agent":["curl/7.64.1"],"X-Auth-Token":["top secret"]}
> [example] RELEASE_VERSION=d770d49e68
> [example] PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
> [example] CUSTOM_ENV_VAR1=hello
> [example] WEBHUG_REQUEST_REMOTE_ADDR=172.17.0.1:48240
> [example] CUSTOM_ENV_VAR2=world!
> [example] WEBHUG_REQUEST_METHOD=POST
> [example] PWD=/etc/webhug

```

## Docker

There is also a docker image available. If you prefer running webhug inside a container try this:
``` 
docker run --rm -p 8080:8080 gregthebunny/webhug
```

Within the image the default config is located at `/etc/webhug/config.yaml`. To provide your own config you
might either mount it into the container
``` 
docker run --rm -p 8080:8080 -v /path/to/your/config.yaml:/etc/webhug/config.yaml gregthebunny/webhug
```

or build your on image eg:

```Dockerfile
FROM gregthebunny/webhug

COPY /path/to/your/config.yaml /etc/webhug/config.yaml
```

As the image comes with a docker-in-docker setup (`FROM docker:19.03`) it enables you to
run docker commands within your actions. To do so you need to mount your hosts docker socket into the 
container:
``` 
docker run --rm -p 8080:8080 \
    -v /path/to/your/config.yaml:/etc/webhug/config.yaml \
    -v /var/run/docker.sock:/var/run/docker.sock \
    gregthebunny/webhug 
```
    
## Config

All configuration has to be done in the `config.yaml` file. A working example can be found 
in the project root:

```yaml
---
webhug:
  listen: ":8080"
  webhooks:
    # "example" is the name of the webhook as well as the the endpoint uri eg:
    # curl -H 'x-auth-token: top secret' -X POST -d '{"some": ["random", "json"]}' http://localhost:8080/example
    example:
      # does not yet have any meaning due to only "custom" is supported by now
      format: custom
      security:
        # one of "header" (match header value), "none" (no security at all), "deny" (deny all)
        type: header
        key: x-auth-token
        value: top secret
      action:
        # one of "shell" (execute shell command), "none" (do nothing)
        type: shell
        # if true stdout will be returned in the http response
        response: true
        # the executable to run
        cmd: "/bin/sh"
        # arguments that will be passed along to the executable
        args: ["-c", "cat; env; sleep 10;"]
        # some extra environment variables that will be injected.
        env:
          - "CUSTOM_ENV_VAR1=hello"
          - "CUSTOM_ENV_VAR2=world!"
``` 

### Security

The security section lets you configure the authentication for each webhook.

#### Deny

This will deny every request. Might be useful to temporary disable a certain webhook:
```yaml
webhug:
  webhooks:
    example:
      security:
        type: deny
``` 
#### None

*Please do not run any "open" webhooks in production .. nasty things can happen!*

The following example will leave your webhook open for everyone: 
```yaml
webhug:
  webhooks:
    example:
      security:
        type: none
``` 
#### Header

The most basic way of access control lets you define a shared secret to be present in the request header.
This is not recommended but better then nothing. Always make sure that you make your endpoint SSL only!

```yaml
webhug:
  webhooks:
    example:
      security:
        type: header
        key: x-any-header-you-like
        value: abc123topsecrettoken
``` 

Trigger your webhook like that:
``` 
curl -H 'x-any-header-you-like: abc123topsecrettoken' -X POST http://localhost:8080/example
```

#### Github

This will check for githubs `X-Hub-Signature` header and validates the digest using the shared secret:
```yaml
webhug:
  webhooks:
    example:
      security:
        type: github
        secret: 123456789abcdefghi
``` 

Further readings:
- https://developer.github.com/webhooks/securing/
- https://developer.github.com/webhooks/

### Actions

By now there is only one action supported with more to come.

#### Shell

This will execute the given command and passes the request body to stdin (see example). The following
environment variables will be available by default:

- `WEBHUG_WEBHOOK`: Name of the webhook which triggered the action
- `WEBHUG_REQUEST_METHOD`: The request method
- `WEBHUG_REQUEST_REMOTE_ADDR`: The remote address which triggered the action
- `WEBHUG_REQUEST_HEADER`: All request headers json encoded


## License
``` 
The MIT License (MIT)

Copyright (c) 2020 Maximilian Reichel <info@phramz.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

