#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/time.h>

#define LEN_MARKER 4

int univoci(char a, char b, char c, char d)
{
    if (a == b)
        return 0;
    else if (a == c)
        return 0;
    else if (a == d)
        return 0;
    else if (b == c)
        return 0;
    else if (b == d)
        return 0;
    else if (c == d)
        return 0;
    else
        return 1;
    
}

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
    
    while (getline(&line, &len, stream) != -1)
    {
        int length = strlen(line);
        //printf("line: %s\nlen line: %d\nlength line: %d\n", line, len, length);
        for (int i = LEN_MARKER - 1; i < length; i++)
        {
            //printf("%c%c%c%c\n", line[i], line[i-1], line[i-2], line[i-3]);
            if (univoci(line[i], line[i-1], line[i-2], line[i-3]))
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
