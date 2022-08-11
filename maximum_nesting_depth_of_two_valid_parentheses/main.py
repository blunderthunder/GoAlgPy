from typing import List

class Solution:
    def maxDepthAfterSplit(self, seq: str) -> List[int]:
        depth = 0
        ans = []
        for val in seq:
            if val == "(":
                ans.append(depth % 2)
                depth += 1
            else:
                depth -= 1
                ans.append(depth % 2)
        
        return ans


if __name__ == "__main__":
    solution = Solution()
    assert solution.maxDepthAfterSplit("(()())")==[0,1,1,1,1,0]