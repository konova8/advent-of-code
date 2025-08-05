FILE_NAME = "example.txt"
FILE_NAME = "input.txt"
f = open(FILE_NAME, "r")

sol_1: int = 0
sol_2: int = 0

db = {
    ")": "(",
    "]": "[",
    "}": "{",
    ">": "<",
}

db_reverse = {
    "(": ")",
    "[": "]",
    "{": "}",
    "<": ">",
}

points = {
    ")": 3,
    "]": 57,
    "}": 1197,
    ">": 25137,
}

points_tail = {
    ")": 1,
    "]": 2,
    "}": 3,
    ">": 4,
}

def sol1(data):
    res = []
    for l in data:
        queue = []
        for i, c in enumerate(l):
            if c in db.values():
                queue.append(c)
            else:
                current = queue.pop()
                if current != db[c]:
                    res.append(c)
                    continue
    return sum([points[c] for c in res])

def sol2(data):
    res = []
    for l in data:
        queue = []
        is_corrupt = False
        for i, c in enumerate(l):
            if c in db.values():
                queue.append(c)
            else:
                current = queue.pop()
                if current != db[c]:
                    is_corrupt = True
        if not is_corrupt and len(queue) != 0:
            missing_tail = [db_reverse[e] for e in queue[::-1]]
            res.append(missing_tail)
    res2 = []
    for tail in res:
        score = 0
        for c in tail:
            score *= 5
            score += points_tail[c]
        res2.append(score)
    return sorted(res2)[len(res2)//2]

data = [e.rstrip() for e in f.read().rstrip().split("\n")]

print(f"Q1: {sol1(data)}")
print(f"Q2: {sol2(data)}")
