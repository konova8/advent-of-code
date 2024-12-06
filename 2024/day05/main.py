import numpy as np

FILE_NAME = "input.txt"
f = open(FILE_NAME, "r")

text = f.read().rstrip()
tmp = text.split("\n\n")

rules = [[int(i) for i in e.rstrip().split("|")] for e in tmp[0].split("\n")]
updates = [[int(i) for i in e.rstrip().split(",")] for e in tmp[1].split("\n")]

sol_1 = 0
sol_2 = 0

for i, u in enumerate(updates):
    good = True
    for r in rules:
        if not good:
            continue
        if (r[0] in u and r[1] in u):
            if not u.index(r[0]) < u.index(r[1]):
                good = False
    if good:
        sol_1 += u[len(u)//2]
    else:
        relevant_rules = list(filter(lambda x: x[0] in u and x[1] in u, rules))
        m = {}
        for e in u:
            m[e] = 0
        for r in relevant_rules:
            m[r[0]] += 1
        l = [[k, m[k]] for k in m.keys()]
        l.sort(key=lambda x: x[1], reverse=True)
        l = [e[0] for e in l]
        sol_2 += l[len(l)//2]

print("Q1: ", sol_1)
print("Q2: ", sol_2)
