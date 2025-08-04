FILE_NAME = "example.txt"
FILE_NAME = "input.txt"
f = open(FILE_NAME, "r")

sol_1: int = 0
sol_2: int = 0

def sol(data, n):
    local_data = {e: data.count(e) for e in range(9)}

    for i in range(n):
        new_local_data = {}
        new_local_data[0] = local_data[1]
        new_local_data[1] = local_data[2]
        new_local_data[2] = local_data[3]
        new_local_data[3] = local_data[4]
        new_local_data[4] = local_data[5]
        new_local_data[5] = local_data[6]
        new_local_data[6] = local_data[7] + local_data[0]
        new_local_data[7] = local_data[8]
        new_local_data[8] = local_data[0]
        local_data = new_local_data
        # print(f"--- Day {i+1} ---")
        # print(local_data, len(local_data))
    return sum([local_data[k] for k in local_data.keys()])

data = [int(e.rstrip()) for e in f.read().rstrip().split(",")]

sol_1 = sol(data, 80)
sol_2 = sol(data, 256)

print(f"Q1: {sol_1}")
print(f"Q2: {sol_2}")
