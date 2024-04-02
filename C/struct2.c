#include <stdio.h>
#include <string.h>


typedef struct{
	char raca[10];
	float peso;
}caracDog;

typedef struct{
	char nome[10];
	int olho;
	caracDog caracdog;
}Dog;


int main (){
	Dog holy;
	
	strcpy(holy.nome, "holy");
	holy.olho = 2;
	strcpy(holy.caracdog.raca, "husky");
	holy.caracdog.peso = 12.5;
	
	printf("nome: %s, quantidade de olhos: %d, raca: %s, peso: %f", holy.nome, holy.olho, holy.caracdog.raca, holy.caracdog.peso);

	





return 0;
}


