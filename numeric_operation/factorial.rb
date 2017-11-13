# 階乗計算
# 例）1,1,2,6,24,120...
#
# - n >= 1
# - n! = n * (n-1)!

# 再帰
def factorial(n)
  return 1 if n == 0
  return n * factorial(n-1)
end

p factorial(7)
