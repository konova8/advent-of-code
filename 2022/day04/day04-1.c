#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/time.h>

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
    int firstPair[2];
    int secondPair[2];
    char *firstElf;
    char *secondElf;
    char letter;
    while (getline(&line, &len, stream) != -1)
    {
        int length = strlen(line);
        //printf("line: %slen line: %d\n\n\n", line, length);
        
        firstElf = strtok(line, ",");
        secondElf = strtok(NULL, ",");
        
        //printf("primo elfo: %s \tsecondo elfo: %s\n\n\n", firstElf, secondElf);

        firstPair[0] = atoi(strtok(firstElf, "-"));
        firstPair[1] = atoi(strtok(NULL, "-"));

        secondPair[0] = atoi(strtok(secondElf, "-"));
        secondPair[1] = atoi(strtok(NULL, "-"));

        //printf("first pair: %d \t %d\nsecond pair: %d \t %d\n\n\n", firstPair[0], firstPair[1], secondPair[0], secondPair[1]);

        if ((firstPair[0] <= secondPair[0] && secondPair[1] <=  firstPair[1]) || (secondPair[0] <= firstPair[0] && firstPair[1] <=  secondPair[1]))
            sum++;
    }
    printf("%d\n", sum);
    fclose(stream);

    //
    //
    //

    gettimeofday(&tval_after, NULL);

    timersub(&tval_after, &tval_before, &tval_result);

    printf("Time elapsed: %ld.%06ld seconds\n", (long int)tval_result.tv_sec, (long int)tval_result.tv_usec);
}
