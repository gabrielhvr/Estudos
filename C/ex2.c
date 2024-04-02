#include <stdio.h>
#include <unistd.h> // Para a função sleep
#include <stdlib.h> // Para a função system

int main() {
    for (int h = 0; h < 24; h++) {
        for (int m = 0; m < 60; m++) {
            for (int s = 0; s < 60; s++) {
                printf("\n %d:%d:%d", h, m, s);
                sleep(1); // Espera por 1 segundo
                system("clear"); // Limpa a tela
            }
        }
    }

    return 0;
}
//letra A, a função sleep vai executando a 1 segundo. em 6 segundo

