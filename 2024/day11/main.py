import functools
from math import floor, log10

FILE_NAME = "example.txt"
FILE_NAME = "input.txt"
f = open(FILE_NAME, "r")

sol_1: int = 0
sol_2: int = 0

@functools.cache
def count(x, d):
    if d == 0: return 1
    if x == 0: return count(1, d-1)

    l = floor(log10(x))+1
    if l % 2: return count(x*2024, d-1)

    return count(x // 10**(l//2), d-1) + count(x %  10**(l//2), d-1)

data = [int(e) for e in f.read().rstrip().split(" ")]
print(f"Q1: {sum(map(count, data, [25 for _ in data]))}")
print(f"Q2: {sum(map(count, data, [75 for _ in data]))}")
