def get_answer(a, n, s, f):
    answer_list = [10**10 for _ in range (n)]
    roadList = [-1 for _ in range (n)]
    checkList = [False for _ in range (n)]
    i = s - 1
    answer_list[i] = 0

    while False in checkList:
        checkList[i] = True
        ni = i
        for j in range (n):
            if a[i][j] > 0 and j != i and a[i][j] + answer_list[i] < answer_list[j]:
                answer_list[j] = a[i][j] + answer_list[i]
                roadList[j] = i + 1
            if not checkList[j]:
                ni = j
        for j in range (n):
            if not checkList[j] and answer_list[j] < answer_list[ni]:
                ni = j
        i = ni
    if answer_list[f - 1] == 10**10:
        return -1, []
    roadAns = [f]
    i = f - 1
    while i != s - 1:
        roadAns.append(roadList[i])
        i = roadList[i] - 1
    return answer_list[f - 1], reversed(roadAns)


n = int(input())
a = []
for i in range (n):
    a.append(list(map(int, input().split())))
s, f = map(int, input().split())
_1, _2 = get_answer(a, n, s, f)
print(_1)
if _2 != [] and _1 != 0:
    print(*_2)
