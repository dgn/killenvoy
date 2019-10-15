KillEnvoy
=========

This is a reproducer for [CVE-2019-15225](https://github.com/envoyproxy/envoy/issues/8519) (see the [underlying issue #7728](https://github.com/envoyproxy/envoy/issues/7728)).

Steps to reproduce:
- deploy an Istio control plane and a `httpbin` service (see `httpbin.yaml` in this repo)
- build the binary: `go build killenvoy.go`
- copy the binary to the `httpbin` container: `kubectl cp kill <NS>/<POD>:/tmp/killenvoy -c httpbin`
- run it: `kubectl exec -it -n <NS> <POD> -c httpbin -- /tmp/killenvoy`
- watch envoy segfault


