class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s) ==0:
            return 0
        if len(s) == 1:
            return 1

        hashmap = {}
        longestlen = 0
        track = 0
        for idx, char in enumerate(s):
            if char in hashmap:
                track = max(track, hashmap[char]+ 1)
            hashmap[char] = idx
            longestlen = max(longestlen, idx +1 - track)
        
        return longestlen

if __name__ == "__main__":
    sol = Solution()

    print(sol.lengthOfLongestSubstring("dvdf"))