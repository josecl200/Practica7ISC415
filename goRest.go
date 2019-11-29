package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"unicode"

	"github.com/go-resty/resty"
)

// Create a Resty Client

type estudiante struct {
	Matricula int    `json:"matricula"`
	Nombre    string `json:"nombre"`
	Correo    string `json:"correo"`
	Carrera   string `json:"carrera"`
}

func getEstudiantes() []estudiante {
	client := resty.New()
	var ests []estudiante
	resp, _ := client.R().
		EnableTrace().
		Get("http://localhost:4567/rest/estudiantes/")
	body := resp.Body()
	json.Unmarshal(body, &ests)
	return ests
}

func getEstudiante(matricula int) estudiante {
	client := resty.New()
	var est estudiante
	resp, _ := client.R().
		EnableTrace().
		Get("http://localhost:4567/rest/estudiantes/" + strconv.Itoa(matricula))
	json.Unmarshal(resp.Body(), &est)
	return est
}

func crearEstudiante(newEst estudiante) (res bool) {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody(newEst).
		Post("http://localhost:4567/rest/estudiantes/")
	println(resp.Body())
	if err != nil {
		res = false
	} else {
		res = true
	}
	return
}

func clearMenu() {
	clear := make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func initMenu() {
	clearMenu()
	fmt.Println("Bienvenido al demo de API REST para ISC-415, digite el número necesario para lo que quiere hacer, sino digite q para salir")
	fmt.Println("1) Ver todos los estudiantes")
	fmt.Println("2) Buscar un estudiante por matricula")
	fmt.Println("3) Registrar un nuevo estudiante")
}

func main() {
	doIt := true
	var i rune
	for doIt {
		initMenu()
		_, _ = fmt.Scanf("%c\n", &i)
		switch i {
		case '1':
			estudiantes := getEstudiantes()
			for _, est := range estudiantes {
				println("nombre: " + est.Nombre)
				println("matricula: " + strconv.Itoa(est.Matricula))
				println("correo: " + est.Correo)
				println("carrera: " + est.Carrera)
				println("--------------------------------------------------------------------------------")
			}
		case '2':
			print("Digite el numero de matricula del estudiante")
			var mat int
			_, _ = fmt.Scanf("%d\n", &mat)
			est := getEstudiante(mat)
			if est.Matricula == 0 && est.Carrera == "" && est.Correo == "" && est.Nombre == "" {
				println("No se encontró ese estudiante")
			} else {
				println("nombre: " + est.Nombre)
				println("matricula: " + strconv.Itoa(est.Matricula))
				println("correo: " + est.Correo)
				println("carrera: " + est.Carrera)
			}
		case '3':
			var newEst estudiante
			print("Digite el nombre del estudiante")
			_, _ = fmt.Scanf("%s\n", &(newEst.Nombre))
			print("Digite el correo del estudiante")
			_, _ = fmt.Scanf("%s\n", &(newEst.Correo))
			print("Digite la carrera del estudiante")
			_, _ = fmt.Scanf("%s\n", &(newEst.Carrera))
			if crearEstudiante(newEst) {
				println("operacion exitosa")
			} else {
				println("ERROR")
			}
		case 'q':
			doIt = false
		default:
			println("Digite una opcion valida e intentelo de nuevo (pulse cualquier tecla para continuar)")
			//_, _ = fmt.Scanf("%c", &i)
		}
		correct := false
		for !correct {
			if doIt {
				println("Desea hacer otra operación? [Y/n]")
				var inp rune
				_, _ = fmt.Scanf("%c\n", &inp)
				if unicode.ToUpper(inp) == 'N' {
					doIt = false
					correct = true
				} else if unicode.ToUpper(inp) != 'Y' {
					println("Digite una opcion valida")
				} else {
					correct = true
				}
			}
		}

	}

	/*estudiantes := getEstudiantes()
	for _, est := range estudiantes {
		println("nombre: " + est.Nombre)
		println("matricula: " + strconv.Itoa(est.Matricula))
		println("correo: " + est.Correo)
		println("carrera: " + est.Carrera)
	}
	mat := 20160138
	stud := getEstudiante(mat)
	println("nombre: " + stud.Nombre)
	println("matricula: " + strconv.Itoa(stud.Matricula))
	println("correo: " + stud.Correo)
	println("carrera: " + stud.Carrera)

	var newEstud estudiante
	newEstud.Matricula = 20160138
	newEstud.Nombre = "José Ureña"
	newEstud.Correo = "polquefuequemedicuenta@pablopiddy.do"
	newEstud.Carrera = "ISC"
	fmt.Println(crearEstudiante(newEstud))*/
}
