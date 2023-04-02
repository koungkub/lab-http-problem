# lab-http-connection

## reference

- [connection reset by peer](https://gosamples.dev/connection-reset-by-peer/)
- [broken pipe](https://gosamples.dev/broken-pipe/)
- [http connection and pools](https://dev.to/mstryoda/golang-what-is-broken-pipe-error-tcp-http-connections-and-pools-4699)

## Problem

```text
(1) client --- (2) ---> (3) server 
```

**broken pipe**

when _(1) client_ write the data and sent it into _(3) server_ on a closed connection

**connection reset by peer**

when _(1) client_ read the data from _(3) server_ on a closed connection

> when _(3) server_ closed connection it means _(1) client_ received RST packet
