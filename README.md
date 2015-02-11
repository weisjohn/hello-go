# hello-go

Just a place to store files as I learn 

Most of these files are notes that I made while progressing through the (Feb 2015 version) of the http://tour.golang.org. Many of the files will not run locally without the google code packages.

File that are named "exercise-" are solutions to those exercises.

### concurrency != parallelization

I have one example which is quite helpful to show the benefits of concurrency. Writing programs that are concurrent can (but do not necessarily) benefit from concurrency. This is a stupid example, but it's intended to show that if you run a concurrent Go program in a parallel way, you can achieve speed improvements.


To demonstrate this, I've written a small package in this folder called [counter](./counter) which initializes an array with values 1...n and then returns a `func() int` which, when invoked, will loop through the array and return the last value in it.

I have two implementations: [counter-loop.go](./counter-loop.go) is a simple for-loop, and [counter-goroutine.go](./counter-goroutine.go) another that uses goroutines. Each program prints the duration and the calculated value.

```bash
$ go run counter-long.go 
duration 630.086676ms
sum 100000000
```

```bash
$ go run counter-goroutine.go 
duration 624.47364ms
sum 100000000
```

The `counter-goroutine.go` is actually on average a little slower. Why is this happening? Because Go is only using one OS-level thread.

To use multiple OS-level threads, you must set the `$GOMAXPROCS` environemnt variable, probably to the number of logical cores you have on your CPU. For example, In my `~/.bash_rc` on my Mac, I run: 

```
export GOMAXPROCS=$(sysctl hw.ncpu | cut -d ' ' -f 2)  // returns 8
```

Doing this, we can see dramatic improvements:

```bash
$ GOMAXPROCS=2 go run counter-goroutine.go 
duration 360.352314ms
sum 100000000
```

```bash
$ GOMAXPROCS=4 go run counter-goroutine.go 
duration 244.658832ms
sum 100000000
````

```
$ GOMAXPROCS=8 go run counter-goroutine.go 
duration 206.06136ms
sum 100000000
```

And at this point, we have hit the logical limit, and further speed improvement is (conceptually) impossible, as there are no more CPU cores on which to do the work.

(This is not intended to be a scientific example, but rather a high-level understanding that concurrency can lead to parallelization speed improvements, but will not necessarily do so).

