#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/time.h>

struct istruzioni
{
    int numToMove;
    int fromNr;
    int toNr;
};


int main(int argc, char *argv[])
{
    struct timeval tval_before, tval_after, tval_result;

    gettimeofday(&tval_before, NULL);

    //
    //
    //

    FILE *stream = fopen(argv[1], "r");
    char *line = NULL;
    size_t len = 0;

    int sum = 0;
    struct Node* pile[9] = (struct Node*) malloc(sizeof(struct Node) * 9);
    

    for (int i = 0; i < 8; i++) // Ciclo per le righe
    {
        for (int j = 0; j < 9; j++) // Ciclo per le colonne
        {
            stream++; // '['
            if (stream != ' ')
            {
                addHead(pile[j], stream++);
            }
            stream++; // ']'
            stream++; // ' '
        }
        stream++; // ' '
    }

    // Due righe in più inutili
    getline(&line, &len, stream);
    getline(&line, &len, stream);

    while (getline(&line, &len, stream) != -1)
    {
        struct istruzioni istr = (struct istruzioni*) malloc(sizeof(struct istruzioni));
        scanf("move %d from %d to %d", &istr.numToMove, &istr.from, &istr.to);

        for (int i = 0; i < istr.numToMove; i++)
        {
            char tmp = removeTail(pile[istr.from]);
            pile[istr.to] = addTail(pile[istr.to]);
        }
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
