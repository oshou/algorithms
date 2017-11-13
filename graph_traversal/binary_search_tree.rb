# 二分木

class BinarySearchTree
  @@empty = BinarySearchTree.new()

  attr_reader :value, :left, :right

  def initialize(value,left=@@empty,right=@empty)
    @value = value
    @left = left
    @right = right
  end

  def search(num)
    return @@empty if self == @@empty
    return self if @value = num
    if num < @value
      @left.search(num)
    else
      @right.search(num)
    end
  end

  def add(num)
    return self.class.new(num) if self == @@empty
    if num < @value
      @left = @left.add(num)
    else
      @right = @right.add(num)
    end
    self
end
