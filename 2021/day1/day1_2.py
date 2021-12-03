f = open("input")
l = f.read().splitlines()
f.close()

cont = 0

for i in range(0, len(l)):
    l[i] = int(l[i])

l3 = []
for i in range(2, len(l)):
    l3.append(l[i-2] + l[i-1] + l[i])


cont = 0
for i in range(1, len(l3)):
    if l3[i-1] < l3[i]:
        cont += 1

print(cont)

# Oppure:
#cont1 = 0
#for i in range(3, len(l)):
#    if l[i-3] < l[i]:
#        cont1 += 1
#
#print(cont1)