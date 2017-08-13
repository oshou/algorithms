#            0
#            |
#       +----------+
#       |          |
#    +--1--+    +--4--+
#    |     |    |     |
#    2     3    5     6


class Node
  def initialize(name, *nodes)
    @name = name
    @nodes = nodes
    @adjacent = nodes
    @visited_flag = false
  end

  def visit
    puts @name
  end

  def adjacent()
    @nodes
  end

  def visited
    @visited_flag = true
  end

  def visited?
    return @visited_flag
  end
end

def depth_first_search(node)
  node.visit
  node.visited
  node.adjacent.each do |child|
    if child.visited? == false
      depth_first_search(child)
    end
  end
end

node3 = Node.new(3)
node2 = Node.new(2)
node1 = Node.new(1,node2,node3)

node5 = Node.new(5)
node6 = Node.new(6)
node4 = Node.new(4,node5,node6)

node0 = Node.new(0,node1,node4)

depth_first_search(node0)
