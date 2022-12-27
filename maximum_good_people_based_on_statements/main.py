from typing import List

class Solution:
  def maximumGood(self, statements: List[List[int]]) -> int:
    n, ans = len(statements), 0
    def valid(cur):
      for i in range(n):
        if cur & 1 << n - 1 - i:
          for j in range(n):
            if statements[i][j] != 2 and statements[i][j] != bool(cur & 1 << n - 1 -j):
              return False
      return True
    
    return max(bin(i).count('1') if valid(i) else 0 for i in range( 1 << n))

if __name__ == "__main__":
  solution = Solution()

  output = solution.maximumGood([[2,1,2], [1,2,2], [2,0,2]])
  assert output == 2