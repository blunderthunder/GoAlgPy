## Solution of longest substring without repeating characters

Suppose we have a function boolean allUnique(String substring) which will return true if the characters in the substring are all unique, otherwise false. We can iterate through all the possible substrings of the given string s and call the function allUnique. If it turns out to be true, then we update our answer of the maximum length of substring without duplicate characters.

Now let's fill the missing parts:

- To enumerate all substrings of a given string, we enumerate the start and end indices of them. Suppose the start and end indices are i and j, respectively. Then we have n0≤i<j≤n (here end index j is exclusive by convention). Thus, using two nested loops with i from 0 to n - 1 and j from i+1 to n, we can enumerate all the substrings of s.

- To check if one string has duplicate characters, we can use a set. We iterate through all the characters in the string and put them into the set one by one. Before putting one character, we check if the set already contains it. If so, we return false. After the loop, we return true.

## Summary

| Language | Runtime | Memory  |
| :--------| :------ | :------ |
| python3  | 64 ms   | 14.4 MB |
| go       | 4 ms    | 3.2 MB  |

*With exact same code , Considering both runtime and Memory , i.e.  64/4 and  14.4/3.2 -> 16 x 4.5 -> 72*

>Go is 72 times efficient than python for this example.