n, k = map(int, input().split())
A = list(map(int, input().split()))
stack = []
count = 0
for j in range (n):
    i = A[j]
    if count == k:
        print(stack[0])
        if stack[0] == A[j - k]:
            stack.pop(0)
    while stack != [] and stack[-1] > i:
        stack.pop()
    stack.append(i)

    if count != k: count += 1

print(stack[0])
