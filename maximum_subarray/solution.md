## Solution to Max subarray problem

this problem was discussed by Jon Bentley (Sep. 1984 Vol. 27 No. 9 Communications of the ACM P885)

the paragraph below was copied from his paper (with a little modifications)

algorithm that operates on arrays: it starts at the left end (element A[1]) and scans through to the right end (element A[n]), keeping track of the maximum sum subvector seen so far. The maximum is initially A[0]. Suppose we've solved the problem for A[1 .. i - 1]; how can we extend that to A[1 .. i]? The maximum
sum in the first I elements is either the maximum sum in the first i - 1 elements (which we'll call MaxSoFar), or it is that of a subvector that ends in position i (which we'll call MaxEndingHere).

MaxEndingHere is either A[i] plus the previous MaxEndingHere, or just A[i], whichever is larger.

```
Time Complexity: O(n) ( for n loops with array n having n number of elements)

Space Complexity: O(1) ( Constant: i.e we didn't use any hashtable or array to store the vaue, int storage was enough )
```

## Summary 

| Language | Runtime | Memory  |
| :--------| :------ | :------ |
| python3  | 52 ms   | 15.1 MB |
| go       | 0.6 ms  | 3.4 MB  |

*With exact same code , Considering both runtime and Memory , i.e.  52/0.1 and  14.4/3.2 -> 86.6 x 4.7 -> 404*

>Go is 404 times efficient than python for this example.