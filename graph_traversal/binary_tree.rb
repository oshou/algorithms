# 二分木


class BinaryTree
  attr_reader :value, :left, :right

  def initialize(value,left=nil,right=nil)
    @value = value
    @left = left
    @right = right
  end

  def search(num)
    return true if @value = num
    return true if @left.search(num)
    return true if @right.search(num)
    false
  end
end
