class Solution:
    def isPowerOfFour(self, n: int) -> bool:
        a = 1
        while a <= n:
            if a == n:
                return True
            a *= 4
        return False