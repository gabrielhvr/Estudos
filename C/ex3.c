#include <stdio.h>



int main(){
int soma = 0;

for(int i=1; i<=4; i++){
    soma+=i;
    if (soma>5)
        break;
}// letra C

printf("\n%d\n", soma);

    for (int h = 0; h < 24; h++) {
        for (int m = 0; m < 60; m++) {
            for (int s = 0; s < 60; s++) {
                printf("\n %d:%d:%d", h, m, s);
                if(s==15)
                    break;
            }
        }
    }

//letra D
}

