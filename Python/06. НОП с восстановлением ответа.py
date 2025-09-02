n1 = int(input())
a1 = list(map(str, input().split()))
n2 = int(input())
a2 = list(map(str, input().split()))
answer = []
slovar = dict()
count = 0
matrix = [[0 for _ in range(n1 + 1)] for __ in range(n2 + 1)]
for i in range(1, n2 + 1):
    check = True
    for j in range(1, n1 + 1):
        matrix[i][j] = max(matrix[i - 1][j] , matrix[i][j - 1] )
        if a2[i - 1] == a1[j - 1]:
            matrix[i][j] = matrix[i - 1][j - 1] + 1
# for i in matrix:
#     print(i)
I = n2
J = n1
while I != 0 and J != 0:
    if matrix[I][J] == matrix[I - 1][J]:
        I -= 1
    elif matrix[I][J] == matrix[I][J - 1]:
        J -= 1
    else:
        answer.append(a1[J - 1])
        I -= 1
        J -= 1
answer.reverse()
for i in answer: print(i, end = ' ')
