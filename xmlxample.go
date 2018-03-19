package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"log"
)

func main() {
		
	type Comprobante struct{
		XMLName xml.Name `xml:"cfdi:comprobante"`
		Version string `xml:"version,attr"`
		Serie string `xml:"serie,attr"`
		Folio string `xml:"folio,attr"`
		Fecha string `xml:"fecha,attr"`
		FormaPago string `xml:"formaPago,attr"`
		CondicionesDePago string `xml:"condicionesDePago,attr"`
		SubTotal string `xml:"subtoTal,attr"`
		Descuento string `xml:"descuento,attr"`
		Moneda string `xml:"moneda,attr"`
		TipoCambio string `xml:"tipoCambio,attr"`
		Total string `xml:"total,attr"`
		TipoDeComprobante string `xml:"tipoDeComprobante,attr"`
		MetodoDePago string `xml:"metodoDePago,attr"`
		LugarExpedicion string `xml:"lugarExpedicion,attr"`
		Confirmacion string `xml:"confirmacion,attr"`
		Sello string `xml:"sello,attr"`
		Certificado string `xml:"certificado,attr"`		
	}
	
	type H1C struct{
		numero string `xml:"numeroOrdenDeCompra"`
		fecha string `xml:"fechaOrdenDecompra`
		
	}
	
	
	//~ Creamos el buffer de Bytes y agregamos la cabecera
	w := &bytes.Buffer{}
	w.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"))

	//~ Creamos el xml y lo agregamos al buffer de bytes
	h := &Comprobante{Version: "3.3", Serie: "AAA", Folio: "196067230", Fecha: "2018-02-12T09:37:27"}
	//~ en := xml.NewEncoder(os.Stdout)
	en := xml.NewEncoder(w)
	en.Indent("", "  ")
	if err := en.Encode(h); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println(w.String())	

    //~ Creamos el archivo xml y le agregamos el slice de Bytes		
	file, err := os.Create("result.xml")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
    defer file.Close()
	file.WriteString(w.String())	
}
