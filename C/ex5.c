#include <stdio.h>

int main(){
// no switch case, as chaves são opcionais
    int escolha;

    scanf("%d",&escolha);

    switch(escolha){
        case 1:
        {
            printf("1");
            break;
        }
        case 2:
            printf("Escolheu 2");
            break;
            
        default:
            printf("nao escolheu nenhum!");

}}