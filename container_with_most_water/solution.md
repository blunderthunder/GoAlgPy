# Container with most water

## Approch

We can solve this with 2 general approch:
* **Brute Force approch**  
    * As the name suggest with this method we multiply every possible combination of pillers in height array and return maximum height.
    * Time complexity: O(n^2).

* **Pointer Based Approch**  
    * We took 2 pointers start & end.
    * Calculate the area with the minimum height multiplied by distance.
    * Store the maximum to ans.
    * Move the shorter height among 2 pointers to the next position.
    * Time complexity: O(n).

## Summary


| Language | Runtime | Memory  |
| :--------| :------ | :------ |
| python3  | 819 ms | 27 MB |
| go       | 74 ms  | 8.7 MB  |
| rust     | 11 ms   | 3 MB  |

> Overall rust is `~670` times better than python3 and `~15` times better than golang.
