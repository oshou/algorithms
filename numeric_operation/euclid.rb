# ユークリッドの互除法
#
# 最大公約数を求める
# - R1 >= R2
# - R1 % R2 = R3
# - R2 % R3 = R4 同様にRn = 0になるまで繰り返し
# - Rn-1が最大公約数

# 通常計算
def euclid1(a,b)
  a,b = b,a%b while b > 0
  return a
end

# 再帰
def euclid2(a,b)
  return a if b == 0
  return euclid2(b,a%b)
end

puts euclid1(1920,1080)
puts euclid2(1920,1080)
