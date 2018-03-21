package main

import (
	"fmt"
	"strings"
	"./filextract"
	"./cfdi"
)

type H1F struct {
	version           string
	serie             string
	folio             string
	fecha             string
	formaPago         string
	condicionesDePago string
	subTotal          string
	descuento         string
	moneda            string
	tipoCambio        string
	total             string
	tipoDeComprobante string
	metodoDePago      string
	lugarDeExpedicion string
	confirmacion      string
	sello             string
	certificado       string
}

func main() {

	//~ Ejemplo: lee el directorio actual e imprime el slice con todos los nombres de
	//~ archivos
	//~ c, _ := filextract.ListFiles("./")
	//~ fmt.Printf("%v",c)

	cuentaEncabezados := 0
	arrCampos := filextract.GetFields("GLI961030TU5_ACO0510202G0_AAA_196067230_20180212093727.txt")

	for i := 0; i < len(arrCampos); i++ {
		if strings.Contains(arrCampos[i], "[") {
			cuentaEncabezados++
		}
	}
	
	sl_encabezados := cfdi.GetEncabezados(arrCampos)
	fmt.Printf("%v", sl_encabezados)
	
	h1f := cfdi.GetH1F(arrCampos)

	mH1F := mapH1F(h1f)
	fmt.Printf("%+v", mH1F)
	for key, value := range mH1F {
		fmt.Println("key:", key, "value:", value)
	}
}



//~ toma un arreglo de strings con los campos de la cabecera H1F y devuelve un map
//~ key value con sus respectivos valores
func mapH1F(cadena [17]string) map[string]string {
	//~ Mapa para cabecera H1F
	mH1F := map[string]string{
		"version":           "",
		"serie":             "",
		"folio":             "",
		"fecha":             "",
		"formaPago":         "",
		"condicionesDePago": "",
		"subTotal":          "",
		"descuento":         "",
		"moneda":            "",
		"tipoCambio":        "",
		"total":             "",
		"tipoDeComprobante": "",
		"metodoDePago":      "",
		"lugarDeExpedicion": "",
		"confirmacion":      "",
		"sello":             "",
		"certificado":       "",
	}
	contador := 0
	//~ Vamos a iterar el map que creamos para la cabecera H1F
	for key := range mH1F {
		mH1F[key] = cadena[contador]
		contador++
	}
	return mH1F
}


