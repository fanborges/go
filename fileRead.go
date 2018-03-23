package main

import (
	//~ "fmt"
	//~ "strings"
	"./filextract"
	"./cfdi"
	"./mapxml"
)

func main() {
	//~ Leer el archivo txt de facturación y obteber los campos
	arrCampos := filextract.GetFields("GLI961030TU5_ACO0510202G0_AAA_196067230_20180212093727.txt")
	//~ Obtener los campos de la cabecera H1F
	h1f := cfdi.GetH1F(arrCampos)
	//~ Imprimir el xml con los campos
	cc := mapxml.STH1F(h1f)
	mapxml.EncodeXML(cc)
}
