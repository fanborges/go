package cfdi
import(
	"strings"
)

//~ Toma un slice con los campos del cfdi y devuelve un
//~ arreglo con los elementos de la cabecera H1F_Comprobante
func GetH1F(cadena []string) [17]string {
	count :=0
	var arrH1F [17]string
	for i := 0; i < len(cadena); i++ {
		if strings.Contains(cadena[i], "[") {
			campo := cadena[i]
			if strings.Contains(campo, "H1F_Comprobante") {
				for j := i + 1; j < (i + 17); j++ {
					arrH1F[count] = cadena[j]
					count++
				}
			}
			campo = ""
		}
	}	
	return arrH1F
}

//~ Toma un slice con los campos de la factura y devuelve otro
//~ slice con los encabezados.
func GetEncabezados(cadena []string) []string{
	var slice_encabezados [] string
		for i := 0; i < len(cadena); i++ {
		if strings.Contains(cadena[i], "[") {
			slice_encabezados = append(slice_encabezados,cadena[i])
		}	
	}
	return slice_encabezados
}
