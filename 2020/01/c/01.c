#include <stdio.h>
#include <stdlib.h>

int main()
{
    FILE *fptr = fopen("challenge.txt", "r");

    int number;
    int numbers[200] = {0};

    for (int i = 0; i < 200; i++)
    {
        if (fscanf(fptr, "%d", &number) > 0)
        {
            numbers[i] = number;
        }
    }

    for (int i = 0; i < 200; i++)
    {
        for (int j = 0; j < 200; j++)
        {
            if (numbers[i] + numbers[j] == 2020)
            {
                printf("%d\n", numbers[i] * numbers[j]);
            }
        }
    }
}