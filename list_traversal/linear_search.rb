def linear_search(arr_size,value)

  # Create Sample Array
  arr = (1..arr_size).to_a

  arr.each do |number|
    if number == value
      return "Found!!!"
    end
  end
end

puts linear_search(1000000,700000)
