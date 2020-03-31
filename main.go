package goarea

// Constante de area
const Pi = 3.1416

// responsavel por calcular a area da circuferencia
func Circulo(raio float64) float64 {
	return Pi * raio * raio
}

// Area da um retangulo
func Retangulo(base, altura float64) float64 {
	return base * altura
}
