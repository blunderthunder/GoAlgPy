from typing import List

class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        # sort input array
        nums.sort()
        result = []

        if len(nums) < 3:
            return [[]]
        if len(nums) == 3:
            if (nums[0] + nums[1] + nums[2]) == 0:
                return [[nums[0], nums[1], nums[2]]]

        for num1Idx in range(0,len(nums) - 2 ):
            
            # skill the the duplicates from left
            if num1Idx > 0 and ( nums[num1Idx] == nums[num1Idx - 1]):
                continue

            num2Idx = num1Idx + 1
            num3Idx = len(nums) - 1

            while num2Idx < num3Idx:
                sum = nums[num1Idx] + nums[num2Idx] + nums[num3Idx]

                if sum == 0:
                    result.append([nums[num1Idx], nums[num2Idx], nums[num3Idx]])

                    num3Idx -= 1

                    # skip all duplicates from right
                    while num2Idx < num3Idx and nums[num3Idx] == nums[ num3Idx + 1]:
                        num3Idx -= 1
                elif sum > 0:
                    num3Idx -= 1
                else:
                    num2Idx += 1
        
        return result
        
        

if __name__ == "__main__":
    solution = Solution()
    # input array 
    nums = [-1,0,1,2,-1,-4]
    result = [[-1,-1,2],[-1,0,1]]
    assert solution.threeSum(nums)==result