def abs(a):
    if a < 0:
        return -a
    return a

f = open("input.txt", "r")
lines = f.readlines()
left = []
right = []
right_map = {}
for l in lines:
    l = l.rstrip().split(" ")
    # print(l)
    left.append(int(l[0]))
    right.append(int(l[-1]))
    right_map[int(l[-1])] = right_map.get(int(l[-1]), 0) + 1
left.sort()
right.sort()

count = 0
score = 0
for i in range(0, len(left)):
    count += abs(left[i] - right[i])
    score += left[i] * right_map.get(left[i], 0)

print("Q1: ", count)
print("Q2: ", score)
