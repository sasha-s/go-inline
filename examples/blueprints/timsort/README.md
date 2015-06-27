The code is adapded from [timsort](https://github.com/psilva261/timsort).

Simply elminating the comparator can speed up the code substantially.
Running this code:
<pre>
BenchmarkTimsortIXor100	  500000	      2380 ns/op
BenchmarkStandardSortIXor100	  300000	      4763 ns/op
BenchmarkTimsortISorted100	 2000000	       603 ns/op
BenchmarkStandardSortISorted100	  500000	      3338 ns/op
BenchmarkTimsortIRevSorted100	 2000000	       705 ns/op
BenchmarkStandardSortIRevSorted100	  300000	      3961 ns/op
BenchmarkTimsortIRandom100	  500000	      2574 ns/op
BenchmarkStandardSortIRandom100	  300000	      5339 ns/op
BenchmarkTimsortIXor1K	   50000	     22023 ns/op
BenchmarkStandardSortIXor1K	   20000	     89818 ns/op
BenchmarkTimsortISorted1K	 1000000	      1734 ns/op
BenchmarkStandardSortISorted1K	   30000	     46269 ns/op
BenchmarkTimsortIRevSorted1K	 1000000	      2283 ns/op
BenchmarkStandardSortIRevSorted1K	   30000	     51069 ns/op
BenchmarkTimsortIRandom1K	   20000	     62109 ns/op
BenchmarkStandardSortIRandom1K	   10000	    119133 ns/op
BenchmarkTimsortIXor1M	      50	  32933014 ns/op
BenchmarkStandardSortIXor1M	      20	  79442638 ns/op
BenchmarkTimsortISorted1M	    2000	    909548 ns/op
BenchmarkStandardSortISorted1M	      20	 102517670 ns/op
BenchmarkTimsortIRevSorted1M	    1000	   1426532 ns/op
BenchmarkStandardSortIRevSorted1M	      20	  95939774 ns/op
BenchmarkTimsortIRandom1M	      10	 121904910 ns/op
BenchmarkStandardSortIRandom1M	       5	 245630006 ns/op
</pre>

Running benchmarks for BenchmarkTimsortI in [timsort](https://github.com/psilva261/timsort) gives:
<pre>
BenchmarkTimsortIXor100	  500000	      3884 ns/op
BenchmarkTimsortISorted100	 2000000	       854 ns/op
BenchmarkTimsortIRevSorted100	 1000000	      1021 ns/op
BenchmarkTimsortIRandom100	  300000	      4355 ns/op
BenchmarkTimsortIXor1K	   50000	     36241 ns/op
BenchmarkTimsortISorted1K	  500000	      3571 ns/op
BenchmarkTimsortIRevSorted1K	  300000	      5106 ns/op
BenchmarkTimsortIRandom1K	   20000	     94382 ns/op
BenchmarkTimsortIXor1M	      30	  52208575 ns/op
BenchmarkTimsortISorted1M	    1000	   2342377 ns/op
BenchmarkTimsortIRevSorted1M	     500	   3986431 ns/op
BenchmarkTimsortIRandom1M	      10	 185847098 ns/op
</pre>
on the same machine.

