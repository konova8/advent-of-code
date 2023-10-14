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
    int firstElf[2];
    int secondElf[2];
    char *tmpFirstElf;
    char *tmpSecondElf;
    char letter;
    while (getline(&line, &len, stream) != -1)
    {
        int length = strlen(line);
        //printf("line: %slen line: %d\n\n\n", line, length);
        
        tmpFirstElf = strtok(line, ",");
        tmpSecondElf = strtok(NULL, ",");
        
        //printf("primo elfo: %s \tsecondo elfo: %s\n\n\n", firstElf, secondElf);

        firstElf[0] = atoi(strtok(tmpFirstElf, "-"));
        firstElf[1] = atoi(strtok(NULL, "-"));

        secondElf[0] = atoi(strtok(tmpSecondElf, "-"));
        secondElf[1] = atoi(strtok(NULL, "-"));

        // Casi possibili
        // 1-3, 3-5 SI
        // 1-3, 4-5 NO
        // 3-5, 1-3 SI
        // 4-5, 1-3 NO

        //printf("First pair: %d-%d\tSecond pair: %d-%d\n", firstElf[0], firstElf[1], secondElf[0], secondElf[1]);
        if (firstElf[1] >= secondElf[0] && firstElf[0] <= secondElf[1])
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
