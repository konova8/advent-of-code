FILE_NAME = "example.txt"
FILE_NAME = "input.txt"
f = open(FILE_NAME, "r")

sol_1: int = 0
sol_2: int = 0

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
    if x+1 < max_x and y+1 < max_y:
        res.add((x+1, y+1))
    if y+1 < max_y and x-1 >= 0:
        res.add((x-1, y+1))
    if x-1 >= 0 and y-1 >= 0:
        res.add((x-1, y-1))
    if y-1 >= 0 and x+1 < max_x:
        res.add((x+1, y-1))
    return res

def sol1(data, n):
    res = 0
    max_x = -1
    max_y = -1
    for k in data.keys():
        max_x = max(max_x, k[0]+1)
        max_y = max(max_y, k[1]+1)
    for i in range(n):
        for k in data.keys():
            data[k] += 1

        to_check = [k for k in data.keys() if data[k] == 10]
        # print(f"{to_check = }")
        while len(to_check) > 0:
            for k in to_check:
                # print_grid(data)
                data[k] = 0
                near = compute_near(k, max_x, max_y)
                for p_n in near:
                    if data[p_n] != 0:
                        data[p_n] = min(data[p_n]+1, 10)
            to_check = [k for k in data.keys() if data[k] == 10]

        for k in data.keys():
            if data[k] == 0:
                res += 1
        # print_grid(data)
        # input()
    return res

def sol2(data):
    res = 0
    max_x = -1
    max_y = -1
    for k in data.keys():
        max_x = max(max_x, k[0]+1)
        max_y = max(max_y, k[1]+1)
    i = 0
    all_sync = False
    while not all_sync:
        for k in data.keys():
            data[k] += 1

        to_check = [k for k in data.keys() if data[k] == 10]
        # print(f"{to_check = }")
        while len(to_check) > 0:
            for k in to_check:
                # print_grid(data)
                data[k] = 0
                near = compute_near(k, max_x, max_y)
                for p_n in near:
                    if data[p_n] != 0:
                        data[p_n] = min(data[p_n]+1, 10)
            to_check = [k for k in data.keys() if data[k] == 10]

        for k in data.keys():
            if data[k] == 0:
                res += 1
        # print_grid(data)
        # input()
        if res == max_x * max_y:
            all_sync = True
        res = 0
        i += 1
    return i

def print_grid(data):
    max_x = -1
    max_y = -1
    for k in data.keys():
        max_x = max(max_x, k[0]+1)
        max_y = max(max_y, k[1]+1)
    grid = [[-1 for _ in range(max_x)] for _ in range(max_y)]
    for x in range(max_x):
        for y in range(max_y):
            grid[y][x] = data[(x, y)]
    print("Grid:")
    print("\n".join(["".join([str(e) for e in l]) for l in grid]))


grid = [[int(x) for x in e.rstrip()] for e in f.read().rstrip().split("\n")]
data = {(x, y): v for y, line in enumerate(grid) for x, v in enumerate(line)}
data_bis = {(x, y): v for y, line in enumerate(grid) for x, v in enumerate(line)}
# print(data)

print(f"Q1: {sol1(data, 100)}")
print(f"Q2: {sol2(data_bis)}")
