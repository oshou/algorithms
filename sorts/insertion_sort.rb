class Array
  def insert_sort
    # default: []
    inject([]){|mem,var| mem.insert_with_order(var)}
  end

  def insert_with_order(item)
    pos = find_index{|n| item <= n} || length
    insert(pos,item)
  end
end
