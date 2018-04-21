S = gets.to_i
h = S / 3600
m = S % 3600 / 60
s = S % 3600 % 60
puts "#{h}:#{m}:#{s}";
