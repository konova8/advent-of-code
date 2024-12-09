from copy import deepcopy

FILE_NAME = "example.txt"
FILE_NAME = "input.txt"
f = open(FILE_NAME, "r")

line = f.read().rstrip()
# print(line)
is_free = False
l = []
l_bis = []
i = 0
blank_count: int = 0
for c in line:
    # print(c)
    if is_free:
        l += [-1] * int(c)
        blank_count += int(c)
    else:
        l += [i] * int(c)
        i += 1
    is_free = not is_free

# print(l)
l_bis = deepcopy(l)
first_free = l.index(-1)

while blank_count > 0 and first_free < len(l):
    # print()
    # print(l)
    # print(len(l))
    # print(first_free)
    # print(blank_count)
    # print()
    l[first_free] = l[-1]
    l.pop()
    blank_count -= 1
    try:
        first_free = l.index(-1, first_free)
    except ValueError:
        continue

# print(l)

def count(n: int, l: list[int], reverse: bool) -> int:
    res = 0
    for i in range(0, len(l)):
        e = 0
        if not reverse:
            e = l[i]
        else:
            e = l[-(i+1)]
        # print(f"{l = }")
        # print(f"{e = }")
        if e == n:
            res += 1
        elif e != n and res == 0:
            continue
        else:
            break
    return res

def find_sub_list(sl,l):
    sll=len(sl)
    for ind in (i for i,e in enumerate(l) if e==sl[0]):
        if l[ind:ind+sll]==sl:
            return ind,ind+sll-1
    return -1, -1


different_number_count = l_bis[-1]+1
# print(f"{list(range(different_number_count-1, -1, -1)) = }")
for e in range(different_number_count-1, -1, -1):
    # print(f"Before: {l_bis}")
    if e % 100 == 0:
        print("Needs to get to 0")
        print(f"{e = }")
    card = count(e, l_bis, True)
    # print(f"{card = }")
    a, b = find_sub_list([-1]*card, l_bis)
    # print(f"{a = }")
    # print(f"{b = }")
    # input()
    if a == -1 and b == -1 or a > l_bis.index(e):
        continue
    # add condition to move only if element is placed better
    for i, x in enumerate(l_bis):
        if x == e:
            l_bis[i] = -1
    l_bis[a:b+1] = [e]*card
    while l_bis[-1] == -1:
        l_bis = l_bis[:-1]
    # print(f"After : {l_bis}")
    # input()


sol_1: int = 0
sol_2: int = 0

for i, e in enumerate(l):
    sol_1 += e * i
for i, e in enumerate(l_bis):
    if e != -1:
        sol_2 += e * i

print("Q1: ", sol_1)
print("Q2: ", sol_2)
