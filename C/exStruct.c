#include <stdio.h>
#include <stdbool.h>


struct Endereco{
    char logradouro[3];
    int numero;
};

struct Alunos {
    char Nome[30];
    int Idade;
    bool Aprovado;
    float Notas[10];
    char sexo;
    float Media;
    struct Endereco endComercial;
    struct Endereco endResidencial;
};

int main() {
    struct Alunos aluno[10];
    int qtdNotas, Resposta, qtdAlunos=0, qtdFeminino, qtdMasculino=0;
    float media, porcentagemMasculino=0, porcentagemFeminino=0;
    printf("Informe a media de aprovacao na faulcade: ");
    scanf("%f", &media);
    printf("Informe a quantidade de notas: ");
    scanf("%i",&qtdNotas);
do{
    qtdAlunos+=1;
        printf("Informe o sexo do %i. aluno(a)[M/F]: ", qtdAlunos);
        scanf("%s", &aluno[qtdAlunos].sexo);
    if (aluno[qtdAlunos].sexo == "F" or aluno)
        }
}

}