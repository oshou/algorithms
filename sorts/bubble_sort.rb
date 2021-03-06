class Array
  def bubble_sort
    tmp = self.dup
    res = []
    res.push(tmp.bubbling) until tmp.empty?
    return res
  end

  def bubbling
    (self.length-1).times do |i|
      self[i],self[i+1] = self[i+1],self[i] if self[i] < self[i+1]
    end
    delete_at(-1)
  end
end

ary = [3,4,8,7,1,6,2,11,10]
print ary.bubble_sort
