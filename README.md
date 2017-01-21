
#A bunch of benchmarks for basic Go things.

##Current results
On a 2015 MacBook Pro, go version go1.7.4 darwin/amd64

```
BenchmarkChannelOneByte-8        	20000000	        98.5 ns/op	  10.16 MB/s	       0 B/op	       0 allocs/op
BenchmarkChannelOneByteMulti-8   	10000000	       230 ns/op	   4.34 MB/s	       0 B/op	       0 allocs/op
BenchmarkCopy-8                  	2000000000	         0.27 ns/op	3756.12 MB/s	       0 B/op	       0 allocs/op
BenchmarkSync-8                  	50000000	        25.3 ns/op
BenchmarkAlloc-8                 	50000000	        32.9 ns/op	       8 B/op	       1 allocs/op
BenchmarkAddToMap-8              	 3000000	       427 ns/op	      84 B/op	       0 allocs/op
BenchmarkCheckAddToMap-8         	 2000000	       635 ns/op	     121 B/op	       0 allocs/op
BenchmarkCheckAddToBigMap-8      	 3000000	       625 ns/op	     121 B/op	       0 allocs/op
BenchmarkSizedCheckAddToMap-8    	10000000	       255 ns/op	       2 B/op	       0 allocs/op
BenchmarkSharedBuffer-8          	20000000	        94.8 ns/op	  10.55 MB/s	       0 B/op	       0 allocs/op
BenchmarkSharedBufferMulti-8     	20000000	        94.5 ns/op	  10.58 MB/s	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_1-8   	100000000	        12.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_10-8  	50000000	        23.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_25-8  	30000000	        43.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_100-8 	10000000	       222 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_500-8 	 1000000	      1178 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_1000-8         	 1000000	      2302 ns/op	       0 B/op	       0 allocs/op
```