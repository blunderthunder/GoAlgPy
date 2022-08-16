# solution of compare version

## Algorithm

*  Split version1 and version2 by "." and make it int(remove leading zeros) => O(M+N)
* Iterate through v1 and v2, check which reversion is greater. Once we exceed a version's length while the other is not we check if the remaining version have not 0 value => O(max(M,N)

complexity : O( M + N )

## Summary

| Language | Runtime | Memory  |
| :--------| :------ | :------ |
| python3  | 53 ms | 13.8 MB |
| go       | 0.5 ms  | 2 MB  |
| rust     | 2 ms   | 2.2 MB  |

> Overall go is `4` times better than rust and `~1000` times better than python.
