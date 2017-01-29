
#A bunch of benchmarks for basic Go things.

## Headlines

- Locking and unlocking an uncontended mutex takes ~25 ns
- Allocating some memory takes ~30 ns
- deferring a function takes ~90 ns
- Pushing a byte through a channel takes ~100 to 250 ns
- Checking if something is in a map, then adding it takes ~250 ns if the map has enough room
- Checking if something is in an unordered slice, then appending it takes ~222 ns up to around 100 items if the slice has cpacity
- A type assertion takes ~0.90 ns
- Storing a value in an interface type takes ~33 ns and involves an allocation, even for relatively small types
- Storing an existing pointer in an interface type takes ~8ns and no allocations.
- It's slightly faster to do a type assertion on an interface type then a call on the concrete type (1.1 ns), than to call a method on the interface type directly (2.3 ns)

##Current results
On a 2015 MacBook Pro, go version go1.7.4 darwin/amd64

```
BenchmarkChannelOneByte-8                          	20000000	        99.3 ns/op	  10.07 MB/s	       0 B/op	       0 allocs/op
BenchmarkChannelOneByteUnbuffered-8                	 5000000	       260 ns/op	   3.84 MB/s	       0 B/op	       0 allocs/op
BenchmarkChannelOneByteMultiReceive-8              	10000000	       210 ns/op	   4.74 MB/s	       0 B/op	       0 allocs/op
BenchmarkChannelOneByteMultiSendRecv-8             	10000000	       200 ns/op	   5.00 MB/s	       0 B/op	       0 allocs/op
BenchmarkChannelOneByteMultiSendRecvUnbuffered-8   	 2000000	       606 ns/op	   1.65 MB/s	       0 B/op	       0 allocs/op
BenchmarkCopy-8                                    	2000000000	         0.16 ns/op	6360.59 MB/s	       0 B/op	       0 allocs/op
BenchmarkDefer-8                                   	20000000	        97.2 ns/op
BenchmarkInterfaceTypeAssertion-8                  	2000000000	         0.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterfaceCall-8                           	1000000000	         2.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterfaceCallComparison-8                 	2000000000	         0.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterfaceCallTypeAssertion-8              	2000000000	         1.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterfaceStore-8                          	50000000	        26.0 ns/op	       8 B/op	       1 allocs/op
BenchmarkInterfaceStorePointer-8                   	2000000000	         0.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkSync-8                                    	100000000	        23.4 ns/op
BenchmarkAlloc-8                                   	50000000	        28.2 ns/op	       8 B/op	       1 allocs/op
BenchmarkAddToMap-8                                	 3000000	       398 ns/op	      84 B/op	       0 allocs/op
BenchmarkCheckAddToMap-8                           	 3000000	       454 ns/op	      84 B/op	       0 allocs/op
BenchmarkCheckAddToBigMap-8                        	 3000000	       618 ns/op	     121 B/op	       0 allocs/op
BenchmarkSizedCheckAddToMap-8                      	10000000	       249 ns/op	       2 B/op	       0 allocs/op
BenchmarkSharedBuffer-8                            	20000000	        95.5 ns/op	  10.48 MB/s	       0 B/op	       0 allocs/op
BenchmarkSharedBufferMulti-8                       	20000000	       290 ns/op	   3.44 MB/s	       0 B/op	       0 allocs/op
BenchmarkSharedBufferMultiSendRecv-8               	10000000	       266 ns/op	   3.75 MB/s	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_1-8                     	100000000	        12.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_10-8                    	100000000	        25.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_25-8                    	30000000	        38.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_100-8                   	10000000	       180 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_500-8                   	 1000000	      1063 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_1000-8                  	 1000000	      2181 ns/op	       0 B/op	       0 allocs/op
```