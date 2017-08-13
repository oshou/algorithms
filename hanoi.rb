RODS = ["A","B","C"]

def hanoi(n, from, to)
  return if n == 0
  temp, = RODS - [from,to]
  hanoi(n-1, from, temp)
  puts "move disk##{n} from #{from} to #{to}"
  hanoi(n-1, temp, to)
  return
end

hanoi(4,RODS.first,RODS.last)
