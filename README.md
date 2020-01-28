Webhug 🤗
==========

A generic webhook dispatcher.

## Build & Run

``` 
git clone https://github.com/phramz/webhug.git
make build
./webhug

> INFO : 2020/01/25 18:25:03.520102 webhug.go:16: reading config ...
> INFO : 2020/01/25 18:25:03.521061 webhug.go:20: setting up webhook 'example' at path '/example/'
> INFO : 2020/01/25 18:25:03.521134 webhug.go:37: 🤗 webhug listening on :8080 ...
```

With the default config you should now be able to try this:
``` 
curl -H 'x-auth-token: top secret' -X POST -d '{"some": ["random", "json"]}' http://localhost:8080/example/

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
webhug:
  listen: ":8080"
  webhooks:
    # "example" is the name of the webhook as well as the the endpoint uri eg:
    # curl -H 'x-auth-token: top secret' -X POST -d '{"some": ["random", "json"]}' http://localhost:8080/example/
    example:
      security:
        # one of "header" (match header value), "github" (match x-hub-signature), "none" (no security at all), "deny" (deny all)
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
        args: ["-c", "cat; env;"]
        # some extra environment variables that will be injected.
        env:
          - "CUSTOM_ENV_VAR1=hello"
          - "CUSTOM_ENV_VAR2=world!"

    # ... of course you can configure more than one webhook!
    # This one here uses githubstyle authentication:
    # curl -H 'x-hub-signature: sha1=01f8da98455877d8beefe6e8276d59ac102857d5' -X POST -d '{"some": ["random", "json"]}' http://localhost:8080/example-github/
    example-github:
      security:
        type: github
        secret: super secret
      action:
        type: shell
        response: true
        cmd: "/bin/sh"
        args: ["-c", "cat; env;"]
        env:
          - "CUSTOM_ENV_VAR1=hello"
          - "CUSTOM_ENV_VAR2=world!"

    # this example shows how to pass along values from the request:
    # curl -v -X POST -d '{"some": [{"random": "value", "nested": {"value": "json"}}]}'  -H 'Content-type: application/json' http://localhost:8080/example-with-template-insecure/
    #
    # WARNING: If your planing to use values from request as arguments
    #          beware of command-injection attacks. Take security measures, eg proper escaping etc
    #          https://owasp.org/www-community/attacks/Command_Injection
    #          NEVER DO IT THIS WAY!! It is just an example to show what is possible.
    example-with-template-insecure:
      security:
        type: none
      action:
        type: shell
        response: true
        cmd: "/bin/sh"
        args:
          - '-c'
          - |
              echo 'CUSTOM_ENV_VAR_COMMAND=$CUSTOM_ENV_VAR_COMMAND';
              echo 'Request from: {{ .Request.RemoteAddr }}';
              echo 'Content-type: {{ index .Request.Header "content-type" }}';
              echo 'Body: {{ .Request.Body }}';
              echo 'Value: "{{ index .Request.Json.some 0 "random" }}" from "{{ index .Request.Json.some 0 "nested" "value" }}"';
        env:
          - 'CUSTOM_ENV_VAR_COMMAND={{ index .Env "PWD" }}'

    # ... this example shows two more secure ways of doing the same
    # curl -v -X POST -d '{"some": [{"random": "value", "nested": {"value": "json"}}]}'  -H 'Content-type: application/json' http://localhost:8080/example-with-template/
    example-with-template:
      security:
        type: none
      action:
        type: shell
        response: true
        cmd: "/bin/sh"
        args:
          - '-c'
          - |
              # either using environment variables like this
              echo CUSTOM_ENV_VAR_COMMAND=${CUSTOM_ENV_VAR_COMMAND};
              echo Request from: ${CUSTOM_ENV_VAR_REQUEST_FROM};
              echo Content-type: ${CUSTOM_ENV_VAR_REQUEST_CTYPE};
              echo Body: ${CUSTOM_ENV_VAR_REQUEST_BODY};
              echo Value: ${CUSTOM_ENV_VAR_REQUEST_VALUE};

              # or take the arguments directly
              echo CUSTOM_ENV_VAR_COMMAND=${0};
              echo Request from: ${1};
              echo Content-type: ${2};
              echo Body: ${3};
              echo Value: ${4};

          - '{{ index .Env "PWD" }}'
          - '{{ .Request.RemoteAddr }};'
          - '{{ index .Request.Header "content-type" }}'
          - '{{ .Request.Body }}'
          - '"{{ index .Request.Json.some 0 "random" }}" from "{{ index .Request.Json.some 0 "nested" "value" }}"'
        env:
          - 'CUSTOM_ENV_VAR_COMMAND={{ index .Env "PWD" }}'
          - 'CUSTOM_ENV_VAR_REQUEST_FROM={{ .Request.RemoteAddr }}'
          - 'CUSTOM_ENV_VAR_REQUEST_CTYPE={{ index .Request.Header "content-type" }}'
          - 'CUSTOM_ENV_VAR_REQUEST_BODY={{ .Request.Body }}'
          - 'CUSTOM_ENV_VAR_REQUEST_VALUE="{{ index .Request.Json.some 0 "random" }}" from "{{ index .Request.Json.some 0 "nested" "value" }}"'
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
curl -H 'x-any-header-you-like: abc123topsecrettoken' -X POST http://localhost:8080/example/
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

Furthermore the raw request body will be piped to stdin of the shell command.

### Templating

Some config keys will be interpolated to add some further flexibility to your configuration.
The string interpolation will be done via Golang templating. If your not familiar with the syntax this
cheat sheet might come in handy:
https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet

The actual keys which values could be templated the following:
- `webhug.webhooks.<*>.security[type=header].key`
- `webhug.webhooks.<*>.security[type=header].value`
- `webhug.webhooks.<*>.security[type=github].secret`
- `webhug.webhooks.<*>.action[type=shell].args.<*>`
- `webhug.webhooks.<*>.action[type=shell].env.<*>`

During rendering you have access to this data:
```go
type Context struct {
    // all environment variable of the context webhug is running in are here
    Env            map[string]string
    Webhook struct {
        Name       string
        Format     string
    }
    Request struct {
        Body       string
        Method     string
        // if a POST request comes with a json content-type header
        // the body will be automatically deserialized and could be access via
        // this property. 
        Json       interface{}
        Uri        string
        Host       string
        RemoteAddr string
        Query      string
        Scheme     string
        Username   string
        Password   string
        // all headers lowercase CanonicalMIMEHeaderKeys 
        Header     map[string]string
        // all variables from the requests query string
        Get        map[string][]string
        // see https://golang.org/pkg/net/http/#Cookie
        Cookie     map[string]*http.Cookie
    }
}
```

## Github Actions

If you want to trigger Webhugs out of your Github Actions you might find this useful:
https://github.com/phramz/webhug-action

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

