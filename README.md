# Web app de catálogo de tarjetas gráficas

**NOTA**: Por problemas de salud, el proyecto se actualizará más lento de lo previsto.

### ¿Qué es y qué motivaciones hay en el proyecto?

Como su nombre indica es una simple aplicación web consistente en un breve catalógo de tarjetas gráficas. Debido a que no lo considero un proyecto muy ambicioso, el backend sólo se emplea para crear un API con un end point, la cual generar un archivo en JSON. Este es leído y parseado en Angular para mostrar cada tarjeta como elemento. 

La motivación es, evidentemente, el aprendizaje como introducción al desarrollo Full stack.

### Uso

Actualmente sólo está disponible el backend. Accede a la carpeta del mismo nombre y desde el terminal ejecuta el siguiente comando:

`$ go run main.go`

Si se quisiera generar un ejecutable con extensión .exe, ejecuta en el mismo directorio:

`$ go build`

En ambos casos se crea un servidor en el puerto 8000 con un único endpoint: `http://localhost:8000//graphic-cards`

<br>


***Nota***: El contenido de este apartado se actualizará cuando el front-end esté operacional.

### Tecnologías usadas

**Back-End**
- Golang: Creación de un simple Rest Api con un endpoint. 

**Front-End**
- Angular 2+
- SASS (*No implementado actualmente*)
- Material UI

**Otras tecnologías**
- BEM (Jerarquización de uso en clases CSS)
- Eslint (Linter JS)

