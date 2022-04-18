## Summary

### Approch
We make use of a HashMap to solve the given problem in a different way in this approach. Firstly, we traverse over the whole list1list1 and create an entry for each element of list1list1 in a HashMap mapmap, of the form (list[i], i)(list[i],i). Here, ii refers to the index of the i^{th}ith element, and list[i]list[i] is the i^{th}ith element itself. Thus, we create a mapping from the elements of list1list1 to their indices.

Now, we traverse over list2list2. For every element ,list2[j]list2[j], of list2list2 encountered, we check if the same element already exists as a key in the mapmap. If so, it means that the element exists in both list1list1 and list2list2. Thus, we find out the sum of indices corresponding to this element in the two lists, given by sum = map.get(list[j]) + jsum=map.get(list[j])+j. If this sumsum is lesser than the minimum sum obtained till now, we update the resultant list to be returned, resres, with the element list2[j]list2[j] as the only entry in it.

If the sumsum is equal to the minimum sum obtained till now, we put an extra entry corresponding to the element list2[j]list2[j] in the resres list.

## Result

| Language | Runtime | Memory  |
| :--------| :------ | :------ |
| python3  | 316 ms  | 14.4 MB |
| go       | 26.6 ms | 7.1 MB  |
| rust     | 12 ms   | 2.4 MB  |

> Overall rust is 22 times better than python3 and ~6 times better than golang.
