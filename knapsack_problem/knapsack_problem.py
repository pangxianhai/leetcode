import pprint

#商品重量
w = [3,6,3,8,6]
#商品价值
v = [4,6,6,12,10]
# 背包容积
W = 10

c = [ [0] * (W+1) for _ in range(0,len(w)+1)]

for i in range(1,len(w)+1):
    for j in range(1,W+1):
        if j - w[i-1] >= 0:
            c[i][j] = max(c[i-1][j],c[i-1][j-w[i-1]]+v[i-1])
        else:
            c[i][j] = c[i-1][j]

i,j = len(w),W

print("最优解:",c[i][j])

choice_goods = []
while i>=0 and j>=0:
    if c[i][j] == c[i-1][j]:
        print("商品",i,"不选")
        i = i - 1
    else:
        choice_goods.append(i)
        print("商品",i,"选")
        j = j - w[i-1]
        i = i - 1

print("选择商品列表",choice_goods)
