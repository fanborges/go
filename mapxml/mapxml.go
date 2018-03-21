package mapxml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"os"
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


//~ Toma una cadena con los campos del encabezado H1F y los agrega al xml.
func mapeaH1F(cadena[] string){
	
	
	
}


