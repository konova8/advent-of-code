#include <stdio.h>
#include <stdlib.h>
#include <sys/time.h>

#define LEN 3

int countZeros(int *arr)
{
	int res = 0;
	for (int i = 0; i < LEN; i++)
		if (arr[i] == 0)
			res++;
	return res;
}

int min(int *arr)
{
	int m = arr[0];
	if (countZeros(arr) > 1)
		return 0;
	for (int i = 1; i < LEN + 1; i++)
		if (m > arr[i] && arr[i] != 0)
			m = arr[i];
	return m;
}

void add(int val, int *arr)
{
	int cont = 0;
	for (int i = 0; i < LEN + 1; i++)
	{
		if (arr[i] == 0)
		{
			if (cont == 0)
				arr[i] = val;
			cont++;
		}
	}

	if (cont != 1)
		return;

	int m = min(arr);

	for (int i = 0; i < LEN + 1; i++)
	{
		if (arr[i] == m)
		{
			arr[i] = 0;
			break;
		}
	}
}

int main(int argc, char *argv[])
{
	struct timeval tval_before, tval_after, tval_result;

	gettimeofday(&tval_before, NULL);

	//
	//
	//

	int max[LEN + 1];
	for (int i = 0; i < LEN + 1; i++)
	{
		max[i] = 0;
	}

	int sum = 0;
	FILE *stream = fopen(argv[1], "r");
	char *line = NULL;
	size_t len = 0;

	while (getline(&line, &len, stream) != -1)
	{
		sum += atoi(line);
		if (atoi(line) == 0)
		{
			add(sum, max);
			sum = 0;
		}
	}
	printf("%d\n", max[0] + max[1] + max[2] + max[3]);

	fclose(stream);

	//
	//
	//

	gettimeofday(&tval_after, NULL);

	timersub(&tval_after, &tval_before, &tval_result);

	printf("Time elapsed: %ld.%06ld seconds\n", (long int)tval_result.tv_sec, (long int)tval_result.tv_usec);
}
