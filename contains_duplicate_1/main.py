from typing import List

def containsDuplicate(nums: List[int]) -> bool:
    if len(nums) ==1:
        return False
    hashmap = {}
    for num in nums:
        if num in hashmap:
            return True
        hashmap[num] =1
    return False

if __name__ == "__main__":
    assert containsDuplicate([1,2,3,1]) , "Wrong answer"