def into_seconds(time):
    seconds = int(time[-2:])
    seconds += int(time[-5: -3]) * 60
    seconds += int(time[: 2]) * 3600
    return seconds

list_time = []
list_speed = []
checking = False
time09 = into_seconds('09;00;00')
time13 = into_seconds('13;00;00')
time14 = into_seconds('14;00;00')
time18 = into_seconds('18;00;00')
N = int(input())
dp = dict()
for i in range (N):
    time, speed = map(str, input().split())
    list_time.append(into_seconds(time))
    list_speed.append(int(speed))
index = -1
list_time.append(time18 + 1)
for i in range (time09, time18 + 1):
    if list_time[index + 1] == i:
        index += 1
    speed = list_speed[index]
    if i > time13 and i < time14:
        checking = True
        dp[i] = dp[i - 1]
        continue
    if i in dp:
        dp[i] = max(dp[i - 1], dp[i])
    elif i - 1 in dp:
        dp[i] = dp[i - 1]
    else:
        dp[i] = 0
    if i + speed <= time13 * (not (checking)) + time18 * checking:
        if i + speed in dp:
            dp[i + speed] = max(dp[i + speed], dp[i] + 1)
        else:
            dp[i + speed] = dp[i] + 1
# print(dp[time13])
print(dp[time18])