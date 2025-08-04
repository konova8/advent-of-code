FILE_NAME = "example.txt"
FILE_NAME = "input.txt"
f = open(FILE_NAME, "r")

sol_1: int = 0
sol_2: int = 0

def sol1(data):
    board = {}
    for pair in data:
        start = pair[0]
        end = pair[1]
        if start[0] != end[0] and start[1] != end[1]:
            continue
        from_x = min(start[0], end[0])
        to_x = max(start[0], end[0]) + 1
        from_y = min(start[1], end[1])
        to_y = max(start[1], end[1]) + 1
        for x in range(from_x, to_x):
            for y in range(from_y, to_y):
                board[(x, y)] = board.get((x, y), 0) + 1

    count = 0
    for k in board.keys():
        if board[k] > 1:
            count += 1
    return count

def sol2(data):
    board = {}
    for pair in data:
        start = pair[0]
        end = pair[1]
        x_list = []
        y_list = []
        max_len = max(abs(start[0] - end[0]), abs(start[1] - end[1])) + 1

        if start[0] > end[0]:
            x_list = [e for e in range(start[0], end[0] - 1, -1)]
        elif start[0] < end[0]:
            x_list = [e for e in range(start[0], end[0] + 1)]
        else:
            x_list = [start[0]] * max_len

        if start[1] > end[1]:
            y_list = [e for e in range(start[1], end[1] - 1, -1)]
        elif start[1] < end[1]:
            y_list = [e for e in range(start[1], end[1] + 1)]
        else:
            y_list = [start[1]] * max_len

        for i in range(0, len(x_list)):
            x = x_list[i]
            y = y_list[i]
            board[(x, y)] = board.get((x, y), 0) + 1

    count = 0
    for k in board.keys():
        if board[k] > 1:
            count += 1
    return count


data = []

for line in f.read().rstrip().split("\n"):
    pairs = line.rstrip().split(" -> ")
    first = tuple(map(int, pairs[0].split(",")))
    second = tuple(map(int, pairs[1].split(",")))
    data.append((first, second))

sol_1 = sol1(data)
sol_2 = sol2(data)

print(f"Q1: {sol_1}")
print(f"Q2: {sol_2}")

