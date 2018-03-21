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
				//~ fmt.Printf("%s\n", campo)
				for j := i + 1; j < (i + 17); j++ {
					arrH1F[count] = cadena[j]
					//~ fmt.Printf("%s\n", arrH1F[count])
					count++
				}
			}
			//~ fmt.Printf("%s\n", arrCampos[i])
			campo = ""
		}
	}	
	return arrH1F
}

func GetEncabezados(cadena []string) []string{
	var slice_encabezados [] string
		for i := 0; i < len(cadena); i++ {
		if strings.Contains(cadena[i], "[") {
			slice_encabezados = append(slice_encabezados,cadena[i])
		}	
	}
	return slice_encabezados	
}
