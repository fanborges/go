package main

import (
	"fmt"
	"os"
	"strings"
	"./fileread"
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
	//~ c, _ := ListFiles("./")
	//~ fmt.Printf("%v",c)
	
	slArchivos,_ := fileread.ListFiles("./")
	fmt.Printf("%+v",slArchivos)
	

	cuentaEncabezados := 0
	arrCampos := DevuelveCampos("GLI961030TU5_ACO0510202G0_AAA_196067230_20180212093727.txt")

	for i := 0; i < len(arrCampos); i++ {
		if strings.Contains(arrCampos[i], "[") {
			cuentaEncabezados++
		}
	}

	count := 0
	//~ fmt.Printf("%d", cuentaEncabezados)

	//~ creamos un arreglo de 17 elementos para la cabecera H1F
	var arrH1F [17]string
	for i := 0; i < len(arrCampos); i++ {
		if strings.Contains(arrCampos[i], "[") {
			campo := arrCampos[i]
			if strings.Contains(campo, "H1F_Comprobante") {
				//~ fmt.Printf("%s\n", campo)
				for j := i + 1; j < (i + 17); j++ {
					arrH1F[count] = arrCampos[j]
					//~ fmt.Printf("%s\n", arrH1F[count])
					count++
				}
			}
			//~ fmt.Printf("%s\n", arrCampos[i])
			campo = ""
		}
	}

	mH1F := mapH1F(arrH1F)
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

//~ Toma el nombre de un archivo de factura y devuelve sus campos en un arreglo
func DevuelveCampos(fileName string) []string {
	archivo, _ := ReadFile(fileName)
	contador := 0
	var sl []string
	campo := ""
	for i := 0; i < len(archivo)-1; i++ {
		if archivo[i] != '|' {
			//~ fmt.Printf("%c",archivo[i])
			campo += string(archivo[i])
		} else {
			//~ sl[contador] = campo
			sl = append(sl, campo)
			contador++
			campo = ""
		}
	}
	return sl
}

//~ Toma el nombre de un archivo y devuelve un string con su contenido
func ReadFile(f string) (string, error) {
	contenido := ""
	file, err := os.Open(f)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return contenido, err
	}
	buffer := make([]byte, 100)
	for n, e := file.Read(buffer); e == nil; n, e = file.Read(buffer) {
		if n > 0 {
			//~ os.Stdout.Write(buffer[0:n])
			contenido += string(buffer[0:n])
		}
	}
	return contenido, err
}

//~ Lee el nombre de un archivo e imprime su contenido en la consola
func ReadAndPrint(f string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	buffer := make([]byte, 100)
	for n, e := file.Read(buffer); e == nil; n, e = file.Read(buffer) {
		if n > 0 {
			os.Stdout.Write(buffer[0:n])
		}
	}
}

//~ Lee un path y devuelve un slice con la lista de archivos
func ListFiles(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
