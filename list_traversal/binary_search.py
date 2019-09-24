#!/usr/bin/env python


class Node:
    def __init__(self, key):
        self.key = key
        self.left = None
        self.right = None


class BST:
    def __init__(self, key):
        self.root = Node(key)

    def search(self, key):
        node = self.root
        while node:
            if node.key == key:
                print("Found!")
                return node
            elif node.key > key:
                node = node.left
            else:
                node = node.right
        return node

    def insert(self, key):
        node = self.root
        while node:
            parent = node
            if node.key = key:
                print("Data already exist")
                return
            elif node.key > key:
                node = node.left
                flag = "left"
            else:
                node = node.right
                flag = "right"

        new = Node(key)
        if flag == "left":
            parent.left = new
        else:
            parent.right = new


def preorder(node):
    if node == None:
        return

    print(node.key)
    preorder(node.left)
    preorder(node.right)


def inorder(node):
    if node == None:
        return

    inorder(node.left)
    print(node.key)
    inorder(node.right)


def postorder(node):
    if node == None:
        return
    postorder(node.left)
    postorder(node.right)
    print(node.key)


root = Node('A')
root.left = Node('B')
root.right = Node('H')
root.left.left = Node('C')
root.left.right = Node('D')
root.left.right.left = Node('E')
root.left.right.right = Node('F')
root.left.right.left.left = Node('G')
print('preorder...')
preorder(root)
print()
print('inorder...')
inorder(root)
print()
print('postorder...')
postorder(root)
