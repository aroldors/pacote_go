package pacote

import "math"

// Pi é uma peoporção numérica definida pela relação entre
// o perímetro de uma circunferência e seu diâmetro
const Pi = 3.1416

// Circ é responsável em calcular a área de uma circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável em calcular a área de um quadrado
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
