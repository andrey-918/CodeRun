def min_cubes_sum_count(N):
    dp = [float('inf')] * (N + 1)  
    dp[0] = 0 

    for i in range(1, int(N ** (1/3)) + 1): 
        cube = i ** 3
        for j in range(cube, N + 1):
            dp[j] = min(dp[j], dp[j - cube] + 1)  

    return dp[N]  

N = int(input())
result = min_cubes_sum_count(N)
print(result)
