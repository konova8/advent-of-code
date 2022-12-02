#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
	int total = 0;
	FILE *stream = fopen(argv[1], "r");
    char *line = NULL;
	size_t len = 0;
	while (getline(&line, &len, stream) != -1) {
        char letter1 = line[0];
        char letter2 = line[2];
		switch (letter1)
        {
        case 'A': // Opponent chose Rock
            switch (letter2)
            {
            case 'X': // I chose Rock
                total += 1;
                total += 3;
                break;
            case 'Y': // I chose Paper
                total += 2;
                total += 6;
                break;
            case 'Z': // I chose Scissors
                total += 3;
                total += 0;
                break;
            }
            break;

        case 'B': // Opponent chose Paper
            switch (letter2)
            {
            case 'X': // I chose Rock
                total += 1;
                total += 0;
                break;
            case 'Y': // I chose Paper
                total += 2;
                total += 3;
                break;
            case 'Z': // I chose Scissors
                total += 3;
                total += 6;
                break;
            }
            break;

        case 'C': // Opponent chose Scissors
            switch (letter2)
            {
            case 'X': // I chose Rock
                total += 1;
                total += 6;
                break;
            case 'Y': // I chose Paper
                total += 2;
                total += 0;
                break;
            case 'Z': // I chose Scissors
                total += 3;
                total += 3;
                break;
            }
            break;
        
        default:
            break;
        }
	}
	printf("\n\n %d \n\n", total);
	fclose(stream);
}

