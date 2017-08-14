# ハノイの塔
#
# 3本のポールとサイズの異なるn個のドーナツ円盤がある
# 出発点のポールに固まっている円盤を全て目的のポールに移す。

POLES = ["A","B","C"]

def hanoi(n, from, to)
  return if n == 0
  temp = POLES - [from,to].to_s
  hanoi(n-1, from, temp)
  puts "move disk##{n} from #{from} to #{to}"
  hanoi(n-1, temp, to)
  return
end

hanoi(4,POLES.first,POLES.last)
