from typing import List

def maxSubArray(nums: List[int]) -> int:
    if len(nums) == 1:
        return nums[0]
    maxsofar , maxendinghere = nums[0],nums[0]
    for i in range(1,len(nums)):
        maxendinghere = max(nums[i], maxendinghere + nums[i])
        maxsofar=max(maxsofar, maxendinghere)
    return maxsofar
        

if __name__ == "__main__":
    assert maxSubArray([-2,1,-3,4,-1,2,1,-5,4]) == 6, " Wrong Answer"