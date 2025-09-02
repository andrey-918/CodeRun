def get_answer(limit, my_army, health, enemy_production):
    rounds = 0
    enemy_army = 0
    while health >= limit:
        if enemy_army >= my_army:
            return 10**9
        health -= my_army - enemy_army
        enemy_army = 0
        if health >= 0:
            enemy_army += enemy_production
        rounds += 1
    while health > 0:
        if my_army <= 0:
            return 10**9
        if health >= my_army:
            health -= my_army
        else:
            enemy_army -= my_army - health
            health = 0
        my_army -= enemy_army
        if health > 0:
            enemy_army += enemy_production
        rounds += 1
    while enemy_army > 0:
        if my_army <= 0:
            return 10**9
        enemy_army -= my_army
        if enemy_army > 0:
            my_army -= enemy_army
        rounds += 1

    return rounds

x = int(input())
y = int(input())
p = int(input())
ans = 10**9

for limit in range(0, y + 2):
    ans = min(ans, get_answer(limit, x, y, p))
if ans == 10**9:
    print(-1)
else:
    print(ans)
