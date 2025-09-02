n, a, b = map(int, input().split())
dp = dict()
dp[1] = 0
for i in range(1, n + 1):
    for j in range (1, i):
        if i in dp:
            dp[i] = min(dp[i], max(dp[j] + a, dp[i - j] + b))
        else:
            dp[i] = max(dp[j] + a, dp[i - j] + b)
print(dp[n])