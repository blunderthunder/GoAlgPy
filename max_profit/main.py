from typing import List
import heapq

def maxProfit(prices: List[int]) -> int:
    """
        Time    O(nlogn)
	Space   O(n) heap
	8 ms, faster than 34.66%
        """
    if len(prices) < 2:
        return 0
    pq = []
    diff = 0
    for price in prices:
        heapq.heappush(pq, price)
        if price - pq[0] > diff:
            diff = price - pq[0]
    return diff


if __name__=="__main__":
    assert  maxProfit([7,1,5,3,6,4]) == 5, print("Wrong answer ")
    assert maxProfit([7,6,4,3,2]) == 0, print("Wrong answer ")