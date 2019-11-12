package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Variaveis Globais
var moeda = "9"
var valorFixoSantander = "9"

func main() {

	fmt.Println(gerarLinhaDigitavelSantander("33", "617756-5", "123456789098", "1234",
		"2002-12-01T12:18:07.990Z", "835,89"))
}

// Funcao para calcular Digito Verificador Modulo 10
func modulo10(num string) int {
	// Variaveis de controle
	tamanhoString := len(num) + 1
	soma := 0
	resto := 0
	dv := 0
	numeros := make([]string, tamanhoString)
	multiplicador := 2
	runes := []rune(num)
	for i := len(num); i > 0; i-- {
		// Multiplica da direita pra esquerda, alternando os algarismos 2 e 1
		if multiplicador%2 == 0 {
			// Pega cada numero isoladamente
			conteudo := string(runes[i-1 : i])
			conteudoInt, _ := strconv.Atoi(conteudo)
			calculo := conteudoInt * 2
			numeros[i] = strconv.Itoa(calculo)
			multiplicador = 1
		} else {
			conteudo := string(runes[i-1 : i])
			conteudoInt, _ := strconv.Atoi(conteudo)
			calculo := conteudoInt * 1
			numeros[i] = strconv.Itoa(calculo)
			multiplicador = 2
		}
	}
	for i := (len(numeros) - 1); i > 0; i-- {
		conteudo := numeros[i]
		conteudoInt, _ := strconv.Atoi(conteudo)
		aux := strconv.Itoa(conteudoInt)
		auxiliar := []rune(aux)

		if len(auxiliar) > 1 {
			aux2 := string(auxiliar[0 : len(auxiliar)-1])
			aux3 := string(auxiliar[len(auxiliar)-1 : len(auxiliar)])
			aux4, _ := strconv.Atoi(aux2)
			aux5, _ := strconv.Atoi(aux3)
			aux6 := aux4 + aux5
			variavel := strconv.Itoa(aux6)
			numeros[i] = variavel
		} else {
			numeros[i] = aux
		}
	}
	for i := len(numeros); i > 0; i-- {
		if len(numeros[i-1]) > 0 {
			conteudoSoma, _ := strconv.Atoi(numeros[i-1])
			soma = soma + conteudoSoma
		}
	}
	resto = soma % 10
	dv = 10 - resto
	if dv == 10 {
		dv = 0
	}
	// Retorna o digito verificador
	return dv
}

// Funcao para calcular Digito Verificador Modulo 11
func modulo11(num string) int {
	// Variaveis de controle
	tamanhoString := len(num) + 1
	soma := 0
	resto := 0
	dv := 0
	numeros := make([]string, tamanhoString)
	multiplicador := 2
	runes := []rune(num)
	for i := len(num); i > 0; i-- {
		// Multiplica da direita pra esquerda, incrementando o multiplicador de 2 a 9
		// Caso o multiplicador seja maior que 9 o mesmo recomeça em 2
		if multiplicador > 9 {
			// Pega cada numero isoladamente
			multiplicador = 2
			conteudo := string(runes[i-1 : i])
			conteudoInt, _ := strconv.Atoi(conteudo)
			calculo := conteudoInt * multiplicador
			numeros[i] = strconv.Itoa(calculo)
			multiplicador++
		} else {
			conteudo := string(runes[i-1 : i])
			conteudoInt, _ := strconv.Atoi(conteudo)
			calculo := conteudoInt * multiplicador
			numeros[i] = strconv.Itoa(calculo)
			multiplicador++
		}
	}
	// Realiza a soma de todos os elementos do array e calcula o digito verificador
	// na base 11 de acordo com a regra.
	for i := len(numeros); i > 0; i-- {
		if len(numeros[i-1]) > 0 {
			conteudoSoma, _ := strconv.Atoi(numeros[i-1])
			soma = soma + conteudoSoma
		}
	}
	resto = soma % 11
	dv = 11 - resto
	if dv > 9 || dv == 0 {
		dv = 1
	}
	// Retorna o digito verificador
	return dv
}

