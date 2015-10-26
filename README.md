# Structured logger

Bare minimum structured logger.

* Logs in JSON.
* 3 log levels: info, error, fatal.
* Reusable context logging

```go
ll := log.KV("who", "world")
ll.Info("hello?")

if err := doThing(); err != nil {
    ll.Err(err).Error("this thing failed")
}
ll.KV("why", "no reason").Fatal("abort abort abort!")
```
