with open('input.txt', 'r') as f:
    a = f.read().split()
    print(len(set(a)))