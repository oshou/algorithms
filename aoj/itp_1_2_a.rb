a,b = gets.split(" ").map{|i| i.to_i}
if (a > b)
  puts "a > b"
elsif (a < b)
  puts "a < b"
else
  puts "a == b"
end
