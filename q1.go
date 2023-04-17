package main

import "fmt"

func calculardesconto(compraatual float64, historico []float64) (float64, error) {
	var soma float64

	for _, i := range historico {
		soma += i

	}
	somadecompras := compraatual + soma
	media := somadecompras / float64(len(historico))

	if somadecompras > 1000 {

		desconto1 := compraatual * 0.9

		fmt.Println("Como o valor total de suas compras na loja foi maior do que 1000.00, o seu desconto é de 10%, ficando ", desconto1)
		return desconto1, nil

	}
	if somadecompras <= 1000 {

		desconto2 := compraatual * 0.95
		fmt.Println("Como o valor total de suas compras na loja foi menor ou igual a  1000.00, o seu desconto é de 5%, ficando ", desconto2)
		return desconto2, nil

	}
	if somadecompras <= 500 {
		desconto3 := compraatual * 0.98
		fmt.Println("Como o valor total de suas compras na loja foi menor do que 500.00, o seu desconto é de 2%, ficando ", desconto3)
		return desconto3, nil

	}
	if len(historico) == 0 {
		desconto4 := compraatual * 0.9

		fmt.Println("Como o valor total de suas compras na loja foi maior do que 1000.00, o seu desconto é de 10%, ficando ", desconto4, compraatual)
		return desconto4, nil

	}
	if media > 1000 {

		desconto5 := somadecompras * 0.8

		fmt.Printf("Como o valor total de suas compras na loja foi maior do que 1000.00, o seu desconto é de 20%%, ficando %.2f", desconto5)
		return desconto5, nil
	}
	return somadecompras, nil
}

func main() {
	var resposta string
	var historico []float64
	var valores float64
	var tamanho int
	var compraatual float64

	fmt.Println("Quanto foi sua compra?")
	fmt.Scanln(&compraatual)
	desconto4 := compraatual * 0.9
	fmt.Println("Já comprou conosco antes?")
	fmt.Scanln(&resposta)
	if resposta == "sim" {
		fmt.Println("Quantas vezes?")
		fmt.Scanln(&tamanho)

		for i := 1; i <= tamanho; i++ {

			fmt.Printf("Digite o valor da compra %d", i)
			fmt.Scanln(&valores)
			historico = append(historico, valores)

		}
	}

	if resposta == "nao" {
		fmt.Printf("Seu desconto é de 10%%, portanto vai custar %2.f", desconto4)

	} else {
		calculardesconto(compraatual, historico)
	}
}
