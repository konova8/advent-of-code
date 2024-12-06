import re

def abs(a):
    if a < 0:
        return -a
    return a

f = open("input.txt", "r")
lines = f.readlines()

parings_bis = []
parings_bis += [('do()', '', '')]
for l in lines:
    parings_bis += re.findall(r"(mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\))", l)
# print(parings_bis)


sol_1 = 0
sol_2 = 0
mul = 1
for t in parings_bis:
    if (t[0] == "do()"):
        mul = 1
    elif (t[0] == "don't()"):
        mul = 0
    else:
        sol_1 += int(t[1]) * int(t[2])
        sol_2 += mul * int(t[1]) * int(t[2])

print("Q1: ", sol_1)
print("Q2: ", sol_2)
