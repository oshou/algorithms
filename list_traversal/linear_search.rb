# 線形探索(リニアサーチ)
#
# 特定の値が存在するか、リスト(配列)の先頭から末尾まで順番に検索

def linear_search(ary,val)
  ary.each {|num|
    return "Found!!!" if num == val
  }
end

N = 100
VAL = 80
sample_ary = (1..N).to_a.shuffle
puts linear_search(sample_ary,VAL)
