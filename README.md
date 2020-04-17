# PerfTests
Repository for comparing simple vector addition execution speeds between Python and GoLang 

#### Format
<pre>
"Benchmark Name"         "# of trials"             "average execution time per trial"
</pre>

## Vector addition including Array Creation time
FYI - Python tests are conducted with same # of trials as GoLang tests


<pre>
***Python Test***
3
0.021359586763122128 per operation
10000
0.0236658220528028 per operation
100000
0.08013223049298165 per operation

***GoLang Test***
BenchmarkVecAdd100000-2             1000          10184478 ns/op
(1)*BenchmarkSliceAdd-2*         1000             723641 ns/op
(2)*BenchmarkMatAdd10000-2*              1000            1995261 ns/op
(3)BenchmarkVecAdd10000-2              1000             877142 ns/op
BenchmarkVecAdd10-2                 1000             10907 ns/op
</pre>

Comparing the operations done using using slices, GoNum's Matrix, and GoNum's Vector: we find that vector and slice operations are similar
in time taken. These operations are faster than operations done using GoNum's Matrix

## Trials excluding Array Creation Time
<pre>
***Python Test***
Running Test for Array Size: 10
0.022797769828288232
Running Test for Array Size: 10000
0.02265420018466034
Running Test for Array Size: 100000
0.03405619764523144
Running Test for Array Size: 1000000
0.31470865151855903
Running Test for Array Size: 10000000
3.25162773480987
</pre>
<pre>
***GoLang Test***
BenchmarkSliceAddOnly10-2                   1000                 991 ns/op
BenchmarkSliceAddOnly10000-2                1000             662,123 ns/op
BenchmarkSliceAddOnly100000-2               1000           6,629,379 ns/op
</pre>
<pre>
***GoLang Test***
BenchmarkVecAddOnly10-2             1000               10699 ns/op
BenchmarkVecAddOnly10000-2          1000             549,371 ns/op
(1)*BenchmarkVecAddOnly100000-2         1000           7,905,048 ns/op*
BenchmarkVecAddOnly1000000-2         100         134,033,644 ns/op
BenchmarkVecAddOnly10000000-2         10       1,882,137,820 ns/op
</pre>
### Increase # of operations with Vector size of 100000
<pre>
***Python Test***
10000 Number of operations
Running Test for Array Size: 100000
3.7353716966821207

100000 Number of operations
Running Test for Array Size: 100000
30.713435394721955
</pre>
<pre>
***GoLang Test***
10000 Number of operations
BenchmarkVecAddOnly100000-2                  100         785,419,554 ns/op
100000 Number of operations
BenchmarkVecAddOnly100000-2                   10        8,682,452,270 ns/op
</pre>

## Conclusion so Far
Ignoring any base execution time required for the overhead in Python - GoLang is roughly 5x faster than Python DataFrame in terms of a simple vector addition operation.


## Appendix
Tests performed using dataframes
### Vector addition including Array Creation time
<pre>
***Python Test***
3
0.021359586763122128 per operation
10000
0.0236658220528028 per operation
100000
0.08013223049298165 per operation
</pre>

### Trials excluding Array Creation Time
<pre>
***Python Test***
Running Test for Array Size: 10
0.019733524791558706
Running Test for Array Size: 10000
0.021340448297199784
Running Test for Array Size: 100000
0.03508807013181842
Running Test for Array Size: 1000000
0.39808399950273016
Running Test for Array Size: 10000000
3.405962936630215
</pre>

10000 Number of operations
Running Test for Array Size: 100000
3.363972613837792

100000 Number of operations
Running Test for Array Size: 100000
37.66599165536402


## Links
Gonum vs Numpy eigenalue - https://github.com/gonum/gonum/issues/511

GoLang vs Python General Dynamic/Static Typing - https://stackoverflow.com/questions/12574909/can-go-really-be-that-much-faster-than-python

GoLang Goroutines - 
http://tleyden.github.io/blog/2014/10/30/goroutines-vs-threads/

Looking into Goroutines - 
https://rcoh.me/posts/why-you-can-have-a-million-go-routines-but-only-1000-java-threads/