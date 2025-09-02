def get_answer(n, A):
    dp = [0 for i in range(n)]
    dp[0] = 1000000000000
    dp[1] = A[1] - A[0]
    for i in range(2, n):
        dp[i] = min(dp[i - 2], dp[i - 1]) + A[i] - A[i - 1]
    return dp[-1]
n = int(input())
A = list(map(int, input().split()))
A = sorted(A)
print(get_answer(n, A))