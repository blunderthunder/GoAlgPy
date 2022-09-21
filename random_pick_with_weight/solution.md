# Solution of random pick with weight 

## Approch

- create a list of to store cumulative sum.  O( N )
- generate random number between 0 and 1 sum.  O( 1 )
- binary search of randomly generated number inside cumulative sum list and get index. O( LogN )

## Complexity

O ( N + LogN )

## Summary

| Language | Runtime | Memory  |
| :--------| :------ | :------ |
| python3  | 405 ms | 18.4 MB |
| go       | 84 ms  | 8.4 MB  |
| rust     | 25 ms   | 4.2 MB  |

> Overall rust is 80 times faster than python3 and golang is 40 times faster than python3.