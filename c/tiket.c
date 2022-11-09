// you are making a ticketing system, for children under 3 years old, they get in for free


#include <stdio.h>
#include <stdlib.h>

int main()
{
    int age;
    int total = 0;
    printf("Enter your age: ");
    scanf("%d", &age);
    if (age < 3)
    {
        total += 0;
        printf("You get in for free!");
    }
    else
    {   
        total += 100;
        printf("You have to pay!", total);
    }
    return 0;
}
