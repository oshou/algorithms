from typing import Any


class Queue(object):
    def __init__(self) -> None:
        self.data = []

    def enqueue(self, v) -> None:
        self.data.append(v)

    def dequeue(self) -> Any:
        if self.data:
            v = self.data[0]
            self.data = self.data[1:]
            return v


if __name__ == '__main__':
    queue = Queue()

    queue.enqueue(1)
    print(queue.data)
    queue.enqueue(2)
    print(queue.data)
    queue.enqueue(3)
    print(queue.data)

    print(queue.dequeue())
    print(queue.data)
    print(queue.dequeue())
    print(queue.data)
    print(queue.dequeue())
    print(queue.data)
    print(queue.dequeue())
    print(queue.data)
