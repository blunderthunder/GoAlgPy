class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        result = 0
        list_v1 = { idx: int(i) for idx, i in enumerate(version1.split("."))}
        list_v2 = { idx: int(i) for idx, i in enumerate(version2.split("."))}

        max_length = len(list_v1) if len(list_v1) > len(list_v2) else len(list_v2)

        for i in range(max_length):
            if not list_v1.get(i, False):
                list_v1[i] = 0

            if not list_v2.get(i, False):
                list_v2[i] = 0

            if list_v1[i] > list_v2[i]:
                result = 1
                break
            elif list_v2[i] > list_v1[i]:
                result = -1
                break
            else:
                continue
        
        return result


if __name__ == "__main__":
    solution = Solution()

    version1, version2 = "0.1", "1.1"
    result = solution.compareVersion(version1, version2)
    expected_result = -1

    assert result==expected_result
