import re

f = open("./input.txt")
l = f.read().splitlines()
f.close()
situazioneIniziale = l[0:8]
istruzioni = l[10:]

#print(istruzioni)
#print(situazioneIniziale)

# Inserisco la situazione iniziale in un array di liste
pile = [[], [], [], [], [], [], [], [], []]

# 1 5 9 13 17 ...
# (i-1) / 4
for row in situazioneIniziale:
    for i in range(0, len(row)):
        if (row[i] != ' ' and row[i] != '[' and row[i] != ']'):
            pile[(i-1)//4].append(row[i])
            #print(row[i])

for pila in pile:
    print(pila)

print('---')

# faccio le mosse
for row in istruzioni:
    #print(row)
    res = re.match(r"move (?P<one>[0-9]+) from (?P<two>[0-9]+) to (?P<three>[0-9]+)", row)

    nrToMove, fromNr, toNr = int(res['one']) - 1, int(res['two']) - 1, int(res['three']) - 1
    
    for i in range(0, nrToMove+1):
        tmp = pile[fromNr].pop(0)
        pile[toNr].insert(0, tmp)
    
    #for i in range(len(pile)):
    #    print(i+1, ' - ', pile[i])
    
#for i in range(len(pile)):
#    print(i, ' - ', pile[i])

for pila in pile:
    print(pila)

for pila in pile:
    print(pila[0], end='')
print()