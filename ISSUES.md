## Why not keep returning Copy
When we are using pure type ( string, int, etc ), do copy at every function doesn't effect performance too much.
```
BenchmarkString_NativeReplaceAll-4      10000000               139 ns/op
BenchmarkString_ReplaceAll-4            10000000               141 ns/op
BenchmarkString_ReplaceAllWithCopy-4    10000000               142 ns/op
```
But if it is slice, copy of slice everytime may suddenly increase memory usage and trigger GC then slow down performence.
```
BenchmarkStringSlice_NativeAppend-4      5000000               308 ns/op
BenchmarkStringSlice_Push-4              5000000               301 ns/op
BenchmarkStringSlice_PushWithCopy-4        10000            192625 ns/op
```
For consistency function behavior, I decide to prevent Copy at every return.