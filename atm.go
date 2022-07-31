package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Users struct { //Creación Estructura para recibir información
	Username      string `json:"username"`
	AccountNumber string `json:"account_number"`
	Password      string `json:"password"`
	Balance       int    `json:"balance"`
}

func getUsers() []Users { //Se obtienen los usuarios
	users := make([]Users, 2)                    //Lista de 2 mapas vacios estructura Users
	raw, err := ioutil.ReadFile("database.json") //Lee archivo
	if err != nil {
		fmt.Println(err.Error())
	}
	json.Unmarshal(raw, &users) //Decodificar contenido JSON dentro instancia users
	return users                //Devuelve lista de los usuarios en el JSON
}

func findUser(account_number string) (Users, bool) { //Encuentra al usuario de acuerdo al numero de cuenta
	user_not_found := Users{}    //Crea estructura en caso de n encontrar usuarios
	users := getUsers()          //Llama a todos los usuarios
	for _, user := range users { //Itera usuarios
		if user.AccountNumber == account_number {
			return user, true //Devuelve al usuario con la bandera flag
		}
	}
	return user_not_found, false
}

//Actualiza el valor del saldo del usuario con numero cuenta
func newBalance(account_number string, new_balance int) {
	var list []Users    //Lista usuarios
	users := getUsers() //Obtiene todos usuarios

	for _, user := range users {
		if user.AccountNumber == account_number {
			user.Balance = new_balance //Actualiza el valor
			//fmt.Println("Usuario: ", user)
			list = append(list, user)
		} else {
			list = append(list, user)
		}
	}
	//fmt.Println("List: ", list)
	overwriteBalance(list) //sobreescrive el JSON con nuevo saldo
}

//Funcion sobreescribir JSON
func overwriteBalance(list []Users) {
	content, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("database.json", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func verifyUser() { //Verifica número de cuenta del usuario
	var account_number string
	fmt.Scan(&account_number)

	user, flag := findUser(account_number)
	if flag {
		fmt.Println("\nHola de nuevo, ", user.Username)
		menu(user) //Llama al menu con el usuario especifico
	} else {
		fmt.Println("Usuario no encontrado, intentalo de nuevo mas tarde")
	}
}

func seeBalance(user_access Users) { //Devuelve Balance Usuario
	user, flag := findUser(user_access.AccountNumber) //Consulta JSON
	if flag {
		fmt.Println("\nTienes un saldo de: $", user.Balance)
	}
}

func amountWithdraw(user_access Users) {
	var amount int
	fmt.Println("\nPor favor, ingresa la cantidad a retirar:")
	fmt.Scan(&amount)

	if user_access.Balance-amount > 0 { //Verifica cantidad a retirar
		user, flag := findUser(user_access.AccountNumber) //Consulta JSON
		if flag {
			new_balance := user_access.Balance - amount

			newBalance(user.AccountNumber, new_balance) //Actualiza saldo
			fmt.Println("\nNuevo saldo: ", new_balance)
		}
	} else {
		println("No es posible retirar la cantidad de dinero especificada")
	}
}

func amountToDeposit(user_access Users) {
	var amount int
	fmt.Println("\nPor favor, ingresa la cantidad a ingresar:")
	fmt.Scan(&amount)

	user, flag := findUser(user_access.AccountNumber)
	if flag {
		new_balance := user_access.Balance + amount

		newBalance(user.AccountNumber, new_balance)
		fmt.Println("\nNuevo saldo: ", new_balance)
	}
}

func amountToTransfer(user_access Users) {
	var account_number string
	fmt.Println("\nIngresa numero de la cuenta a transferir:")
	fmt.Scan(&account_number) //Pregunta numero cuenta a transferiri

	user, flag := findUser(account_number)

	//Si encuentra al usuario y es diferente a si mismo
	if flag && user_access.AccountNumber != user.AccountNumber {
		var option string
		fmt.Println("\nTransferir a: ", user.Username, " ¿correcto?")
		fmt.Println("1. Si\n2. No")
		fmt.Scan(&option)
		switch option { //Pregunta si quiere transferir a ese usuario
		case "1":
			var amount int
			fmt.Println("\n¿Cuanto vas a transferir?")
			fmt.Scan(&amount)
			if user_access.Balance-amount > 0 { //Verifica cantidad a transferir
				new_balance_access := user_access.Balance - amount
				newBalance(user_access.AccountNumber, new_balance_access) //Actualiza saldo a tranferir

				new_balance_transfer := user.Balance + amount
				newBalance(user.AccountNumber, new_balance_transfer) //Transfiere saldo
				fmt.Println("\nTransferencia Exitosa\nNuevo saldo: ", user_access.Balance-amount)
			} else {
				println("No es posible transferir la cantidad de dinero especificada")
			}
		case "2":
			amountToTransfer(user_access)
		}
	} else {
		fmt.Println("No encontramos la cuenta, intentalo de nuevo más tarde")
	}
}

func verifyCredentials(user_access Users, option_case string) { //Verifica contraseña usuario
	var password_scan string
	fmt.Println("\nPor favor, ingresa tu contraseña:")
	fmt.Scan(&password_scan)

	if user_access.Password == password_scan {
		switch option_case { //Realiza acción dependiendo caso
		case "1":
			seeBalance(user_access)
		case "2":
			amountWithdraw(user_access)
		case "3":
			amountToDeposit(user_access)
		case "4":
			amountToTransfer(user_access)
		}
	} else {
		fmt.Println("Contraseña incorrecta, intentalo de nuevo") //Vuelve a preguntar contraseña
		verifyCredentials(user_access, option_case)
	}
}

func menu(user_access Users) { // Funcion imprimir Menu
	var option string

	fmt.Println("\n¿Qué deseas realizar?")
	fmt.Print("1. Consultar saldo\n")
	fmt.Print("2. Retirar\n")
	fmt.Print("3. Ingresar dinero\n")
	fmt.Print("4. Transferir\n")
	fmt.Print("5. Salir\n")

	fmt.Scan(&option)

	switch option {
	case "1":
		verifyCredentials(user_access, "1")
		menu(user_access)
	case "2":
		verifyCredentials(user_access, "2")
		menu(user_access)
	case "3":
		verifyCredentials(user_access, "3")
		menu(user_access)
	case "4":
		verifyCredentials(user_access, "4")
		menu(user_access)
	case "5":
		fmt.Println("\n¡Hasta pronto!")
	}
}

func main() {
	fmt.Print("\nBienvenido al cajero\n")
	fmt.Println("Por favor, ingresa tu numero de cuenta")
	verifyUser()
}
