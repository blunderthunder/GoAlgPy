from typing import List

class Solution:
    def hitBricks(self, grid: List[List[int]], hits: List[List[int]]) -> List[int]:
        # one method with 2 usages
        # first call: mark all connected points as 2
        # second call: count total number of newly connected points
        def dfs(x, y):
            if grid[x][y] != 1: return 0 
            grid[x][y], ans = 2, 1
            for xx, yy in map(lambda pair: (pair[0]+x, pair[1]+y), ((-1, 0), (1, 0), (0, -1), (0, 1))):
                if 0 <= xx < m and 0 <= yy < n and grid[xx][yy] == 1: ans += dfs(xx, yy)
            return ans                    
        
        # given (x, y) mark it as 1; then return True if (x, y) is stable        
        def is_stable(x, y):
            grid[x][y] += 1
            if grid[x][y] <= 0: return False
            if (x == 0 and grid[x][y] == 1) or grid[x][y] == 2: return True 
            return any((0 <= xx < m and 0 <= yy < n) and grid[xx][yy] == 2 for xx, yy in map(lambda pair: (pair[0]+x, pair[1]+y), ((-1, 0), (1, 0), (0, -1), (0, 1))))
        
        m, n = len(grid), len(grid[0])
        # decrement hits to make them unstable
        # this is NOT same as grid[x][y] = 0 
        #   because grid[x][y] will be incremented during is_stable() 
        #   and the value of grid[x][y] will be used to check whether its stable
        # if you don't do this, you will need to have a set() to record 
        #   which (x,y) in `hits` was originally stable (grid[x][y] == 1)
        for x, y in hits: grid[x][y] -= 1
        # dfs first call: mark all connected points as 2    
        for j in range(n): dfs(0, j)  
            
        # dfs second call: count total number of newly connected points
        return [(dfs(x,y) - 1) if is_stable(x, y) else 0 for x, y in hits[::-1]][::-1]  


if __name__ == "__main__":
    solution = Solution()

    grid = [[1,0,0,0],[1,1,1,0]]
    hits = [[1,0]]

    result = solution.hitBricks(grid, hits)
    print(" the result is {} ".format(result))

    result = solution.hitBricks([[1,0,0,0],[1,1,0,0]], [[1,1],[1,0]])
    print(" the result is {} ".format(result))

    result = solution.hitBricks([[1,0,1],[1,1,1]], [[0,0],[0,2],[1,1]])
    print(" the result is {} ".format(result))
