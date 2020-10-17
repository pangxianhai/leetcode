# Python 实现  由于Python 能轻松画出树图形 方便查看树的构建结果

import uuid

from graphviz import Digraph


class TreeNote(object):
    def __init__(self, data, left=None, right=None, parent=None):
        self.left = left
        self.right = right
        self.data = data
        self.parent = parent


class BinarySearchTree(object):

    def __init__(self):
        self.root = None

    def search(self, data: int):
        if self.root == None:
            return None
        else:
            note = self.root
            while note:
                if note.data == data:
                    return note
                elif note.data < data:
                    note = note.right
                else:
                    note = note.left

    def insert(self, data: int):
        if self.root:
            note = self.root
            parent_note = note
            while note:
                parent_note = note
                if note.data < data:
                    note = note.right
                else:
                    note = note.left
            newNote = TreeNote(data=data, parent=parent_note)
            if parent_note.data < data:
                parent_note.right = newNote
            else:
                parent_note.left = newNote

        else:
            note = TreeNote(data=data)
            self.root = note

    def delete(self, data: int):
        if not self.root:
            return
        note = self.root
        dataNote = None
        while note:
            if note.data == data:
                dataNote = note
                break
            if data < note.data:
                note = note.left
            else:
                note = note.right
        if not dataNote:
            # 元素不存在
            return
        if not dataNote.left and not dataNote.right:
            # 叶子节点的情况
            if dataNote.parent:
                if dataNote.parent.left == dataNote:
                    dataNote.parent.left = None
                else:
                    dataNote.parent.right = None
            else:
                self.root = None

        elif dataNote.left and dataNote.right:
            # 查找后驱
            afterDataNote = dataNote.right
            while True:
                if afterDataNote.left:
                    afterDataNote = afterDataNote.left
                else:
                    break
            # 删除后驱
            if afterDataNote.parent.left == afterDataNote:
                afterDataNote.parent.left = afterDataNote.right
            else:
                afterDataNote.parent.right = afterDataNote.right
            if afterDataNote.right:
                afterDataNote.right.parent = afterDataNote.parent

            # 用后驱替换当前节点
            afterDataNote.left = dataNote.left
            afterDataNote.right = dataNote.right
            if dataNote.parent:
                if dataNote.parent.left == dataNote:
                    dataNote.parent.left = afterDataNote
                else:
                    dataNote.parent.right = afterDataNote
                afterDataNote.parent = dataNote.parent
            else:
                self.root = afterDataNote
                self.root.parent = None

        else:
            # 只有一个孩子
            child = dataNote.left
            if not child:
                child = dataNote.right
            if dataNote.parent:
                if dataNote.parent.left == dataNote:
                    dataNote.parent.left = child
                else:
                    dataNote.parent.right = child
                child.parent = dataNote.parent
            else:
                self.root = child
                self.root.parent = None

    def view(self):
        if not self.root:
            return
        treeView = Digraph(comment='Binary Search Tree')
        self.viewNote(self.root, None, treeView, None)
        treeView.view()

    def viewNote(self, note: TreeNote, parentKey, treeView, label):
        key = str(uuid.uuid1())
        treeView.node(key, str(note.data))
        if parentKey:
            treeView.edge(parentKey, key, label=label)

        if not note.left and not note.right:
            return
        if note.left:
            self.viewNote(note.left, key, treeView, "L")
        else:
            key1 = str(uuid.uuid1())
            treeView.node(key1, 'None')
            treeView.edge(key, key1, label='L')
        if note.right:
            self.viewNote(note.right, key, treeView, "R")
        else:
            key2 = str(uuid.uuid1())
            treeView.node(key2, 'None')
            treeView.edge(key, key2, label='R')


def main():
    bTree = BinarySearchTree()
    list = [13, 12, 1, 4, 0, 7, 3, 6, 9, 8, 20, 7, 15, 6, 19, 0, 4, 16, 19, 6]
    for data in list:
        bTree.insert(data)

    bTree.delete(8)
    bTree.delete(0)
    bTree.delete(20)
    bTree.delete(4)
    bTree.delete(13)
    bTree.view()


if __name__ == '__main__':
    main()

#
# dot = Digraph(comment='Binary Tree')
#
# dot.node('1', str('A'), style='filled', color='red')
# dot.node('2', str('B'), style='filled', color='red')
#
# dot.edge('1', '2', label='00')
#
# dot.view()
