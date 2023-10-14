f = open("input.txt")
l = f.read().splitlines()
f.close()
listaElfi = []

sum = 0
for i in range(1, len(l)):
    if l[i] != "":
        sum += int(l[i])
    else:
        listaElfi.append(sum)
        sum = 0

max1 = max(listaElfi)
listaElfi.remove(max1)

max2 = max(listaElfi)
listaElfi.remove(max2)

max3 = max(listaElfi)
listaElfi.remove(max3)

print(max1 + max2 + max3)