# Solution of Best Time to Buy and Sell

Suppose for this, Given Array

> [7, 1, 5, 3, 6, 4]

![Solution](https://leetcode.com/media/original_images/121_profit_graph.png)

The points of interest are the peaks and valleys in the given graph. We need to find the largest peak following the smallest valley. We can maintain two variables - minprice and maxprofit corresponding to the smallest valley and maximum profit (maximum difference between selling price and minprice) obtained so far respectively.

## Summary

| Language | Runtime | Memory  |
| :--------| :------ | :------ |
| python3  | 1008 ms | 25.9 MB |
| go       | 116 ms  | 8.2 MB  |

*With exact same code , Considering both runtime and Memory , i.e.  1008/116 and  25.9/8.2 -> 16 x 4.5 -> 11.84

>Go is 11.84 times efficient than python for this piece of code.