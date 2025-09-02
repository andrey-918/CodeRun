N = int(input())
answer = 0
for i in range (1, N + 1):
    answer += (N - i + 2) * (N - i + 1)
    if (N - 2 * i + 1) > 0:
        answer += (N - 2 * i + 2) * (N - 2 * i + 1)
print(answer // 2)