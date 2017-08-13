def factorial(n)
  return 1 if n == 0
  return n * factorial(n-1)
end

p factorial(5)
