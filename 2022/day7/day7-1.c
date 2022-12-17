#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/time.h>

#define LINE_MAX 100

int main(int argc, char *argv[])
{
    struct timeval tval_before, tval_after, tval_result;

    gettimeofday(&tval_before, NULL);

    //
    //
    //

    int res = 0;

    char *firstPart = NULL;
    char *secondPart = NULL;

    FILE *fp = fopen(argv[1], "r");
    char line[LINE_MAX];
    while (fgets(line, sizeof line, fp) != NULL)
    {
        int length = strlen(line);
        printf("line: %s", line);

        firstPart = strtok(line, " ");
        secondPart = strtok(line, " ");

        printf("firstPart: %s, secondPart: %s\n", line, firstPart, secondPart);
        if (firstPart[0] != '$' && firstPart[0] != 'd')
        {
            res += atoi(firstPart);
        }
    }
    printf("%d\n", res);
    fclose(fp);

    //
    //
    //

    gettimeofday(&tval_after, NULL);

    timersub(&tval_after, &tval_before, &tval_result);

    printf("Time elapsed: %ld.%06ld seconds\n", (long int)tval_result.tv_sec, (long int)tval_result.tv_usec);
}
