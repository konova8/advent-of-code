f = open("input").read().splitlines()
l = []

moves_f = 0
moves_d = 0

print(f)

for e in f:
    if e[0] == "f":
        moves_f += int(e[8:])
    elif e[0] == "d":
        moves_d += int(e[5:])
    elif e[0] == "u":
        moves_d -= int(e[3:])

print(moves_f, moves_d, "\t", moves_f * moves_d)