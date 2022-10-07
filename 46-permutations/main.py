class Solution:

    def permute(self, nums):
        return self._recombo([], nums)

    @staticmethod
    def _recombo(current, left):
        result = []
        for i, x in enumerate(left):
            new_current = current + [x]
            new_left = left.copy()
            del new_left[i]

            if len(new_left) == 0:
                result.append(new_current)
            else:
                result += Solution._recombo(new_current, new_left)

        return result


if __name__ == "__main__":
    s = Solution()
    print(s.permute([1, 2, 3]))
