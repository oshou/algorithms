class Array
  def select_sort
    tmp = self.dup
    res = []
    res.push(tmp.pickup_min) until tmp.empty?
    return res
  end

  def pickup_min
    (self.length-1).times do |i|
      self[i],self[i+1] = self[i+1],self[i] if self[i] < self[i+1]
    end
    delete_at(-1)
  end
end
