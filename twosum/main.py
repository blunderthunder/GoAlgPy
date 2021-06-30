def twoSum(nums, target):
    hashmap = {}
    for i, num in enumerate(nums):
        m = target - num
        if m in hashmap:
            return [hashmap[m], i]
        else:
            hashmap[num] = i

if __name__ == "__main__":
    print(twoSum([2,7,11,15], 9))