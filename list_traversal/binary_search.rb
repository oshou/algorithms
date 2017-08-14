# 二分探索(バイナリサーチ)
#
# 特定の値が存在するか、リスト(配列)の先頭から末尾まで順番に検索
# 与えられる配列はソート済である前提とする

def binary_search(ary,val)
  left = 0
  right = ary.last
  mid = 0

  while left <= right do
    mid = (left + right) / 2
    if ary[mid] == val
      return "Found!!!"
    elsif ary[mid] < val
      left = mid + 1
    else
      right = mid - 1
    end
  end
end

N = 100
VAL = 70
sample_ary = (1..N).to_a
puts binary_search(sample_ary,VAL)
