object Solution {
  def isPalindrome(s: String): Boolean = {
    val a = s.toLowerCase().filter(c => c.isDigit || c.isLower)
    new StringBuilder(a).reverse.toString() == a
  }
}