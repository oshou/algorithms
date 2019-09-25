#!/usr/bin/env python


class Node:
    def __init__(self, val):
        self.left = None
        self.right = None
        self.val = val


def binary_insert(root, node):
    if root is None:
        root = node
    else:
        if root.val > node.val:
            if root.left is None:
                root.left = node
            else:
                binary_insert(root.left, node)
        else:
            if root.right is None:
                root.right = node
            else:
                binary_insert(root.right, node)


def pre_order_print(root):
    if not root:
        return
    print(root.val)
    pre_order_print(root.left)
    pre_order_print(root.right)


def in_order_print(root):
    if not root:
        return
    in_order_print(root.left)
    print(root.val)
    in_order_print(root.right)


def post_order_print(root):
    if not root:
        return
    post_order_print(root.left)
    post_order_print(root.right)
    print(root.val)


if __name__ == '__main__':
    r = Node(3)
    binary_insert(r, Node(7))
    binary_insert(r, Node(1))
    binary_insert(r, Node(5))

print("-----pre_order_print-----")
print(pre_order_print(r))
print("-----in_order_print-----")
print(in_order_print(r))
print("-----post_order_print-----")
print(post_order_print(r))
