# Lendo argumentos de linha de comando

Processar alguma entrada a fim de produzir alguma saída é, basicamente, a definição de computação. Alguns programas podem gerar seus próprios dados, porem, como mais frequencia, a entrada vem de uma fonte externa: um arquivo, uma conexão de rede, a saída de outro programa, um usuário cem um teclado, argumentos de linha de comando ou algo semelhante.

Iniciando o estudo de "A Linguagem de Programação Go", de Alan A. A. Donovan e Brian Q. Kernighan, nos deparamos com uma série de tutoriais introdutórios que abordam algumas dessas alternativas começando pelos argumetos de linha de comando que são estudados a partir de agora.

Como exemplo, construiremos algumas variações do famoso comando `echo` do Unix:

## Echo 1

```go
// Echo1 exibe seus argumentos de linha de comando
package main

import (
    "fmt"
    "os"
)
func main() {
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
       s += sep + os.Args[i] 
       sep = " "
    }
    fmt.Println(s)
}
```
Para lidar com o sistema operacional, independentemente da plataforma, usamos o pacote `os` que nos dá acesso a variável `Args` que armazenará os argumentos. A variável `Args` é uma *fatia* (sliece) de strings. Fatia podem ser pensadas, nesse momento como uma sequencia **s** de elementos de uma array dimensionada dinamicamente, em que elementos individuais podem ser acessados como `s[i]` e uma subsequencias contígua, como `s[m:n]`. O numero de elementos é dado por `len(s)`. Com intervalos *semiabertos* que incluem o primeiro item, mas excluem o último, pois isso simplifica a lógica. 

Fora do pacote `os`, seu nome é `os.Args` e seu primeiro elemento e `os.Args[0]`, que, em nosso exemplo se refere ao nome do programa propriamente dito; os outros elementos são os argumentos apresentados ao programa quando sua execução for iniciada. 

Uma expressão de fatia na forma `s[m:n]` produz uma fatia que se refere aos elementos `m` a `n -1`, protanto, os elementos de que precisamos em nosso exemplo são aqueles que estão na fatia `os.Args[1:len(os.Args)]`.Se `m`ou `n` forem omitidos, os defaults serão 0 ou `len(s)` respectivamente, desta forma, podemos abreviar a fatia como `os.Args[1:]`

No pacote *echo1*, após as importações, na função `main`, são declaradas duas variáveis `s` e `sep` do tipo string, inicializadas implicitamente (como seus *valores zero*). A expressão `sep + os.Args[i]` representa a concatenação das strings `sep` e os argumrntos armazenados em `os.Args[]` e a instrução `s += sep + os.Args[i]` é uma *istrução de atribuição* que seria o mesmo que  
```
    s = s + sep + os.Args[i]
```
em que o resultado é atribuído de volta a `s`.

Em outras palavras, a string `s` nasce vazia, isto é, seu valor zero é "", e cada passagem do loop acrescenta um texto a ela; após a primeira iteração, haverá um espaço entre as strings. (Note que, a última linha do corpo do loop reatribui um espaço vazio a `sep`).

```
for inicialização; cindução; pós {
    // zero ou mais instruções
}

// Um loop "while" tradicional
for condição {
     //...
}

// um loop infinito tradicional 
for {
    //...
}
```
Em Go, parenteses jamais são usados em torni dis três componemte de um loop `for`.

Outro detalhe importante é que `i++` é uma instrução de incremento e não uma
expressão, como é o caso de muitas linguagens da familia C. Portanto, `j = i++`
não é permitido. Outro detalhe é que apenas a forma pósfixa é permitida, isto
é `++i` também não existe em Go.

Esse é um processo quaderático que pode ser custoso se o número de argumentos for grande. Os próximos exemplos mostraram versões melhoradas do programa.

## Echo 2
Outra forma de loop `for` permite iterar sobre um intervalo (*range*) de
valores der um tipo de dado como uma string ou uma fatia.

```go 
package main

import (
        "fmt"
        "os"
)

func main() {
    s, sep := "", ""
        for _, arg := range os.Args[1:] {
            s += sep + arg
            sep = " "
        }
        fmt.Println(s)
}
```


