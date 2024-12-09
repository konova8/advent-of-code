from copy import deepcopy

FILE_NAME = "example.txt"
FILE_NAME = "input.txt"
f = open(FILE_NAME, "r")
sol_1: int = 0
sol_2: int = 0

lines = [l.rstrip() for l in f.readlines()]

max_i = len(lines)
max_j = len(lines[0])

D: dict[str, set[complex]] = {}
marked_1 = set()
marked_2 = set()
for l in lines:
    for e in l:
        if e != ".":
            D[e] = set()
for i, l in enumerate(lines):
    for j, e in enumerate(l):
        if e == ".":
            continue
        # print(l)
        # print(e)
        # print(f"{i = }, {j = }")
        # input()

        x = complex(i, j)
        # Handle new point, get pairs of antinodes for this and all the other antennas with the same key
        for k in D[e]:
            diff_real = abs(k.real - x.real)
            diff_imag = abs(k.imag - x.imag)
            if k.real >= x.real and k.imag >= x.imag:
                a = complex(k.real + diff_real, k.imag + diff_imag)
                if a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_1.add(a)
                while a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_2.add(a)
                    a = complex(a.real + diff_real, a.imag + diff_imag)
                a = complex(x.real - diff_real, x.imag - diff_imag)
                if a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_1.add(a)
                while a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_2.add(a)
                    a = complex(a.real - diff_real, a.imag - diff_imag)
            elif k.real <= x.real and k.imag <= x.imag:
                a = complex(k.real - diff_real, k.imag - diff_imag)
                if a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_1.add(a)
                while a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_2.add(a)
                    a = complex(a.real - diff_real, a.imag - diff_imag)
                a = complex(x.real + diff_real, x.imag + diff_imag)
                if a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_1.add(a)
                while a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_2.add(a)
                    a = complex(a.real + diff_real, a.imag + diff_imag)
            elif k.real <= x.real and k.imag >= x.imag:
                a = complex(k.real - diff_real, k.imag + diff_imag)
                if a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_1.add(a)
                while a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_2.add(a)
                    a = complex(a.real - diff_real, a.imag + diff_imag)
                a = complex(x.real + diff_real, x.imag - diff_imag)
                if a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_1.add(a)
                while a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_2.add(a)
                    a = complex(a.real + diff_real, a.imag - diff_imag)
            elif k.real >= x.real and k.imag <= x.imag:
                a = complex(k.real + diff_real, k.imag - diff_imag)
                if a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_1.add(a)
                while a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_2.add(a)
                    a = complex(a.real + diff_real, a.imag - diff_imag)
                a = complex(x.real - diff_real, x.imag + diff_imag)
                if a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_1.add(a)
                while a.real >= 0 and a.imag >= 0 and a.real < max_j and a.imag < max_i:
                    marked_2.add(a)
                    a = complex(a.real - diff_real, a.imag + diff_imag)
        D[e].add(complex(i, j))
        marked_2.add(complex(i, j))

sol_1 = len(marked_1)
sol_2 = len(marked_2)

print("Q1: ", sol_1)
print("Q2: ", sol_2)
