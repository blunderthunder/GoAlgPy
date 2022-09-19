# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution(object):

    def allPossibleFBT(self, N):
        dp = { 0: [], 1: [TreeNode()]}
        def backtrack(n):
            if n in dp:
                return dp[n]
            
            res = []
            for l in range(n):
                r = n - 1 - l
                lefttrees, righttrees =  backtrack(l), backtrack(r)

                for left in lefttrees:
                    for right in righttrees:
                        res.append(TreeNode(0, left, right))
            
            dp[n] = res
            return res

        return backtrack(N)


if __name__ == "__main__":
    sol = Solution()

    print(sol.allPossibleFBT(7))