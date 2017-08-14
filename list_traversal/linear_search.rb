def linear_search(value)
  self.each do |number|
    if number == value
      return "Found!!!"
    end
  end
end

def create_random_array(ary_size)
  return (1..ary_size).to_a.shuffle
end

N = 100
sample_ary = create_random_array(N)
print sample_ary
sample_ary.linear_search(80)
