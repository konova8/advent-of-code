import functools
from math import floor, log10

FILE_NAME = "example.txt"
FILE_NAME = "input.txt"
f = open(FILE_NAME, "r")

sol_1: int = 0
sol_2: int = 0

def sol1(lines):
    res = [0 for _ in range(len(lines[0]))]
    for l in lines:
        for i, c in enumerate(l):
            if c == '1':
                res[i] += 1
    res.reverse()
    k = 1
    gamma = 0
    epsilon = 0
    for e in res:
        if e > len(lines)/2:
            gamma += k
        else:
            epsilon += k
        k *= 2
    return gamma * epsilon

def get_most(lines, i, is_most):
    if len(lines) == 1:
        return lines[0]
    ones = 0
    zeros = 0
    for l in lines:
        if l[i] == '1':
            ones += 1
        else:
            zeros += 1
    to_keep = '1'
    if ones >= zeros:
        if is_most:
            to_keep = '1'
        else:
            to_keep = '0'
    else:
        if is_most:
            to_keep = '0'
        else:
            to_keep = '1'
    lines = list(filter(lambda x: x[i] == to_keep, lines))
    return get_most(lines, i+1, is_most)

def sol2(lines):
    most = get_most(lines, 0, True)
    less = get_most(lines, 0, False)
    return int(most, 2) * int(less, 2)


data = [e.rstrip() for e in f.read().rstrip().split("\n")]
print(f"Q1: {sol1(data)}")
print(f"Q2: {sol2(data)}")
