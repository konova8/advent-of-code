from copy import deepcopy

FILE_NAME = "example.txt"
FILE_NAME = "input.txt"
f = open(FILE_NAME, "r")

lines: list[list[int]] = [[int(c) for c in l.rstrip()] for l in f.readlines()]

def pretty_print(lines: list[list[int]]):
    for l in lines:
        for e in l:
            print(e, end="")
        print()

# pretty_print(lines)

def inside_boundaries(map: list[list[int]], p: complex):
    max_real = len(map[0])
    max_imag = len(map)
    return p.real >= 0 and p.imag >= 0 and p.real < max_real and p.imag < max_imag

def get_value(map: list[list[int]], p: complex):
    return map[int(p.imag)][int(p.real)]


zeros: list[complex] =  []
for y, l in enumerate(lines):
    for x, e in enumerate(l):
        if e == 0:
            p = complex(x, y)
            zeros.append(p)

sol_1: int = 0
sol_2: int = 0

for s in zeros:
    nines: set[complex] = set()
    seen: set[complex] = set()
    todo: list[complex] = []
    todo.append(s)
    # print(f"Starting with {s}")
    while len(todo) > 0:
        current = todo.pop(0)
        seen.add(current)
        if get_value(lines, current) == 9:
            nines.add(current)
            sol_2 += 1
            continue
        # UP
        x = complex(current.real, current.imag+1)
        if inside_boundaries(lines, x) and get_value(lines, x) == get_value(lines, current) + 1 and x not in seen:
            todo.append(x)
        # DOWN
        x = complex(current.real, current.imag-1)
        if inside_boundaries(lines, x) and get_value(lines, x) == get_value(lines, current) + 1 and x not in seen:
            todo.append(x)
        # RIGHT
        x = complex(current.real+1, current.imag)
        if inside_boundaries(lines, x) and get_value(lines, x) == get_value(lines, current) + 1 and x not in seen:
            todo.append(x)
        # LEFT
        x = complex(current.real-1, current.imag)
        if inside_boundaries(lines, x) and get_value(lines, x) == get_value(lines, current) + 1 and x not in seen:
            todo.append(x)
    # print(f"{len(nines) = }")
    # print()
    sol_1 += len(nines)

print("Q1: ", sol_1)
print("Q2: ", sol_2)
