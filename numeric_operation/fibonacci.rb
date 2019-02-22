# フィボナッチ数列
#
# 第3項以降が直前の2項の和になる数列
# 例) 1,1,2,3,5,8,13,...
# - n >=0
# - F(0) = 0
# - F(1) = 1
# - F(2) = F(n-1) + F(n-2)

# 通常計算
def fib1(n)
  return n if n < 2
  p2 = 0
  p1 = 1
  2.upto(n){p2,p1 = p1,p2+p1}
  return p1
end


# 再帰
def fib2(n)
  return n if n < 2;
  return fib2(n-1) + fib2(n-2)
end


# メモ化
def fib3(n)
  return n if n < 2
  return fib3(n-1) + fib3(n-2)
end

def fib3_memo(n)
  @cache ||= []
  @cache[n] ||= fib3(n)
end

N = 40
N.times{|n| puts "fib1(#{n}): #{fib1(n)}"}
# N.times{|n| puts "fib2(#{n}): #{fib2(n)}"}
N.times{|n| puts "fib3(#{n}): #{fib3_memo(n)}"}
