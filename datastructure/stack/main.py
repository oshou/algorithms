from typing import Any


class Stack(object):
    def __init__(self) -> None:
        self.data = []

    def push(self, data) -> None:
        self.data.append(data)

    def pop(self) -> Any:
        if self.data:
            return self.data.pop()


if __name__ == '__main__':
    stack = Stack()
    print(stack.data)
    stack.push(1)
    print(stack.data)
    stack.push(2)
    print(stack.data)
    stack.push(3)
    print(stack.data)

    print(stack.pop())
    print(stack.data)
    print(stack.pop())
    print(stack.data)
    print(stack.pop())
    print(stack.data)
    print(stack.pop())
    print(stack.data)
