class Array
  def quick_sort
    return self if length <= 1
    # 任意の1要素を抽出
    base = pop
    # 選択要素を基準に大小2つの配列に分ける
    smaller,bigger = partition{|e| e < base}
    # 選択した要素をbaseにまず追加
    push(base)
    smaller.quick_sort + [base] + bigger.quick_sort
  end
end

ary = [3,4,8,7,1,6,2,11,10]
print ary.quick_sort
