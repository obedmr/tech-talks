#include <stdio.h>

int main() {

    char *p;
    char *q = NULL;


    printf("Address of p = %u\n", p);
    strcpy(p, "Hello");
    printf("%s\n", p);
    printf("About to copy \"Goodbye\" to q\n");
    strcpy(q, "Goodbye");
    printf("String copied\n");
    printf("%s\n", q);

    return 0;
}
