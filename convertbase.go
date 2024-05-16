package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimalForm := AtoiBase(nbr, baseFrom)
	strBaseTo := NbrBase(decimalForm, baseTo)

	return strBaseTo
}
