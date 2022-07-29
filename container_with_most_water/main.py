from typing import List

class Solution:
    def maxArea(self, height: List[int]) -> int:
        if len(height) <= 1:
            return 0
        if len(height) == 2:
            if height[0] > height[1]:
                return height[1]
            else:
                return height[0]
        
        ans = 0
        startidx = 0
        endidx = len(height) - 1

        while True:
            if startidx >= endidx:
                break
            
            if height[startidx] > height[endidx]:
                newans = height[endidx] * ( endidx - startidx )
                if newans > ans:
                    ans = newans
                endidx = endidx - 1
            else:
                newans = height[startidx] * ( endidx - startidx )
                if newans > ans:
                    ans = newans
                startidx = startidx + 1

        return ans



if __name__ == "__main__":
    solution = Solution()
    result = solution.maxArea([1,8,6,2,5,4,8,3,7])
    assert result==49, result