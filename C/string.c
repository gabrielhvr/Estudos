#include <stdio.h>
#include <string.h>

void removerCaractere(char *str, int index) {
    int i;
    int tamanho = strlen(str);

    // Desloca os caracteres à direita do índice um para a esquerda
    for (i = index; i < tamanho - 1; i++) {
        str[i] = str[i + 1];
    }

    // Adiciona o caractere nulo ao final da nova string
    str[tamanho - 1] = '\0';
}

int main() {
    char nome[] = "ggabriel";
    printf("String original: %s\n", nome);

    // Removendo o primeiro caractere
    removerCaractere(nome, 0);

    printf("String modificada: %s\n", nome);

    return 0;
}
