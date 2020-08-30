# Python 实现  由于Python 能轻松画出树图形 方便查看树的构建结果

import uuid

from graphviz import Digraph


# 红黑树
# 每个节点红色或黑色
# 根节点是黑色
# 叶子节点(Nil)是黑色
# 如果节点为红色 则它的字节点为黑色
# 对每个节点，从该节点开始到其子孙节点的所有路径上包含相同数目的黑色节点

class TreeNote(object):
  def __init__(self, data, isRed=True, left=None, right=None, parent=None):
    self.left: TreeNote = left
    self.right: TreeNote = right
    self.data: int = data
    self.parent: TreeNote = parent
    self.red: bool = isRed


class RedBlackTree(object):
  def __init__(self):
    self.root: TreeNote = None

  def search(self, data: int):
    note = self.root
    if not note:
      return None
    while note:
      if note.data == data:
        return data
      elif data < note.data:
        note = note.left
      else:
        note = note.right
    return None

  def insert(self, data: int):
    if not self.root:
      self.root = TreeNote(data, False)
      return
    else:
      p = self.root
      while True:
        if data <= p.data:
          if p.left == None:
            break
          else:
            p = p.left
        else:
          if p.right == None:
            break
          else:
            p = p.right
      newNote = TreeNote(data)  # 默认红色
      newNote.parent = p
      if newNote.data <= p.data:
        p.left = newNote
      else:
        p.right = newNote
      # 普通二叉树到此可结束 但红黑树 由于添加新节点 为了保证红黑树的性质需要对树进行旋转
      self.rebuildTree(newNote)

  def rebuildTree(self, z: TreeNote):
    if not z.parent:
      z.red = False
      return
    if not z.parent.red:
      return
    p = z.parent
    pp = p.parent
    if pp is None:
      # 这种情况不存在
      return
    if p == pp.left:
      pr = pp.right
      if pr and pr.red:
        p.red = False
        pr.red = False
        pp.red = True
        self.rebuildTree(pp)
        return
      else:
        if z == p.right:
          self.leftRotate(p)
          p = z
        self.rightRotate(pp)
        p.red = False
        pp.red = True
        return
    else:
      pl = pp.left
      if pl and pl.red:
        p.red = False
        pl.red = False
        pp.red = True
        self.rebuildTree(pp)
        return
      else:
        if z == p.left:
          self.rightRotate(p)
          p = z
        self.leftRotate(pp)
        p.red = False
        pp.red = True
        return

  def delete(self, data: int):
    note = self.root
    dataNote: TreeNote = None
    while note:
      if note.data == data:
        dataNote = note
      if data <= note.data:
        note = note.left
      else:
        note = note.right
    if not dataNote:
      # 数据不存在 删除失败
      return
    if dataNote.left and dataNote.right:
      afDataNote = dataNote.right
      while True:
        if afDataNote.left:
          afDataNote = afDataNote.left
        else:
          break
      dataNote.data = afDataNote.data
      dataNote = afDataNote

    child = dataNote.left
    if not child:
      child = dataNote.right

    if dataNote.parent:
      if dataNote.parent.left == dataNote:
        dataNote.parent.left = child
      else:
        dataNote.parent.right = child
    else:
      self.root = child

    if child:
      delete = not child.red
      child.parent = dataNote.parent
      child.red = dataNote.red
      if delete:
        self.rebuildDeleteTree(child, dataNote.parent)
    else:
      if not dataNote.red:
        self.rebuildDeleteTree(child, dataNote.parent)

  def rebuildDeleteTree(self, z: TreeNote, p: TreeNote):
    if p is None:
      self.root = z
      if self.root:
        self.root.parent = None
      return
    if self.root == z or (z is not None and z.red):
      return
    if z == p.left:
      if p.right is None:
        p.red = False
        self.rebuildDeleteTree(p, p.parent)
        return
      else:
        if p.right.red:
          self.leftRotate(p)
          p.red = True
          p.parent.red = False
          self.rebuildDeleteTree(z, p)
          return
        else:
          r = p.right
          rLb = r.left is None or not r.left.red
          rRb = r.right is None or not r.right.red
          if rLb and rRb:
            r.red = True
            if p.red:
              p.red = False
            else:
              self.rebuildDeleteTree(p, p.parent)
            return
          elif rRb:
            self.rightRotate(r)
            # 由于旋转 右节点重新定义
            r = p.right
            r.red = False
            if r.right:
              r.right.red = True
            self.leftRotate(p)
            p.parent.red = p.red
            p.parent.right.red = False
            p.red = False
          else:
            self.leftRotate(p)
            p.parent.red = p.red
            p.parent.right.red = False
            p.red = False
    else:
      if p.left is None:
        p.red = False
        self.rebuildDeleteTree(p, p.parent)
        return
      else:
        if p.left.red:
          self.rightRotate(p)
          p.red = True
          p.parent.red = False
          self.rebuildDeleteTree(z, p)
          return
        else:
          l = p.left
          lLb = l.left is None or not l.left.red
          lRb = l.right is None or not l.right.red
          if lLb and lRb:
            l.red = True
            if p.red:
              p.red = False
            else:
              self.rebuildDeleteTree(p, p.parent)
            return
          elif lLb:
            self.leftRotate(l)
            l = p.left
            l.red = False
            if l.left:
              l.left.red = True
            self.rightRotate(p)
            p.parent.red = p.red
            p.parent.left.red = False
            p.red = False
          else:
            self.rightRotate(p)
            p.parent.red = p.red
            p.parent.left.red = False
            p.red = False

  # 以为中心左旋 x
  def leftRotate(self, x: TreeNote):
    if x is None:
      return
    p = x.parent
    r = x.right
    if r is None:
      return
    x.right = r.left
    x.parent = r
    if x.right:
      x.right.parent = x
    r.left = x
    if p is None:
      self.root = r
      self.root.parent = None
    else:
      r.parent = p
      if p.left == x:
        p.left = r
      else:
        p.right = r

  def rightRotate(self, x: TreeNote):
    if x is None:
      return
    l = x.left
    p = x.parent
    if l is None:
      return
    x.parent = l
    x.left = l.right
    if x.left:
      x.left.parent = x
    l.right = x
    l.parent = p
    if p is None:
      self.root = l
      self.root.parent = None
    else:
      if p.left == x:
        p.left = l
      else:
        p.right = l

  def view(self, filename):
    if not self.root:
      return
    treeView = Digraph(comment='Binary Search Tree')
    self.viewNote(self.root, None, treeView, None)
    treeView.view(filename=filename)

  def viewNote(self, note: TreeNote, parentKey, treeView, label):
    key = str(uuid.uuid1())
    color = 'red'
    if not note.red:
      color = 'black'
    treeView.node(key, str(note.data), style='filled', color=color,
                  fontcolor='white')
    if parentKey:
      treeView.edge(parentKey, key, label=label)

    if not note.left and not note.right:
      return
    if note.left:
      note = note.left.parent  # 校验 parent 指针是否正确
      self.viewNote(note.left, key, treeView, "L")
    if note.right:
      note = note.right.parent  # 校验 parent 指针是否正确
      self.viewNote(note.right, key, treeView, "R")


