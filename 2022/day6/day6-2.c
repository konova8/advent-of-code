#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/time.h>

#define LEN_MARKER 14

int main(int argc, char *argv[])
{
    struct timeval tval_before, tval_after, tval_result;

    gettimeofday(&tval_before, NULL);

    //
    //
    //

    int res = 0;
    FILE *stream = fopen(argv[1], "r");
    char *line = NULL;
    size_t len = 0;

    int possible = 1;

    while (getline(&line, &len, stream) != -1)
    {
        int length = strlen(line);
        //printf("line: %s\nlen line: %d\nlength line: %d\n", line, len, length);
        for (int i = LEN_MARKER - 1; i < length; i++) // Controllo tutti i possibili 14-etti
        {
            possible = 1;
            //printf("%c%c%c%c\n", line[i-3], line[i-2], line[i-1], line[i]);
            for (int j = 0; j < LEN_MARKER && possible; j++) // 
            {
                for (int k = 1; k < LEN_MARKER - j && possible; k++)
                {
                    //printf("Confronto: %c - %c\t", line[i-j], line[i-j-k]);
                    if (line[i-j] == line[i-j-k])
                    {
                        //printf("Uguali\n");
                        possible = 0;
                    }
                    else
                    {
                        //printf("Diversi\n");
                    }
                }
            }

            if (possible == 1)
            {
                res = i+1;
                break;
            }
        }
    }
    printf("%d\n", res);
    fclose(stream);

    //
    //
    //

    gettimeofday(&tval_after, NULL);

    timersub(&tval_after, &tval_before, &tval_result);

    printf("Time elapsed: %ld.%06ld seconds\n", (long int)tval_result.tv_sec, (long int)tval_result.tv_usec);
}
