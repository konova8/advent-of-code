#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
	int max[3] = {0, 0, 0};
	int sum = 0;
	FILE *stream = fopen(argv[1], "r");
    char *line = NULL;
	size_t len = 0;
	int min = 0;
	int posToSum = 0;
	while (getline(&line, &len, stream) != -1) {
		sum += atoi(line);
		if (atoi(line) == 0) {
			if (sum > max[0] || sum > max[1] || sum > max[2]) {
				for (int i = 0; i < 3; i++) {
					if (max[i] < min) {
						min = max[i];
					}
				}
				for (int i = 0; i < 3; i++) {
					if (max[i] = min) {
						posToSum = i;
					}
				}
				max[posToSum] = sum;
			}
			sum = 0;
		}
	}
	printf("\n\n %d \n\n", max[0] + max[1] + max[2]);
	fclose(stream);
}