def main():
  rbTree = RedBlackTree()
  list = [13, 12, 1, 4, 0, 7, 3, 6, 9, 8, 20, 7, 15, 6, 19, 0, 4, 16, 19, 6]
  # list = [13, 12, 1, 4]
  for data in list:
    rbTree.insert(data)
  # rbTree.view("begind")
  # print("开始删除0")
  # rbTree.delete(0)
  # print("开始删除1")
  # rbTree.delete(1)
  # print("开始删除0")
  # rbTree.delete(0)
  # #
  # # # rbTree.leftRotate(rbTree.root)
  # print("开始删除4")
  # rbTree.delete(4)
  # rbTree.delete(4)
  # rbTree.delete(3)

  print("开始删除12")
  rbTree.delete(12)
  print("开始删除15")
  rbTree.delete(15)
  print("开始删除13")
  rbTree.delete(13)
  print("开始删除19")
  rbTree.delete(19)
  print("开始删除19")
  rbTree.delete(19)
  print("开始删除20")
  rbTree.delete(20)
  # # rbTree.view("a")
  print("开始删除9")
  rbTree.delete(9)
  rbTree.delete(6)
  rbTree.delete(7)
  rbTree.delete(8)
  rbTree.delete(6)
  rbTree.delete(7)
  rbTree.delete(6)
  rbTree.delete(16)
  rbTree.delete(4)
  rbTree.delete(4)
  rbTree.delete(3)
  rbTree.delete(0)
  rbTree.delete(1)
  rbTree.delete(0)
  rbTree.view("finalResult")


if __name__ == '__main__':
  main()
