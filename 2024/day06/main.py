FILE_NAME = "example.txt"
FILE_NAME = "input.txt"
f = open(FILE_NAME, "r")

lines: list[str] = [l.rstrip() for l in f.readlines()]

sol_1: int = 0
sol_2: int = 0

ROCKS: set[complex] = set()
starting_position: complex = 0+0j
marked: set[complex] = set()

for i, l in enumerate(lines[::-1]):
    for j, c in enumerate(l):
        if c == "^":
            me = complex(i, j)
            starting_position = complex(i, j)
        elif c == "#":
            ROCKS.add(complex(i, j))
me: complex = complex(starting_position.real, starting_position.imag)

BOUND_REAL: int = max([int(e.real) for e in ROCKS])
BOUND_IMAG: int = max([int(e.imag) for e in ROCKS])
DIRECTIONS: list[complex] = [
    1+0j, # Going up
    0+1j, # Going right
    -1+0j, # Going down
    0-1j, # Going left
]
i = 0

while (me.real >=0 and me.real <= BOUND_REAL and me.imag >= 0 and me.imag <= BOUND_IMAG):
    if me in ROCKS:
        me -= DIRECTIONS[i]
        i = (i+1) % len(DIRECTIONS)
    else:
        marked.add(complex(me.real, me.imag))
        me += DIRECTIONS[i]

sol_1 = len(marked)

done: int = 0
for c in marked:
    done += 1
    me: complex = complex(starting_position.real, starting_position.imag)
    O = complex(c.real, c.imag)
    if O == me:
        continue
    print(f"Done {done} out of {BOUND_REAL*BOUND_IMAG} total cases, {(done / (BOUND_REAL*BOUND_IMAG)) * 100:.2f}%")
    steps: set[tuple[int, complex]] = set()
    i = 0
    in_loop = False
    while (me.real >=0 and me.real <= BOUND_REAL and me.imag >= 0 and me.imag <= BOUND_IMAG) and not in_loop:
        current = (i, complex(me.real, me.imag))
        if current in steps:
            sol_2 += 1
            in_loop = True
            continue
        if me in ROCKS or me == O:
            me -= DIRECTIONS[i]
            i = (i+1) % len(DIRECTIONS)
        else:
            steps.add((i, complex(me.real, me.imag)))
            me += DIRECTIONS[i]

print("Q1: ", sol_1)
print("Q2: ", sol_2)
