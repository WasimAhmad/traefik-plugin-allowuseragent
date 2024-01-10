# Block User-Agent

Allow User-Agent is a middleware plugin for [Traefik](https://github.com/traefik/traefik) which sends an HTTP `403 Forbidden` 
response when the requested HTTP User-Agent header does not matches one the configured [regular expressions](https://github.com/google/re2/wiki/Syntax).

## Configuration
## Static

```toml
[experimental.plugins.allowuseragent]
    modulename = "github.com/EmergingPixels/traefik-plugin-allowuseragent"
    version = "v0.0.1"
```

## Dynamic

To configure the `Block User-Agent` plugin you should create a [middleware](https://docs.traefik.io/middlewares/overview/) in 
your dynamic configuration as explained [here](https://docs.traefik.io/middlewares/overview/). The following example creates
and uses the `blockuseragent` middleware plugin to block all HTTP requests with a User-Agent like `\bTheAgent\b`. 

```toml
[http.routers]
  [http.routers.my-router]
    rule = "Host(`localhost`)"
    middlewares = ["block-foo"]
    service = "my-service"

# Block all paths starting with /foo
[http.middlewares]
  [http.middlewares.allow-foo.plugin.allowuseragent]
    regex = ["\bTheAgent\b"]

[http.services]
  [http.services.my-service]
    [http.services.my-service.loadBalancer]
      [[http.services.my-service.loadBalancer.servers]]
        url = "http://127.0.0.1"
```
