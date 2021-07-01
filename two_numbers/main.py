import sys

class ListNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next    

class Solution:

    @staticmethod
    def create_ListNode(arr : list):
        l = None
        while len(arr):
            l = ListNode(arr.pop(), l)
        return l

    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        carry = 0
        root = n = ListNode(0)
        while l1 or l2 or carry:
            v1 = v2 = 0
            if l1:
                v1 = l1.val
                l1 = l1.next
            if l2:
                v2 = l2.val
                l2 = l2.next
            carry, val = divmod(v1+v2+carry, 10)
            n.next = ListNode(val)
            n = n.next
        return root.next


if __name__ == "__main__":
    sol = Solution()
    l1 = Solution.create_ListNode([2,4,3])
    l2 = Solution.create_ListNode([5,6,4])
    l = sol.addTwoNumbers(l1, l2)
    print(l.next, l.val)