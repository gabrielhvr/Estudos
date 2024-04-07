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
    if (aluno[qtdAlunos].sexo == 'f' | aluno[qtdAlunos].sexo == 'f'){ //C
        printf(" Informe o nome %i. aluna: ", qtdAlunos);
        qtdFeminino+=1;}
    else if (aluno[qtdAlunos].sexo == "M" | aluno[qtdAlunos].sexo == "m"){
        printf(" Informe o nome do %i. aluno: ", qtdAlunos);
        qtdMasculino+=1;
        }
    else
        printf(" O valor informado para sexo nao eh vailido! ");
    scanf("%s", &aluno[qtdAlunos].Nome);
    printf("Informe a idade de %s: ", aluno[qtdAlunos].Nome);
    scanf("%i", &aluno[qtdAlunos].Idade);
    aluno[qtdAlunos].Notas[0] = 0; // D verdadeiro
    for (int i=1; i <=qtdNotas; i++){
        printf("Informe a %i. nota: ", i);
        scanf ("%f", &aluno[qtdAlunos].Notas[i]);
        aluno[qtdAlunos].Notas[0] += aluno[qtdAlunos].Notas[i];
    }
    aluno[qtdAlunos].Media = aluno[qtdAlunos].Notas[0] / qtdNotas;
    aluno[qtdAlunos].Aprovado = aluno[qtdAlunos].Media >= media;
    printf("Deseja informar o proximo aluno [1/sim ou 2/nao]: ");
    scanf("%i", &Resposta);}
    while (Resposta == 1);
    printf("\n Relatoorio de alunos aprovados ");
    for(int i=1; i<= qtdAlunos; i++)
        if (aluno[i].Aprovado)
            printf("\n * %s", aluno[i].Nome);
        printf("\n Relatorio de alunos Reprovados ");
        for(int i=1; i <= qtdAlunos; i++)
            if(!aluno[i].Aprovado)
                printf("\n * %s", aluno[i].Nome);
        porcentagemFeminino = (float)qtdFeminino/qtdAlunos*100;
        porcentagemMasculino = (float)qtdMasculino/qtdAlunos*100;
        printf ("\n\n Nessa turma %.2f%% sao do sexo Feminino", porcentagemFeminino);
        printf("\n\n Nessa turma %.2f%% sao do sexo Masculino", porcentagemMasculino);

        }