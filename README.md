
#A bunch of benchmarks for basic Go things.

## Headlines

- Locking and unlocking an uncontended mutex takes ~25 ns
- Allocating some memory takes ~30 ns
- deferring a function takes ~90 ns
- Pushing a byte through a channel takes ~100 to 250 ns
- Checking if something is in a map, then adding it takes ~250 ns if the map has enough room
- Checking if something is in an unordered slice, then appending it takes ~222 ns up to around 100 items if the slice has cpacity

##Current results
On a 2015 MacBook Pro, go version go1.7.4 darwin/amd64

```
BenchmarkChannelOneByte-8                	20000000	        99.2 ns/op	  10.08 MB/s	       0 B/op	       0 allocs/op
BenchmarkChannelOneByteMultiReceive-8    	10000000	       222 ns/op	   4.49 MB/s	       0 B/op	       0 allocs/op
BenchmarkChannelOneByteMultiSendRecv-8   	10000000	       195 ns/op	   5.11 MB/s	       0 B/op	       0 allocs/op
BenchmarkCopy-8                          	2000000000	         0.20 ns/op	4936.86 MB/s	       0 B/op	       0 allocs/op
BenchmarkDefer-8                         	20000000	        91.4 ns/op
BenchmarkSync-8                          	100000000	        23.0 ns/op
BenchmarkAlloc-8                         	50000000	        28.6 ns/op	       8 B/op	       1 allocs/op
BenchmarkAddToMap-8                      	 3000000	       454 ns/op	      84 B/op	       0 allocs/op
BenchmarkCheckAddToMap-8                 	 3000000	       470 ns/op	      84 B/op	       0 allocs/op
BenchmarkCheckAddToBigMap-8              	 3000000	       645 ns/op	     121 B/op	       0 allocs/op
BenchmarkSizedCheckAddToMap-8            	 5000000	       256 ns/op	       2 B/op	       0 allocs/op
BenchmarkSharedBuffer-8                  	20000000	        93.2 ns/op	  10.73 MB/s	       0 B/op	       0 allocs/op
BenchmarkSharedBufferMulti-8             	 5000000	       334 ns/op	   2.99 MB/s	       0 B/op	       0 allocs/op
BenchmarkSharedBufferMultiSendRecv-8     	 5000000	       264 ns/op	   3.77 MB/s	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_1-8           	100000000	        13.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_10-8          	50000000	        22.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_25-8          	30000000	        68.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_100-8         	10000000	       234 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_500-8         	 1000000	      1164 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_1000-8        	 1000000	      2267 ns/op	       0 B/op	       0 allocs/op
```