k, n = map(int, input().split())
answer = []
tupik = [0 for _ in range (k + 1)]
check = 0
for _ in range (n):
    arrive, leave = map(int, input().split())
    check = 0
    for i in range (1, k + 1):
        if tupik[i] < arrive:
            answer.append(i)
            tupik[i] = leave
            check = 1
            break
    if check == 0:
        print(0, _ + 1)
        break
if check != 0:
    for a in answer:
        print(a)