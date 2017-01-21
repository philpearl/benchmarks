
#A bunch of benchmarks for basic Go things.

##Current results
On a 2015 MacBook Pro, go version go1.7.4 darwin/amd64

```
BenchmarkChannelOneByte-8        	20000000	       100 ns/op	  10.00 MB/s	       0 B/op	       0 allocs/op
BenchmarkChannelOneByteMulti-8   	 5000000	       278 ns/op	   3.59 MB/s	       0 B/op	       0 allocs/op
BenchmarkCopy-8                  	2000000000	         0.83 ns/op	1212.02 MB/s	       0 B/op	       0 allocs/op
BenchmarkSync-8                  	100000000	        23.3 ns/op
BenchmarkAlloc-8                 	50000000	        30.1 ns/op	       8 B/op	       1 allocs/op
BenchmarkAddToMap-8              	 3000000	       400 ns/op	      84 B/op	       0 allocs/op
BenchmarkCheckAddToMap-8         	 3000000	       466 ns/op	      84 B/op	       0 allocs/op
BenchmarkCheckAddToBigMap-8      	 3000000	       615 ns/op	     121 B/op	       0 allocs/op
BenchmarkSizedCheckAddToMap-8    	10000000	       264 ns/op	       2 B/op	       0 allocs/op
BenchmarkSharedBuffer-8          	 5000000	       256 ns/op	   3.90 MB/s	       0 B/op	       0 allocs/op
BenchmarkSharedBufferMulti-8     	 5000000	       256 ns/op	   3.89 MB/s	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_1-8   	100000000	        13.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_10-8  	 5000000	       267 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_25-8  	 1000000	      1003 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_100-8 	  100000	     20366 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_500-8 	    3000	    487358 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedList/add_1000-8         	    1000	   2125965 ns/op	       0 B/op	       0 allocs/op
```