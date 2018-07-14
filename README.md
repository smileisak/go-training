# GoLang Training [![Go Report Card](https://goreportcard.com/badge/github.com/smileisak/go-training)](https://goreportcard.com/report/github.com/smileisak/go-training)

This repository contains my training path to master GoLang in one month.

Resources can be found [here](http://harrymoreno.com/2016/06/30/How-to-learn-Golang-in-1-month.html).


# Step 1: Go By Examples
Trying out [Go by Example](https://gobyexample.com/) by [@mmcgrana](https://twitter.com/mmcgrana), i've created one go file for each topic.

1. [Variables](./variables.go)
1. [Constants](./constants.go)
1. [For](./for.go)
1. [If/Else](./ifelse.go)
1. [Switch](./switch.go)
1. [Arrays](./arrays.go)
1. [Slices](./slices.go)
1. [Maps](./maps.go)
1. [Range](./range.go)
1. [Functions](./functions.go)
1. [Multiple Return Values](./functions.go)
1. [Variadic Functions](./functions.go)
1. [Closures](./closures.go)
1. [Recursion](./recursive.go)
1. [Pointers](./pointers.go)
1. [Structs](./structs.go)
1. [Methods](./methods.go)
1. [Interfaces](./interfaces.go)
1. [Errors](./errors.go)
1. [Goroutines](./go-routines.go)
1. [Channels](./channels.go)
1. [Channel Buffering](./channel-buffering.go)
1. [Channel Synchronization](./channel-sync.go)
1. [Channel Directions](./channel-directions.go)
1. [Select](select.go)
1. [Non-Blocking Channel Operations](./channels-non-blocking.go)
1. [Closing Channels](./close-channels.go)
1. [Range over Channels](./range-channels.go)
1. [Timers](./timers.go)
1. [Tickers](./tickers.go)
1. [Worker Pools](./worker-pools.go)
1. [Rate Limiting](./rate-limiting.go)
1. [Atomic Counters](./atomic-counter.go)
1. [Mutexes](./mutexes.go)
1. [Stateful Goroutines](./stateful-goroutines.go)
1. [Sorting](./sorting.go)
1. [Sorting by Functions](./sorting-by-func.go)
1. [Panic](./panic.go)
1. [Defer](./defer.go)
1. [Collection Functions](./collection-functions.go)
1. [String Functions](./string-functions.go)
1. [String Formatting](./string-formatting.go)
1. [Regular Expressions](./regex.go)
1. [JSON](./json.go)
1. [Time](./time.go)
1. [Epoch](./epoch.go)
1. [Time Formatting / Parsing](./time-formatting.go)
1. [Random Numbers](./random.go)
1. [Number Parsing](./number-parsing.go)
1. [URL Parsing](./url-parsing.go)
1. [SHA1 Hashes](./sha1-hashes.go)
1. [Base64 Encoding](./base64.go)
1. [Reading Files](./reading-files.go)
1. [Writing Files](./writing-files.go)
1. [Line Filters](./filters/line-filter.go)
1. [Command-Line Arguments](./cmd/args/args.go)
1. [Command-Line Flags](./cmd/flags/flags.go)
1. [Environment Variables](./envvars.go)
1. [Spawning Processes](./cmd/process/spawn/main.go)
1. [Exec'ing Processes](./cmd/process/exec/exec.go)
1. [Signals](./signals.go)
1. [Exit](./exit.go)

# Discovering Go packages:

This section is for discovering Go official or unofficial packages. It is like a sandbox for more advanced examples and tools that can be useful when doing awsome things.

1. [ssh](./ssh/cmd/main.go)

    This example uses `golang.org/x/crypto/ssh` to run an arbitary commands in a remote server. May be useful when creating programs that needs to make sshing stuff in a remote server.


1. [ssh-tunneling](./ssh/tunneling/main.go)

    This example will establish an ssh tunnel listening from localhost passing through a bastion to a remote server.

1. [Kubernetes go client](https://github.com/kubernetes/client-go/)
    1. [Out cluster Example](./k8s/examples/out-cluster/main.go): Discovering k8s go client to list pods and nodes from outside the cluster.
