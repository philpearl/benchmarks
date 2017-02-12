
#A bunch of benchmarks for basic Go things.

## Headlines

- Locking and unlocking an uncontended mutex takes ~25 ns
- Allocating some memory takes ~30 ns
- Using a sync.Pool takes ~25ns
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
BenchmarkChannelOneByte-8                          	20000000	        96.7 ns/op	  10.34 MB/s	       0 B/op	       0 allocs/op
BenchmarkChannelOneByteUnbuffered-8                	 5000000	       282 ns/op	   3.54 MB/s	       0 B/op	       0 allocs/op
BenchmarkChannelOneByteMultiReceive-8              	 5000000	       285 ns/op	   3.50 MB/s	       0 B/op	       0 allocs/op
BenchmarkChannelOneByteMultiSendRecv-8             	10000000	       194 ns/op	   5.14 MB/s	       0 B/op	       0 allocs/op
BenchmarkChannelOneByteMultiSendRecvUnbuffered-8   	 3000000	       570 ns/op	   1.75 MB/s	       0 B/op	       0 allocs/op
BenchmarkCopy-8                                    	2000000000	         0.27 ns/op	3723.21 MB/s	       0 B/op	       0 allocs/op
BenchmarkDefer-8                                   	20000000	        98.9 ns/op
BenchmarkInterfaceTypeAssert-8                     	2000000000	         0.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterfaceTypeAssertInterface-8            	100000000	        13.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterfaceTypeSwitch-8                     	1000000000	         2.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterfaceCall-8                           	1000000000	         2.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterfaceCallComparison-8                 	2000000000	         0.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterfaceCallTypeAssertion-8              	1000000000	         2.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterfaceStore-8                          	50000000	        26.0 ns/op	       8 B/op	       1 allocs/op
BenchmarkInterfaceStorePointer-8                   	2000000000	         0.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkLevenshtein_philpearl-8                   	 3000000	       400 ns/op	       0 B/op	       0 allocs/op
BenchmarkLevenshtein_dgryski-8                     	 3000000	       432 ns/op	      80 B/op	       1 allocs/op
BenchmarkLevenshtein_texttheater-8                 	 1000000	      1277 ns/op	    1200 B/op	      11 allocs/op
BenchmarkLevenshtein_kse-8                         	 5000000	       388 ns/op	       0 B/op	       0 allocs/op
BenchmarkLevenshtein_honzab-8                      	  300000	      3810 ns/op	     192 B/op	       2 allocs/op
BenchmarkLevenshtein_arbovm-8                      	 3000000	       458 ns/op	      80 B/op	       1 allocs/op
BenchmarkLevenshtein_agnivade-8                    	 2000000	       686 ns/op	     192 B/op	       2 allocs/op
BenchmarkSync-8                                    	100000000	        25.0 ns/op
BenchmarkAlloc-8                                   	50000000	        30.5 ns/op	       8 B/op	       1 allocs/op
BenchmarkAddToMap-8                                	 3000000	       434 ns/op	      84 B/op	       0 allocs/op
BenchmarkCheckAddToMap-8                           	 3000000	       492 ns/op	      84 B/op	       0 allocs/op
BenchmarkCheckAddToBigMap-8                        	 3000000	       643 ns/op	     121 B/op	       0 allocs/op
BenchmarkSizedCheckAddToMap-8                      	10000000	       255 ns/op	       2 B/op	       0 allocs/op
BenchmarkSharedBuffer-8                            	20000000	        92.9 ns/op	  10.76 MB/s	       0 B/op	       0 allocs/op
BenchmarkSharedBufferMulti-8                       	 5000000	       326 ns/op	   3.06 MB/s	       0 B/op	       0 allocs/op
BenchmarkSharedBufferMultiSendRecv-8               	 5000000	       263 ns/op	   3.80 MB/s	       0 B/op	       0 allocs/op
BenchmarkSyncPool-8                                	50000000	        25.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_1-8                     	100000000	        13.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_10-8                    	100000000	        30.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_25-8                    	30000000	        45.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_100-8                   	10000000	       226 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_500-8                   	 1000000	      1162 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_1000-8                  	 1000000	      2438 ns/op	       0 B/op	       0 allocs/op
```