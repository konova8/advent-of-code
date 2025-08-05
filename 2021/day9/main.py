import functools

FILE_NAME = "example.txt"
FILE_NAME = "input.txt"
f = open(FILE_NAME, "r")

sol_1: int = 0
sol_2: int = 0

def sol1(data):
    res = 0
    for y, line in enumerate(data):
        for x, e in enumerate(line):
            if y-1 >= 0 and e >= data[y-1][x]:
                continue
            if y+1 < len(data) and e >= data[y+1][x]:
                continue
            if x-1 >= 0 and e >= data[y][x-1]:
                continue
            if x+1 < len(line) and e >= data[y][x+1]:
                continue
            res += e+1
            # print(f"({x=}, {y=}) = {e}")
    return res

def compute_near(p, max_x, max_y):
    x = p[0]
    y = p[1]
    res = set()
    if x+1 < max_x:
        res.add((x+1, y))
    if y+1 < max_y:
        res.add((x, y+1))
    if x-1 >= 0:
        res.add((x-1, y))
    if y-1 >= 0:
        res.add((x, y-1))
    return res

def sol2(data):
    lowest_points = set()
    for y, line in enumerate(data):
        for x, e in enumerate(line):
            if y-1 >= 0 and e >= data[y-1][x]:
                continue
            if y+1 < len(data) and e >= data[y+1][x]:
                continue
            if x-1 >= 0 and e >= data[y][x-1]:
                continue
            if x+1 < len(line) and e >= data[y][x+1]:
                continue
            lowest_points.add((x, y))

    res = {}
    for p in lowest_points:
        x = p[0]
        y = p[1]
        # print(f"Checking point ({x}, {y})")
        seen = set()
        seen.add(p)
        to_check = compute_near(p, len(data[0]), len(data))
        while len(to_check) != 0:
            p1 = to_check.pop()
            x1 = p1[0]
            y1 = p1[1]
            e1 = data[y1][x1]
            if e1 == 9:
                continue
            seen.add(p1)

            # Check if good
            if y1-1 >= 0 and e1 < data[y1-1][x1]:
                to_check.add((x1, y1-1))
            if y1+1 < len(data) and e1 < data[y1+1][x1]:
                to_check.add((x1, y1+1))
            if x1-1 >= 0 and e1 < data[y1][x1-1]:
                to_check.add((x1-1, y1))
            if x1+1 < len(data[0]) and e1 < data[y1][x1+1]:
                to_check.add((x1+1, y1))
        res[p] = len(seen)

    return functools.reduce(lambda x, y: x*y, sorted(res.values(), reverse=True)[0:3])

data = [[int(x) for x in e.rstrip()] for e in f.read().rstrip().split("\n")]

# print(data)

print(f"Q1: {sol1(data)}")
print(f"Q2: {sol2(data)}")
