class Stack:
    def __init__(self):
        self.stack = []

    def push(self, x):
        self.stack.append(x)
        return stack.stack

    def pop(self):
        self.stack.pop(-1)
        return stack.stack


count = int(input("任意の個数を入力して下さい: "))
stack = Stack()
for i in range(count):
    stack.push(i)

for i in range(count):
    stack.pop()
