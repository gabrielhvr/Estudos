#include <stdio.h>
#include <string.h>


typedef struct{
    int dia, mes, ano;
}DataNas;

typedef struct{
    DataNas dataNas;
    int idade;
    char sexo;
    char nome[100];

}Pessoa;



int main () {

    Pessoa pessoa;

    printf("Digite sua data de nascimento no formato dd mm aaaa:");
    scanf("%d%d%d", &pessoa.dataNas.dia,&pessoa.dataNas.mes, &pessoa.dataNas.ano);

printf("Data: %d/%d/%d", pessoa.dataNas.dia,pessoa.dataNas.mes,pessoa.dataNas.ano);







return 0;
}
