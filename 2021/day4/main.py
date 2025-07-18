import functools
from math import floor, log10

FILE_NAME = "example.txt"
FILE_NAME = "input.txt"
f = open(FILE_NAME, "r")

sol_1: int = 0
sol_2: int = 0

def sol1(order, boards):
    for n in order:
        n = int(n, 10)
        for b, board in enumerate(boards):
            for y, row in enumerate(board):
                for x, e in enumerate(row):
                    if e == n:
                        boards[b][y][x] = -1
        # Check if board is good
        # Rows
        for b, board in enumerate(boards):
            for y, row in enumerate(board):
                if all([e == -1 for e in row]):
                    return (b, n)
        # Cols
        for b, board in enumerate(boards):
            for c_index in [0, 1, 2, 3, 4]:
                if all([row[c_index] == -1 for row in board]):
                    return (b, n)
    return (-1, -1)

def sol2(order, boards):
    good_boards = [1 for _ in range(len(boards))]
    b_last = -1
    for n in order:
        n = int(n, 10)
        if sum(good_boards) == 1:
            b_last = good_boards.index(1)
        for b, board in enumerate(boards):
            for y, row in enumerate(board):
                for x, e in enumerate(row):
                    if e == n:
                        boards[b][y][x] = -1
        # Check if board is good
        # Rows
        for b, board in enumerate(boards):
            for y, row in enumerate(board):
                if all([e == -1 for e in row]):
                    good_boards[b] = 0
        # Cols
        for b, board in enumerate(boards):
            for c_index in [0, 1, 2, 3, 4]:
                if all([row[c_index] == -1 for row in board]):
                    good_boards[b] = 0
        if sum(good_boards) == 0:
            return (b_last, n)
    return (-1, -1)


data = [e.rstrip() for e in f.read().rstrip().split("\n\n")]
order = data[0].rstrip().split(',')
boards = [[[int(e, 10) for e in l.split(" ") if len(e)>0] for l in b.split("\n")] for b in data[1:]]

(i, n) = sol1(order, boards)
for row in boards[i]:
    for e in row:
        if e != -1:
            sol_1 += e
sol_1 *= n

print(f"Q1: {sol_1}")


(i, n) = sol2(order, boards)

for row in boards[i]:
    for e in row:
        if e != -1:
            sol_2 += e
sol_2 *= n
print(f"Q2: {sol_2}")
