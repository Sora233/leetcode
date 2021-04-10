class Solution:
    def isPowerOfTwo(self, n: int) -> bool:
        a = 1
        while a <= n:
            if a == n:
                return True
            a *= 2
        return False