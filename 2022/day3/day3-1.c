#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/time.h>

int esisteIn(char c, char *word)
{
    for (int i = 0; i < strlen(word); i++)
        if (c == word[i])
            return 1;
    return 0;
}


int main(int argc, char *argv[])
{
    struct timeval tval_before, tval_after, tval_result;

    gettimeofday(&tval_before, NULL);

    //
    //
    //

    int sum = 0;
    FILE *stream = fopen(argv[1], "r");
    char *line = NULL;
    size_t len = 0;
    char *firPart;
    char *secPart;
    char letter;
    while (getline(&line, &len, stream) != -1)
    {
        int length = strlen(line);
        //strncpy(firPart, line, length / 2);
        //strcpy(secPart, line + length / 2);
        firPart = strndup(line, length / 2);
        secPart = strdup(line + length / 2);
        //printf("line: %slen line: %d\nfirst part: %ssecond part: %s\n", line, length/2, firPart, secPart);

        // Trova la lettera in comune tra le due parti
        for (int i = 0; i < strlen(firPart); i++)
        {
            if (esisteIn(firPart[i], secPart))
            {
                letter = firPart[i];
                break;
            }
        }

        if (letter >= 'A' && letter <= 'Z')
            sum += letter - 38;
        else if (letter >= 'a' && letter <= 'z')
            sum += letter - 96;
    }
    printf("\n\n %d \n\n", sum);
    fclose(stream);

    //
    //
    //

    gettimeofday(&tval_after, NULL);

    timersub(&tval_after, &tval_before, &tval_result);

    printf("Time elapsed: %ld.%06ld seconds\n", (long int)tval_result.tv_sec, (long int)tval_result.tv_usec);
}
