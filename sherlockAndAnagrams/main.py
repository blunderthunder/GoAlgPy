from collections import Counter
from itertools import combinations
val = "cdcd"

substrings = []
# size of the window must be lower than length of string and greater than equal to 1
for i in range (1, len(val)):
    start = 0
    end = start + i
    while True:
        if end > len(val):
            break
        data = val[start:end]
        start += 1
        end = start + i
        substrings.append(''.join(sorted(data)))

combinators  = filter(lambda x: x >1 , list(Counter(substrings).values()))

print(sum([len(list(combinations(list(range(i)), 2))) for i in combinators]))