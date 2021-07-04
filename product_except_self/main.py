from typing import List

def productExceptSelf(nums: List[int]) -> List[int]:
    if not nums: return False
        
    arr = [1] * len(nums)
    pi = pj = 1

    for i in range(len(nums)):
        j = -1-i

        arr[i]*=pi; arr[j]*=pj
        pi *= nums[i]; pj *= nums[j]
        print(arr)
                
    return arr

if __name__ == "__main__":
    assert productExceptSelf([1,2,3,4]) == [24, 12, 8, 6], "Result Not Matched"