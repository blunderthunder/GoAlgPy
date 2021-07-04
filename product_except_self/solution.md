## Solution of product of array except self

To solve this problem in `O(n)` complexity, you need to create two pointers of loop through array.

Left pointer will store the cumulative product from left of the array to the right end and right pointer will store cumulative product from right side of the array to the left end. 

For example:
```
arr = [1,2,3,4]

initilize 
    leftpointer = 1
    rightpointer = 1
    result = [1,1,1,1]

loop 1:
    [leftpointer, 1 , 1 , rightpointer]
    leftpointer = 1 * 1
    rightpointer = 1 * 4
loop 2:
    [leftpointer, leftpointer  , rightpointer, rightpointer]
    leftpointer = 1 * 1 * 2
    rightpointer = 1 * 4 * 3
loop 3:
    [leftpointer, leftpointer * rightpointer, rightpointer * leftpointer , rightpointer]
    leftpointer = 1 * 1 * 2 * 3
    rightpointer = 1 * 4 * 3 * 2
loop 4:
    [leftpointer * rightpointer, leftpointer * rightpointer, rightpointer * leftpointer, rightpointer * leftpointer]

Analysis

Time Complexity: O(n)
space Complexity: O(1)
```

## Summary

| Language | Runtime | Memory  |
| :--------| :------ | :------ |
| python3  | 228 ms   | 83.1 MB |
| go       | 28 ms    | 7.2 MB  |

*With exact same code , Considering both runtime and Memory , i.e.  83/7.2 and  228/28 -> 8 x 10.5 -> 84*

>Go is 84 times efficient than python for this example.

