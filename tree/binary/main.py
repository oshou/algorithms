class Node:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None


class BinaryTree:
    def __init__(self, root_value):
        self.root = Node(root_value)

    # root -> left -> right
    def preorder_print(self, root):
        print(root.value)
        if root.left is not None:
            self.preorder_print(root.left)
        if root.right is not None:
            self.preorder_print(root.right)

    # left -> root -> right
    def inorder_print(self, root):
        if root.left is not None:
            self.inorder_print(root.left)
        print(root.value)
        if root.right is not None:
            self.inorder_print(root.right)

    # root -> right -> left
    def postorder_print(self, root):
        print(root.value)
        if root.right is not None:
            self.postorder_print(root.right)
        if root.left is not None:
            self.postorder_print(root.left)


bt = BinaryTree(1)
bt.root.left = Node(2)
bt.root.right = Node(3)

print("pre")
bt.preorder_print(bt.root)
print("in")
bt.inorder_print(bt.root)
print("post")
bt.postorder_print(bt.root)
