# DOME - Call Metrics

This metric package is a simple wrapper around VictoriaMetrics client.

At this time the client is attached to each service and can be called easily by handlers.

Example usage:

```go
s.Metric().Call("method", "api", apiKey)
```
