FILE_NAME = "example.txt"
FILE_NAME = "input.txt"
f = open(FILE_NAME, "r")

sol_1: int = 0
sol_2: int = 0

def sol1(data):
    res = {e: 0 for e in range(10)}
    for l in data:
        for e in l[1]:
            res[len(e)] = res.get(len(e), 0) + 1

    return res[2] + res[3] + res[4] + res[7]

def sol2(data):

    res = 0
    for l in data:
        schema = {}
        # 1, 4, 7, 8
        for e in l[0]:
            if len(e) == 2:
                schema[1] = e
            elif len(e) == 3:
                schema[7] = e
            elif len(e) == 4:
                schema[4] = e
            elif len(e) == 7:
                schema[8] = e
        # 3
        for e in l[0]:
            if len(e) == 5 and all([x in e for x in schema[1]]):
                schema[3] = e
        # 5, 2
        for e in l[0]:
            if len(e) == 5 and e != schema[3] and sum([x in schema[4] for x in e]) == 3:
                schema[5] = e
            if len(e) == 5 and e != schema[3] and sum([x in schema[4] for x in e]) == 2:
                schema[2] = e
        # 6, 9
        for e in l[0]:
            if len(e) == 6 and sum([x in schema[1] for x in e]) == 1:
                schema[6] = e
            if len(e) == 6 and sum([x in schema[3] for x in e]) == 5:
                schema[9] = e

        # 0
        for e in l[0]:
            if e not in schema.values():
                schema[0] = e

        if len(set(schema.values())) != 10:
            exit(1)

        good_schema = {schema[k]: k for k in schema.keys()}
        # print(good_schema)

        mul = 1
        for e in l[1][::-1]:
            res += good_schema[e]*mul
            mul *= 10

    return res

lines = [e.rstrip() for e in f.read().rstrip().split("\n")]
data = []
for l in lines:
    pair = l.split(" | ")
    first = ["".join(sorted(e)) for e in pair[0].split(" ")]
    second = ["".join(sorted(e)) for e in pair[1].split(" ")]
    data.append((first, second))

# print(data)

print(f"Q1: {sol1(data)}")
print(f"Q2: {sol2(data)}")
