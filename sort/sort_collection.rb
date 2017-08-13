class SortCollections
  def setup(n)
    @ary = []
    n.times { @ary << rand(n)}
    return @ary
  end

  # BubbleSort
  def bubble_sort(ary)
    return ary if ary.length <= 1
    res = ary.dup
    for i in 0..(res.length - 1)
      for j in 0..(res.length - 2 - i)
        res[j],res[j+1] = res[j+1],res[j] if res[j] > res[j+1]
      end
    end
    return res
  end

  # MergeSort
  def merge_sort(ary)
    return ary if ary.length <= 1
    mid = ary.length / 2
    left = merge_sort(ary[0...mid])
    right = merge_sort(ary[mid..-1])
    merge(left,right)
  end

  def merge(a,b)
    res = []
    until a.empty? && b.empty?
      res <<
      case
      when a.empty? then b.shift
      when b.empty? then a.shift
      when a[0] < b[0] then a.shift
      else b.shift
      end
    end
    return res
  end

  # QuickSort
  def quick_sort(ary)
    return ary if ary.length <= 1
    pivot = ary[0]
    left = []
    right = []
    (1..ary.length-1).each do |i|
      if ary[i] <= pivot
        left.push(ary[i])
      else
        right.push(ary[i])
      end
    end
    left = ary.quick_sort(left)
    right = ary.quick_sort(right)
    return left + [pivot] + right
  end
end
