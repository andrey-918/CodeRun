heights = list(map(int, input().split()))
length = heights[0]
heights.pop(0)
mass = [length for _ in range (length)]
stack = []
for i in range (length):
    while stack != [] and heights[stack[-1]] > heights[i]:
        mass[stack[-1]] = i
        stack.pop()
    stack.append(i)
stack.clear()
for i in range (length - 1, -1, -1):
    while stack != [] and heights[stack[-1]] > heights[i]:
        mass[stack[-1]] -= i + 1
        stack.pop()
    stack.append(i)
answer = -1
for i in range (length):
    answer = max(answer, mass[i] * heights[i])
print(answer)