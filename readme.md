# Resolvedor de Sopa de Letras
## Descripción
Este programa resuelve sopas de letras. Dada una sopa de letras y una lista de palabras, el programa debe encontrar todas las palabras en la sopa de letras.
- Las palabras pueden estar en cualquier dirección (horizontal, vertical, diagonal) y pueden estar escritas de izquierda a derecha, de derecha a izquierda, de arriba a abajo, de abajo a arriba, en diagonal de izquierda a derecha, en diagonal de derecha a izquierda, en diagonal de arriba a abajo o en diagonal de abajo a arriba. 
- Las palabras pueden estar en cualquier posición de la sopa de letras.
- Las palabras pueden solaparse. 
- El programa puede devolver cada posición de las letras de la palabra.

## Uso
El programa se ha de compilar con el siguiente comando:
```
go build main.go
```

Y se ha de ejecutar con el siguiente comando:
```
./main <sopa_de_letras> <lista_de_palabras> [<paralelización=[Y|N]>] [<salida=[Y|N]>]
```

## Ejemplo
Dada la siguiente sopa de letras:
```
A B C D E
F G H I J
K L M N O
P Q R S T
U V W X Y
```

Y la siguiente lista de palabras:
```
ABCD
EFGH
IMQU
MNOP
````

El programa debe devolver:
```
ABCD: (0,0) (0,1) (0,2) (0,3)
EFGH: Not found
IMQU: (1,3) (2,2) (3,1) (4,0)
MNO: (2,2) (2,3) (2,4)
```