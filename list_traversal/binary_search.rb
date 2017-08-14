# prerequisite: sorted

def binary_search(ary_size,value)
  ary = (1..ary_size).to_a.shuffle

  left = 0
  right = ary.last
  mid = 0

  while left <= right do
    mid = (left + right) / 2
    if ary[mid] == value
      return "Found!!!"
    elsif ary[mid] < value
      left = mid + 1
    else
      right = mid - 1
    end
  end
end

puts binary_search(1000000,700000)