// Funcao para calcular Digito Verificador Modulo 11 - "Santander"
func modulo11Santander(num string) int {
	// Variaveis de controle
	tamanhoString := len(num) + 1
	soma := 0
	resto := 0
	dv := 0
	numeros := make([]string, tamanhoString)
	multiplicador := 2
	runes := []rune(num)
	for i := len(num); i > 0; i-- {
		// Multiplica da direita pra esquerda, incrementando o multiplicador de 2 a 9
		// Caso o multiplicador seja maior que 9 o mesmo recomeça em 2
		if multiplicador > 9 {
			// Pega cada numero isoladamente
			multiplicador = 2
			conteudo := string(runes[i-1 : i])
			conteudoInt, _ := strconv.Atoi(conteudo)
			calculo := conteudoInt * multiplicador
			numeros[i] = strconv.Itoa(calculo)
			multiplicador++
		} else {
			conteudo := string(runes[i-1 : i])
			conteudoInt, _ := strconv.Atoi(conteudo)
			calculo := conteudoInt * multiplicador
			numeros[i] = strconv.Itoa(calculo)
			multiplicador++
		}
	}
	// Realiza a soma de todos os elementos do array e calcula o digito verificador
	// na base 11 de acordo com a regra.
	for i := len(numeros); i > 0; i-- {
		if len(numeros[i-1]) > 0 {
			conteudoSoma, _ := strconv.Atoi(numeros[i-1])
			soma = soma + conteudoSoma
		}
	}
	resto = soma % 11
	if resto == 0 {
		dv = 0
	} else if resto == 1 {
		dv = 0
	} else if resto == 10 {
		dv = 1
	} else if resto > 9 {
		dv = 0
	} else {
		dv = 11 - resto
	}
	// Retorna o digito verificador
	return dv
}

// Funcao para Formatar Data De Vencimento Para Realizar o Calculo Do Fator De Vencimento
func formatarDataVencimento(stringDataVencimento string) string {
	runes := []rune(stringDataVencimento)
	dataVencimentoParte1 := string(runes[8:10])
	dataVencimentoParte2 := string(runes[5:7])
	dataVencimentoParte3 := string(runes[0:4])
	stringData := dataVencimentoParte1 + "/" + dataVencimentoParte2 + "/" + dataVencimentoParte3
	return stringData
}

// Funcao para Formatar Numero de Convenio - "Santander"
func formatarConvenioSantander(stringConvenio string) string {
	stringConvenio = strings.Replace(stringConvenio, "-", "", -1)
	return stringConvenio
}

// Funcao para Formatar Valor do Boleto
func formatarValorBoleto(stringValorBoleto string) string {
	stringValorBoleto = strings.Replace(stringValorBoleto, ",", "", -1)
	stringValorBoleto = strings.Replace(stringValorBoleto, ".", "", -1)
	return stringValorBoleto
}

// Funcao para Acrescentar Numeros Zeros Entre Fator de Vencimento e Valor do Boleto
func acrescentarNumerosZeros(stringValorBoleto string) string {
	stringValorBoleto = formatarValorBoleto(stringValorBoleto)
	for i := 0; len(stringValorBoleto) < 10; i++ {
		stringValorBoleto = "0" + stringValorBoleto
	}
	return stringValorBoleto
}

// Funcao para Gerar Fator de Vencimento de Acordo com a Data de Vencimento
func gerarFatorVencimentoGenerico(stringVencimentoDocumento string) string {
	dataVencimento := formatarDataVencimento(stringVencimentoDocumento)
	runes := []rune(dataVencimento)
	dataParte1 := string(runes[0:2])
	dataParte2 := string(runes[3:5])
	dataParte3 := string(runes[6:10])
	int1, _ := strconv.Atoi(dataParte1)
	int2, _ := strconv.Atoi(dataParte2)
	int3, _ := strconv.Atoi(dataParte3)
	st := time.Date(1997, time.Month(10), 7, 00, 00, 00, 0, time.UTC)
	en := time.Date(int3, time.Month(int2), int1, 00, 00, 00, 0, time.UTC)
	diff := (en.Sub(st))
	calculoMilisegundo := diff.Milliseconds() + 3600000
	fatorVencimento := calculoMilisegundo / 86400000
	retorno := strconv.FormatInt(fatorVencimento, 10)
	return retorno
}

