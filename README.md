# PerfTests
Repository for comparing execution speeds

### Array of 3
#### Python
Time: 0.021359586763122128 per operation

#### GoLang Slices
pkg: perf/main/sliceadd
BenchmarkSliceAdd-2         1000              1001 ns/op

#### GoLang GoNum
BenchmarkVedAcc-2           1000             11723 ns/op

#### Result
GoLang using GoNum is roughly half the time


100 Operations 

### Array of 10000
#### Python
Time: 0.028487594093662742 per operation

#### GoLang Slices
pkg: perf/main/sliceadd
BenchmarkSliceAdd-2         1000             723641 ns/op

#### GoLang GoNum
BenchmarkVedAcc-2           1000             667037 ns/op

#### Result
GoLang using GoNum is roughly half the time


## With Array Creation
10000
0.0236658220528028
100000
0.08013223049298165

BenchmarkVecAdd100000-2             1000          10184478 ns/op
BenchmarkVecAdd10000-2              1000            877142 ns/op
BenchmarkVecAdd10-2                 1000             10907 ns/op
BenchmarkMatAdd10000-2              1000           1995261 ns/op


## Without Array Creation
Running Test for Size: 10
0.019733524791558706
Running Test for Size: 10000
0.021340448297199784
Running Test for Size: 100000
0.03508807013181842
Running Test for Size: 1000000
0.39808399950273016
Running Test for Size: 10000000
3.405962936630215


BenchmarkSliceAddOnly10-2                   1000                 991 ns/op
BenchmarkSliceAddOnly10000-2                1000             662,123 ns/op
BenchmarkSliceAddOnly100000-2               1000           6,629,379 ns/op

BenchmarkVecAddOnly10-2             1000               10699 ns/op
BenchmarkVecAddOnly10000-2          1000             549,371 ns/op
*BenchmarkVecAddOnly100000-2         1000           7,905,048 ns/op*
BenchmarkVecAddOnly1000000-2         100         134,033,644 ns/op
BenchmarkVecAddOnly10000000-2         10       1,882,137,820 ns/op

### 10000 Number of operations

Running Test for Size: 100000
3.363972613837792

BenchmarkVecAddOnly100000-2                  100         785,419,554 ns/op

### 100000 Number of operations

Running Test for Size: 100000
37.66599165536402

BenchmarkVecAddOnly100000-2                   10        8,682,452,270 ns/op