class Array
  def merge_sort
    tmp = self.dup
    return tmp if tmp.length <= 1
    a,b = self.half.map{|e| e.merge_sort}
    merge(a,b)
  end

  def half
    mid = length/2
    return slice(0..mid),slice(mid..-1)
  end

  def merge(a,b)
    res = []
    until a.empty? && b.empty?
      res <<
        case
        when a.empty? then b.shift
        when b.empty? then a.shift
        when a.first < b.first then a.shift
        else b.shift
        end
    end
    res
  end
end

ary = [3,4,8,7,1,6,2,11,10]
print ary.merge_sort
