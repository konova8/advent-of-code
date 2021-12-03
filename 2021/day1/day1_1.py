f = open("input")
l = f.read().splitlines()
f.close()

cont = 0
for i in range(1, len(l)):
    if int(l[i-1]) <= int(l[i]):
        cont += 1

print(cont)
