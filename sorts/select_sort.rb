class Array
  def select_sort
    tmp = self.dup
    res = []
    res.push(tmp.pickup_min) until tmp.empty?
    return res
  end

  def pickup_min
    min_idx  =  find_index { |item| item == self.min }
    delete_at(min_idx)
  end
end

ary = [3,4,8,7,1,6,2,11,10]
print ary.select_sort
