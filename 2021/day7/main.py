FILE_NAME = "example.txt"
FILE_NAME = "input.txt"
f = open(FILE_NAME, "r")

sol_1: int = 0
sol_2: int = 0

def sol1(data):
    best = -1
    for e in data:
        cost = 0
        for c in data:
            cost += abs(c - e)
        if best == -1 or cost < best:
            best = cost
    return best

def sol2(data):
    best = -1
    for e in range(min(data), max(data)+1):
        cost = 0
        for c in data:
            cost += (abs(c-e) * (abs(c-e)+1)) // 2
        if best == -1 or cost < best:
            best = cost
    return best

data = [int(e.rstrip()) for e in f.read().rstrip().split(",")]

print(f"Q1: {sol1(data)}")
print(f"Q2: {sol2(data)}")
