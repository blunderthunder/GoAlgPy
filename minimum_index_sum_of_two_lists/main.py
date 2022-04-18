from typing import List

class Solution:
    def findRestaurant(self, list1: List[str], list2: List[str]) -> List[str]:
        common = set(list1).intersection(list2)
        if len(common) == 1:
            return list(common)
        item_index = {}
        for item in list(common):
            item_index[list1.index(item) + list2.index(item)] = item_index.get(list1.index(item) + list2.index(item), [])
            item_index[list1.index(item) + list2.index(item)].append(item)
        return item_index[sorted(list(item_index.keys()))[0]]


if __name__ == "__main__":
    solution = Solution()

    val1 = ["Shogun","Tapioca Express","Burger King","KFC"]
    val2 = ["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"]
    assert solution.findRestaurant(val1, val2) == ["Shogun"]