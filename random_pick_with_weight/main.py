from typing import List
import random
class Solution:

    def __init__(self, w: List[int]):
        self.prefixes = []
        self.total = 0
        for weight in w:
            self.total += weight
            self.prefixes.append(self.total)

    def pickIndex(self) -> int:
        target = self.total * random.random()
        
        l, r = 0, len(self.prefixes) - 1
        
        while l <= r:
            mid = (l + r) // 2
            if self.prefixes[mid] <= target:
                l = mid + 1
            else:
                r = mid - 1
                
        return l


if __name__ == "__main__":
  sol = Solution([1,3, 2, 1, 4])