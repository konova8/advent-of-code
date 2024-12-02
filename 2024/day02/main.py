def abs(a):
    if a < 0:
        return -a
    return a

def is_safe(l: list[int]):
    return is_inc(l) or is_dec(l)

def is_inc(l: list[int]):
    if len(l) == 1:
        return True
    diff = int(l[0]) - int(l[1])
    if diff < 1 or diff > 3:
        return False
    return is_inc(l[1:])

def is_dec(l: list[int]):
    if len(l) == 1:
        return True
    diff = int(l[1]) - int(l[0])
    if diff < 1 or diff > 3:
        return False
    return is_dec(l[1:])


def is_safe_with_one_bad_level(l: list[int]):
    if is_dec(l) or is_inc(l):
        return True
    for i in range(0, len(l)):
        new_l = l[:i] + l[i+1:]
        if is_dec(new_l) or is_inc(new_l):
            return True
    return False

f = open("input.txt", "r")
lines = f.readlines()

count = 0
count_bis = 0
for l in lines:
    l = l.rstrip().split(" ")
    l_int = []
    for s in l:
        l_int.append(int(s))
    if is_safe(l_int):
        count += 1
    if is_safe_with_one_bad_level(l_int):
        count_bis += 1

print("Q1: ", count)
print("Q2: ", count_bis)
