# ユークリッドの互除法
#
# - 1から100までの数をプリントするプログラムを書け。
# - ただし3の倍数のときは数の代わりに「Fizz」、5の倍数のときは「Buzz」とプリントし、3と5両方の倍数の場合には「FizzBuzz」とプリントする。

# for
def fizzbuzz_for(n)
  for i in (1..n)
    if i%15 == 0
      puts "FizzBuzz!"
    elsif i%3 == 0
      puts "Fizz!"
    elsif i%5 == 0
      puts "Buzz!"
    else
      puts i
    end
  end
end

# upto
def fizzbuzz_upto(n)
  1.upto(n) do |i|
    if i%15 == 0
      puts "FizzBuzz!"
    elsif i%3 == 0
      puts "Fizz!"
    elsif i%5 == 0
      puts "Buzz!"
    else
      puts i
    end
  end
end

# each
def fizzbuzz_each(n)
  (1..n).each do |i|
    if i%15 == 0
      puts "FizzBuzz!"
    elsif i%3 == 0
      puts "Fizz!"
    elsif i%5 == 0
      puts "Buzz!"
    else
      puts i
    end
  end
end

# times
def fizzbuzz_times(n)
  (1..n).times do |i|
    if i%15 == 0
      puts "FizzBuzz!"
    elsif i%3 == 0
      puts "Fizz!"
    elsif i%5 == 0
      puts "Buzz!"
    else
      puts i
    end
  end
end

# while
def fizzbuzz_while(n)
  i = 1
  while n <= 30
    if i%15 == 0
      puts "FizzBuzz!"
    elsif i%3 == 0
      puts "Fizz!"
    elsif i%5 == 0
      puts "Buzz!"
    else
      puts i
    end
    i += 1
  end
end

N = 30
fizbuzz_for(N)
fizbuzz_upto(N)
fizbuzz_each(N)
fizbuzz_times(N)
fizbuzz_while(N)
