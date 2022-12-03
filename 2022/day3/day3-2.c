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
    char *thiPart;
    char letter;
    int found = 0;
    int k = 0;
    while (getline(&line, &len, stream) != -1)
    {
        if (k % 3 == 0)
            firPart = strdup(line);
        if (k % 3 == 1)
            secPart = strdup(line);
        if (k % 3 == 2)
            thiPart = strdup(line);

        // Trova la lettera in comune tra le tre parti
        if (k % 3 == 2)
        {
            //printf("first part: %ssecond part: %sthird part: %s\n", firPart, secPart, thiPart);
            for (int i = 0; i < strlen(firPart) && !found; i++)
            {
                for (int j = 0; j < strlen(secPart) && !found; j++)
                {
                    for (int h = 0; h < strlen(thiPart) && !found; h++)
                    {
                        if (firPart[i] == secPart[j] && secPart[j] == thiPart[h])
                        {
                            letter = firPart[i];
                            found = 1;
                        }
                    }
                }
            }
            //printf("%c\n", letter);
            if (letter >= 'A' && letter <= 'Z')
                sum += letter - 38;
            else if (letter >= 'a' && letter <= 'z')
                sum += letter - 96;
            found = 0;
        }

        k++;
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
