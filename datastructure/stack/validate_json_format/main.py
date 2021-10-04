from typing import Any


class Stack(object):
    def __init__(self) -> None:
        self.data = []

    def push(self, v) -> None:
        self.data.append(v)

    def pop(self) -> Any:
        if self.data:
            return self.data.pop()

    def is_empty(self) -> bool:
        return len(self.data) == 0


def validate_json_format(stack: Stack, chars: str) -> bool:
    lookup = {'{': '}', '[': ']', '(': ')'}
    for char in chars:
        if char in lookup.keys():
            stack.push(lookup[char])
        if char in lookup.values():
            if not stack:
                return False
            if char != stack.pop():
                return False

    if stack.is_empty():
        return True
    else:
        return False


if __name__ == '__main__':
    chars1 = "{'key1': 'value1', 'key2': [1, 2, 3], 'key3': (1, 2, 3)}"
    chars2 = "[{'key1': 'value1', 'key2': [1, 2, 3], 'key3': (1, 2, 3)}"
    stack = Stack()

    print(validate_json_format(stack, chars1))
    print(validate_json_format(stack, chars2))
