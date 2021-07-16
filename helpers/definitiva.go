package helpers

func CalcularDefinitiva(nota_1 float64, nota_2 float64, nota_3 float64) (def float64) {

	def = ((nota_1 * 0.35) + (nota_2 * 0.35) + (nota_3 * 0.3))

	return def
}
