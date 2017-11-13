# ツリー構造


class Tree
  attr_accessor :value,:children

  def initialize(value,children=[])
    @value = value
    @children = children
  end

  # dfs(深さ優先探索)
  def dfs(v)
    puts @value
    return true if @value == v
    @children.each {|c|
      return true if c.dfs(v)
    }
    false
  end

  # bfs(幅優先探索)
  def bfs(num)
    queue = Queue.new
    queue.enq(num)
    puts queue
    while(queue.size!=0)
      puts @value
      @children.each do |child|
        queue.push(child)
      end
    end
  end
end

t1 = Tree.new(1,[
  Tree.new(2,[
    Tree.new(5),
    Tree.new(6)
  ]),
  Tree.new(3,[
    Tree.new(7),
    Tree.new(8)
  ]),
  Tree.new(4)
])
puts t1.dfs(8)
puts t1.bfs(8)
