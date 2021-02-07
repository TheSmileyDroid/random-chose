package main

import (
	"math/rand"
)

func main() {
	genrer := genrerGetter()

	println("Gênero: ", genrer[rand.Intn(len(genrer))])
	theme := themesGetter()

	println("Tema: ", theme[rand.Intn(len(theme))])
	geral := geralGetter()

	println("Assunto aleatório: ", geral)
}
