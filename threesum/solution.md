# 3 Sum

## Approch

* First step is to sort the given array nums. Sorting the array helps us identify duplicate triplets using our loops by skipping certain numbers that would result in duplicate triplets. This helps us avoid using a hashmap to identify the duplicates (like in solution 1) there by improving the space complexity(Keep reading to know how duplicate triplets can be skipped). Also sorting the array helps efficiently increment/decrement our index variables depending on whether the sum is less than or greater than 0.

* Next we need two loops. Outer loop index  num1Idx represents the index of the first element in the triplet. Inner loop contains two indexes num2Idx and num3Idx representing the indexes of the 2nd and 3rd triplet elements respectively.

* Initially num1Idx points to the first element in the given array and num2Idx, num3Idx point to the 2nd and last elements in the given array. We fix the outer loop index num1Idx and move the two inner loop indexes inwards as long as num3Idx > num2Idx. Once the condition num3Idx > num2Idx is false we stop the inner loop and increment the outer loop index num1Idx, also update num2Idx and num3Idx, num2Idx=num1Idx+1 and num3Idx=n-1.

* Take a variable sum to store the triplet sum. sum = nums[num1Idx] + nums[num2Idx] + nums[num3Idx]. Now there are three possibilities: a. If sum is equal to 0 we add it to our result. b. If sum is greater than 0 we need to decrease the sum value to make it equal to 0, so we decrement num3Idx index. c. If sum is less than 0 we need to increase sum value to make it equal to 0, so we increment num2Idx index.

* The inner loop should run as long as num3Idx > num2Idx for each iteration of the outer loop. We return the result once all the triplet combinations are processed.

* The above 4 steps ensure that we find all triplets whose sum is equal to 0. But it will also add duplicates to the result array. To skip duplicate triplets we need to add two conditions to our algorithm, one in the outer loop and one in the inner loop. In the outer loop if nums[num1Idx] == nums[num1Idx-1] i.e. if current num1Idx value is same as previous number (num1Idx-1) we skip the current number (we don't have to consider the current number for calculating our result). This condition ensures that we skip all duplicates from the left side of the array. Similarly to skip all numbers from the right side of the array, once we find a triplet with sum equal to zero we keep decrementing num3Idx until nums[num3Idx] != nums[num3Idx +1] (in the inner loop)

## Summary

| Language | Runtime | Memory  |
| :--------| :------ | :------ |
| python3  | 1747 ms | 18.1 MB |
| go       | 29 ms  | 7.3 MB  |
| rust     | 24 ms   | 3.4 MB  |

> Overall rust is `~430` times efficient than python3 and `~3` times efficient than golang.
