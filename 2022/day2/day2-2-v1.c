#include <stdio.h>
#include <stdlib.h>
#include <sys/time.h>

int main(int argc, char *argv[]) {
    struct timeval tval_before, tval_after, tval_result;

	gettimeofday(&tval_before, NULL);

	//
	//
	//

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
            case 'X': // I need to lose
                total += 3;
                total += 0;
                break;
            case 'Y': // I need to draw
                total += 1;
                total += 3;
                break;
            case 'Z': // I need to win
                total += 2;
                total += 6;
                break;
            }
            break;

        case 'B': // Opponent chose Paper
            switch (letter2)
            {
            case 'X': // I need to lose
                total += 1;
                total += 0;
                break;
            case 'Y': // I need to draw
                total += 2;
                total += 3;
                break;
            case 'Z': // I need to win
                total += 3;
                total += 6;
                break;
            }
            break;

        case 'C': // Opponent chose Scissors
            switch (letter2)
            {
            case 'X': // I need to lose
                total += 2;
                total += 0;
                break;
            case 'Y': // I need to draw
                total += 3;
                total += 3;
                break;
            case 'Z': // I need to win
                total += 1;
                total += 6;
                break;
            }
            break;
        
        default:
            break;
        }
	}
	printf("\n\n %d \n\n", total);
	fclose(stream);

	//
	//
	//

	gettimeofday(&tval_after, NULL);

	timersub(&tval_after, &tval_before, &tval_result);

	printf("Time elapsed: %ld.%06ld seconds\n", (long int)tval_result.tv_sec, (long int)tval_result.tv_usec);
}

