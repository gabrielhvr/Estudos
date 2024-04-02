#include <stdio.h>




int main (){
int Total = 0;
for (int j=0; j<7; j++){
    if (j==5)
        continue; // pula para proxima execução no j=5
    Total= Total +j;


}
printf("Total: %d", Total);
} // letra D