// Funcao para Gerar Fator de Vencimento de Acordo com a Data de Vencimento
func gerarCampoLivreSantander(stringValorFixo string, stringConvenio string,
	stringNossoNumero string, stringDvNossoNumero, stringCarteira string) string {
	return stringValorFixo + stringConvenio + stringNossoNumero + stringDvNossoNumero + stringCarteira
}

// Funcao para Gerar Digito Verificador Nosso Numero - "Santander"
func gerarDvNossoNumeroSantander(stringNossoNumero string) string {
	digitoVerificador := modulo11Santander(stringNossoNumero)
	return strconv.Itoa(digitoVerificador)
}

// Funcao para Gerar Linha Digitavel Santander
func gerarLinhaDigitavelSantander(codigo_banco string, cod_empresa_banco string,
	nossoNumero string, carteira string, vencimento string, valor_boleto string) string {
	convenio := formatarConvenioSantander(cod_empresa_banco)
	digitoVerificador := gerarDvNossoNumeroSantander(nossoNumero)
	valorFixo := valorFixoSantander
	campoLivreSantander := gerarCampoLivreSantander(valorFixo, convenio, nossoNumero, digitoVerificador, carteira)
	campo1 := campo1(codigo_banco, campoLivreSantander)
	campo2 := campo2(campoLivreSantander)
	campo3 := campo3(campoLivreSantander)
	campo5 := campo5(vencimento, valor_boleto)
	campo4 := campo4(codigo_banco, campo5, nossoNumero, cod_empresa_banco)
	retorno := campo1 + campo2 + campo3 + campo4 + campo5
	return retorno
}

// Funcao para Gerar Campo 1 da Linha Digitavel
func campo1(codigo_banco string, campoLivreSantander string) string {
	campoLivre := campoLivreSantander
	runes := []rune(campoLivre)
	campoLivre = string(runes[0:5])
	if codigo_banco == "33" {
		codigo_banco = "033"
	}
	dvVerificador := strconv.Itoa(modulo10(codigo_banco + moeda + campoLivre))
	return codigo_banco + moeda + campoLivre + dvVerificador
}

// Funcao para Gerar Campo 2 da Linha Digitavel
func campo2(campoLivreSantander string) string {
	campoLivre := campoLivreSantander
	runes := []rune(campoLivre)
	campoLivre = string(runes[5:15])
	dvVerificador := strconv.Itoa(modulo10(campoLivre))
	return campoLivre + dvVerificador
}

// Funcao para Gerar Campo 3 da Linha Digitavel
func campo3(campoLivreSantander string) string {
	campoLivre := campoLivreSantander
	runes := []rune(campoLivre)
	campoLivre = string(runes[15:25])
	dvVerificador := strconv.Itoa(modulo10(campoLivre))
	return campoLivre + dvVerificador
}

// Funcao para Gerar Campo 4 da Linha Digitavel
func campo4(codigo_banco string, campo5 string, nossoNumero string, cod_empresa_banco string) string {
	convenio := formatarConvenioSantander(cod_empresa_banco)
	dvVerificador := strconv.Itoa(modulo11(codigo_banco + moeda + campo5 + nossoNumero + convenio))
	return dvVerificador
}

// Funcao para Gerar Campo 5 da Linha Digitavel
func campo5(vencimento string, valor_boleto string) string {
	fatorVencimento := gerarFatorVencimentoGenerico(vencimento)
	valorBoleto := formatarValorBoleto(valor_boleto)
	valorBoletoFormatado := acrescentarNumerosZeros(valorBoleto)
	return fatorVencimento + valorBoletoFormatado
}
