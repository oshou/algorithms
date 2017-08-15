# 二分探索(バイナリサーチ)
#
# iリストのインデックスの中央値を計算し、その大小で対象の絞込を行う。
# 要素が一致したら終了。

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
