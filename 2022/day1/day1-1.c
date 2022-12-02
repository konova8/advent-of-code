#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
	int max = 0;
	int sum = 0;
	FILE *stream = fopen(argv[1], "r");
    char *line = NULL;
	size_t len = 0;
	while (getline(&line, &len, stream) != -1) {
		//printf("%d\n", atoi(line));
		sum += atoi(line);
		if (atoi(line) == 0) {
			if (sum > max) {
				max = sum;
			}
			sum = 0;
		}
	}
	printf("\n\n %d \n\n", max);
	fclose(stream);
}

