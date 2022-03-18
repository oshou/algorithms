const isParindrome = (text) => {
  let left = 0;
  let right = text.length
  while (left < right) {
    if (text[left] !== text[right]) return false
    left += 1
    right -= 1
  }
  return true
};
console.log(isParindrome("abba"))
