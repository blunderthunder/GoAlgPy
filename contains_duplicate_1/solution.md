
# Solution to  Contains Duplicate

Using Hashtable to solve this problem.

```
Time Comlexity: O(n) ( We do search() and insert() for nn times and each operation takes constant time. )

Space Complexity: O(n) ( The space used by a hash table is linear with the number of elements in it. )
```

## Summary

| Language | Runtime | Memory  |
| :--------| :------ | :------ |
| python3  | 116 ms  | 21.4 MB |
| go       | 16 ms   | 7.9 MB  |

*With exact same code , Considering both runtime and Memory , i.e.  116/16 and  21.4/7.9 -> 7.25 x 3.53 -> 25.6*

>Go is 25.6 times efficient than python for this example.