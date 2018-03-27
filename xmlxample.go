package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func main() {

	type H1C struct {
		XMLName             xml.Name `xml:"cfdi:ejemplo"`
		NumeroOrdenDeCompra string   `xml:"orden"`
		FechaOrdenDeCompra  string   `xml:"fechaorden"`
	}

	type Comprobante struct {
		XMLName           xml.Name `xml:"cfdi:comprobante"`
		Version           string   `xml:"version,attr"`
		Serie             string   `xml:"serie,attr"`
		Folio             string   `xml:"folio,attr"`
		Fecha             string   `xml:"fecha,attr"`
		FormaPago         string   `xml:"formaPago,attr"`
		CondicionesDePago string   `xml:"condicionesDePago,attr"`
		SubTotal          string   `xml:"subtoTal,attr"`
		Descuento         string   `xml:"descuento,attr"`
		Moneda            string   `xml:"moneda,attr"`
		TipoCambio        string   `xml:"tipoCambio,attr"`
		Total             string   `xml:"total,attr"`
		TipoDeComprobante string   `xml:"tipoDeComprobante,attr"`
		MetodoDePago      string   `xml:"metodoDePago,attr"`
		LugarExpedicion   string   `xml:"lugarExpedicion,attr"`
		Confirmacion      string   `xml:"confirmacion,attr"`
		Sello             string   `xml:"sello,attr"`
		Certificado       string   `xml:"certificado,attr"`
		//~ Hola string `xml:"hola>mundo>cruel"`
		H1C
	}

	//~ Creamos el buffer de Bytes y agregamos la cabecera
	w := &bytes.Buffer{}
	w.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"))
	//~ Creamos la estructura xml y la agregamos al buffer de bytes
	h := &Comprobante{Version: "3.3", Serie: "AAA", Folio: "196067230"}
	//~ Agregamos la estructura H1C
	h.H1C = H1C{NumeroOrdenDeCompra: "hola mundo", FechaOrdenDeCompra: "fecha"}
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	if err := enc.Encode(h); err != nil {
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
