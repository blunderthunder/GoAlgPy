class Solution:
    def reverseBits(self, n: int) -> int:
        ret = 0
        for i in range(16):
            ret |= ((n>>i)&1)<<(31-i) | ((n>>(31-i))&1)<<i
        return ret