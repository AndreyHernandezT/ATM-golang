# Simulador de cajero en GO
_Simulador de Cajero ATM desarrollado en GO, tomando (y actualizando) valores de archivo JSON_

## ¿Cómo se usa? 🚀
### Para ejecutar el programa

```bash
go run atm.go
```
### Mensaje de bienvenida
```
Bienvenido al cajero
Por favor, ingresa tu numero de cuenta
```
El número de cuenta se encuentra en el archivo JSON `database.json` de la siguiente manera `"account_number":"__"`

Una vez ingresado el número de cuenta, se mostrará el menú principal

### Menú principal

```
Hola de nuevo,  user_name

¿Qué deseas realizar?
1. Consultar saldo
2. Retirar
3. Ingresar dinero
4. Transferir
5. Salir
```

Al ser un cajero, por cada operación pedirá la contraseña del usuario específico, también en el archivo `database.json` como `"password":"__"`

### 1. Consultar saldo
Para consultar saldo ingresa el número 1. Como se mencionó anteriormente, por cada operación se pedirá la contraseña.

Se valida la contraseña y si es correcta, devolverá el saldo del usuario, leido del archivo `database.json` 
```
1
Por favor, ingresa tu contraseña:
___

Tienes un saldo de: $ ______
```

### 2. Retirar
El proceso es igual, pide la contraseña y una vez validada, pedirá ingresar la cantidad a retirar. Si lo que se desea retirar es menor o igual al saldo correspondiente, se restará dicha cantidad al saldo del usuario en el archivo `database.json`  y mostrará el saldo actual
```
2
Por favor, ingresa tu contraseña:
___

Por favor, ingresa la cantidad a retirar
___

Nuevo saldo:  ___

```
### 3. Ingresar dinero
Se ingresa la contraseña del usuario y posteriormente el dinero que desea consignar en la cuenta. El dato se actualiza en el archivo `database.json` y muestra el saldo actual

```
3
Por favor, ingresa tu contraseña:
___

Por favor, ingresa la cantidad a ingresar:
___

Nuevo saldo:  ___
```

### 4. Transferir
Se ingresa la contraseña del usuario y posteriormente el **número de la cuenta del usuario a transferir**, en caso de ser el mismo número de cuenta del usuario actual, se vuelve a preguntar el número de cuenta. Se muestra el nombre del usuario al que desea transferir y le preguntará si es correcto, se pregunta la cantidad de dinero a transferir y muestra en pantalla el saldo actual

```
4

Por favor, ingresa tu contraseña:
___

Ingresa numero de la cuenta a transferir:
___

Transferir a:  user_name  ¿correcto?
1. Si
2. No
1

¿Cuanto vas a transferir?
___

Transferencia Exitosa
Nuevo saldo:  ___
```

### 5. Salir
Se despide del usuario y finaliza el programa

```
5

¡Hasta pronto!
```

---
⌨️ con ❤️ por [Andrey](https://github.com/AndreyHernandezT) 😊