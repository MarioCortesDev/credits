# Requerimientos

1.- Tener el servicio de docker escuchando

2.- Postman

# Pasos para probar

## Montar primero la imagen en de la BD en docker

Ejecutar en la ruta de credits/database los siguientes comando

```docker build . -t credit-db```

```docker run -p 54321:5432 credit-db```

## Ejecución del programa

Ejecutar en la raíz del protecto la siguiente intrucción

```go run main.go``` 

## Probar en postman

Los 2 endpoints trabajan en POST

```localhost:8080/credit-assignment```

Colocar el cuerpo raw con el siguiente formato json
{
    "investment": 3000
}

En caso de ser positivo retornara 200 con una de las combinaciones posibles para el calculo del prestamo, caso contrario retornará un mensaje con un 400

Para el sigiente endpoint

```localhost:8080/statistics```

Retornará la información de:
 Total de asignaciones realizadas (e.g. 100)
 
 Total de asignaciones exitosas (e.g. 40) 
 
 Total de asignaciones no exitosas (e.g. 60)
 
 Promedio de inversión exitosa (e.g. 3545.6)
 
 Promedio de inversión no exitosa (e.g. 350.3)
