#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[])
{
    FILE *f;
    char *line;

    f = fopen(argv[1], "r");

    if (f == NULL)
    {
        printf ("%s is not a valid file\n", argv[1]);
        return(1);
    }

    while (fgets(line, 1024, f))
    {
        if (line[0] == '\n') {
            continue;
        }
        size_t ln = strlen(line) - 1;
        if (line[ln] == '\n')
            line[ln] = ' ';

        char *word;
        char *longest;
        longest = "";
        word = strtok (line, " ");
        while (word != NULL)
        {
            if (strlen (word) > strlen (longest))
            {
                longest = word;
            }
            word = strtok (NULL, " ");
        }
        printf ("%s\n", longest);
    }

    fclose(f);
    return(0);
}
