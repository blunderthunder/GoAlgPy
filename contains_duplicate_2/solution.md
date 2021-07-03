## Solution to Almost Duplicate

Using Hashtable

```
Time Complexity: O(n) ( since we are storing indices in hashtable, its takes o(1) time to search and get item )

Space Complexity: O(2n) ( since we are stoing index and value of array in hashtable, for each item in array it has allocate 1 space for index and 1 space for value itself )

```

## Summary

| Language | Runtime | Memory  |
| :--------| :------ | :------ |
| python3  | 104 ms  | 18.8 MB |
| go       | 10ms    | 6.2 MB  |


>Go is 30 times efficient than python for this example.