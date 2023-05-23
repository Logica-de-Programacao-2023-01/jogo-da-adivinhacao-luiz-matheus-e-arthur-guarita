# Jogo da Advinhação

O jogo deve funcionar da seguinte forma:

1. O programa deve gerar um número aleatório entre 1 e 100 que será a resposta.
2. O usuário deve tentar adivinhar o número, digitando um valor inteiro.
3. O programa deve comparar o valor digitado com a resposta e dar feedback ao usuário, informando se o número é maior ou menor que a resposta.
4. O usuário deve continuar tentando até acertar o número.
5. O programa deve armazenar o número de tentativas do usuário em um vetor.
6. Quando o usuário acertar o número, o programa deve exibir a quantidade de tentativas que o usuário utilizou para acertar.
7. O programa deve perguntar ao usuário se ele deseja jogar novamente.
8. O programa deve armazenar o número de tentativas utilizadas em todas as jogadas em uma matriz.
9. O usuário pode jogar quantas vezes quiser.
10. O programa deve permitir que o usuário visualize o número de tentativas utilizadas em todas as jogadas.

## Exemplo de saída:

```
Bem-vindo ao jogo da adivinhação!
Digite um número entre 1 e 100: 50
O número é menor que 50.
Digite um número entre 1 e 100: 25
O número é maior que 25.
Digite um número entre 1 e 100: 37
O número é maior que 37.
Digite um número entre 1 e 100: 43
O número é maior que 43.
Digite um número entre 1 e 100: 46
O número é menor que 46.
Digite um número entre 1 e 100: 45
Parabéns, você acertou!
Você utilizou 6 tentativas.
Você deseja jogar novamente? (s/n): s
Digite um número entre 1 e 100: 50
O número é menor que 50.
Digite um número entre 1 e 100: 25
O número é maior que 25.
Digite um número entre 1 e 100: 37
O número é maior que 37.
Digite um número entre 1 e 100: 43
O número é maior que 43.
Digite um número entre 1 e 100: 46
O número é menor que 46.
Digite um número entre 1 e 100: 45
Parabéns, você acertou!
Você utilizou 6 tentativas.
Você deseja jogar novamente? (s/n): n
Você utilizou 6 tentativas na primeira jogada.
Você utilizou 6 tentativas na segunda jogada.
```
