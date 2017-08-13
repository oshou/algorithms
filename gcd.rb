def gcd(a,b)
  return a.abs if b == 0
  return gcd(b,a%b)
end

p gcd(1920,1080)
