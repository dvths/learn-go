// Echo3 exibe seus argumentos de linha de comango
// Esta versão simplificada utiliza a função `Join` do pacote `strings`
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	//fmt.Println(os.Args[1:])
	/*
	   Se não estivermos interessados em formatação, mas quisermos apenas ver os valores,
	   talvez para depuração, podemos deixar `Println` formatar o resultado.
	*/

}
