# Solution of all possible full binary trees 

## Approch

using recursive technique to solve the problem. Full binary tree is tree that has either both  zero children node or two children node. 
- first pick a point on the index which you want to be top most parent, then break them as left tree and right tree
- since both left and right need to be full binary tree its a recursive functions until it find either node with 0 children or 2 children.

## Complexity
O(N^2)

## Summary

| Language | Runtime | Memory  |
| :--------| :------ | :------ |
| python3  | 470 ms | 17.4 MB |
| go       | 15 ms  | 9.4 MB  |
| rust     | 25 ms   | 3.2 MB  |

> Overall rust and golang is 40 times faster than python3.