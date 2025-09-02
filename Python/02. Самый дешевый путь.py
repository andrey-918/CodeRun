def get_answer(n, m, A):
    dp = [[10000000000 for i in range(m + 1)] for j in range (n + 1)]
    dp[1][1] = A[0][0]
    for i in range (1, n + 1):
        for j in range (1, m + 1):
            if i == j and i == 1:continue
            dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + A[i - 1][j - 1]
    return dp[n][m]
n, m = map(int, input().split())
A = []
for i in range (n):
    buf = list(map(int, input().split()))
    A.append(buf)
print(get_answer(n, m, A))