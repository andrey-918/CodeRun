from collections import deque
N = int(input())
head = deque()
tail = deque()
count = 0
v = 0
for i in range (N):
    A = list(input().split())
    if A[0] == '+':
        tail.append(A[1])
    elif A[0] == '*':
        tail.appendleft(A[1])

    else:
        print(head[0])
        head.popleft()
    if len(head) < len(tail):
        head.append(tail[0])
        tail.popleft()

