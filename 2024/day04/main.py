import numpy as np

XMAS = [c for c in 'XMAS']
SAMX = [c for c in 'SAMX']
MAS = [c for c in 'MAS']
SAM = [c for c in 'SAM']

f = open("input.txt", "r")

matrix = np.array([[c for c in l.rstrip()] for l in f.readlines()])
N = len(matrix)
sol_1 = 0
sol_2 = 0

for c in range(0, N):
    for r in range(0, N):
        # Horizontal
        if r < N and c < N-3:
            main_diag = matrix[r, c:c+4]
            main_diag = [str(e) for e in main_diag]
            if np.array_equal(main_diag, XMAS) or np.array_equal(main_diag, SAMX):
                sol_1 += 1
        # Vertical
        if r < N-3 and c < N:
            main_diag = matrix[r:r+4, c]
            main_diag = [str(e) for e in main_diag]
            if np.array_equal(main_diag, XMAS) or np.array_equal(main_diag, SAMX):
                sol_1 += 1
        # Main Diagonal
        if r < N-3 and c < N-3:
            main_diag = matrix[r:r+4, c:c+4].diagonal()
            main_diag = [str(e) for e in main_diag]
            if np.array_equal(main_diag, XMAS) or np.array_equal(main_diag, SAMX):
                sol_1 += 1
        # Second Diagonal
        if r < N-3 and c < N-3:
            main_diag = np.fliplr(matrix[r:r+4, c:c+4]).diagonal()
            main_diag = [str(e) for e in main_diag]
            if np.array_equal(main_diag, XMAS) or np.array_equal(main_diag, SAMX):
                sol_1 += 1

        # Second part
        if r > 0 and r < N-1 and c > 0 and c < N-1 and matrix[r, c] == 'A':
            square = matrix[r-1:r+2, c-1:c+2]
            main_diag = square.diagonal()
            second_diag = np.fliplr(square).diagonal()
            if (np.array_equal(main_diag, MAS) or np.array_equal(main_diag, SAM)) and (np.array_equal(second_diag, MAS) or np.array_equal(second_diag, SAM)):
                sol_2 += 1

print("Q1: ", sol_1)
print("Q2: ", sol_2)
