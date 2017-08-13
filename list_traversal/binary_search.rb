# prerequisite: sorted

def binary_search(arr_size,value)
  arr = (1..arr_size).to_a

  left = 0
  right = arr.last
  mid = 0

  while left <= right do
    mid = (left + right) / 2
    if arr[mid] == value
      return "Found!!!"
    elsif arr[mid] < value
      left = mid + 1
    else
      right = mid - 1
    end
  end
end

puts binary_search(1000000,700000)
