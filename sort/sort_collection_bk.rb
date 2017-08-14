require './bubble_sort.rb'

class Array
  def initialize(n)
    @n = n
  end

  def setup_random_ary
    @ary = []
    @n.times { @ary << rand(@n)}
    return(@ary)
  end
end

sample_ary = Array.new(100).setup_random_ary
print sample_ary
bubble_sort(sample_ary)
