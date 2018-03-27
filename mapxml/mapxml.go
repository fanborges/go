package mapxml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	//~ "log"
	//~ "os"
)

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

}

//~ Toma una slice con los campos del encabezado H1F devuelve
//~ la estructura Comprobante con sus respectivos campos
func STH1F(c [17]string) *Comprobante {

	com := &Comprobante{Version: c[0], Serie: c[1], Folio: c[2], Fecha: c[3],
		FormaPago: c[4], CondicionesDePago: c[5], SubTotal: c[6], Descuento: c[7],
		Moneda: c[8], TipoCambio: c[9], Total: c[10], TipoDeComprobante: c[11],
		MetodoDePago: c[12], LugarExpedicion: c[13], Confirmacion: c[14],
		Sello: c[15], Certificado: c[16]}

	return com
}

//~ Toma una estructura Comprobante e imprime la respectiva codificación
//~ a xml (provisional)

func EncodeXML(h *Comprobante) {
	//~ Creamos el buffer de Bytes y agregamos la cabecera
	w := &bytes.Buffer{}
	w.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"))
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	if err := enc.Encode(h); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println(w.String())
}
