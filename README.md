# FileWatcher
## Watcher para actualización de Firestore

Éste código desarrollado en Golang permite observar los cambios en un directorio o archivo y emitir alertas a la base de datos en caso de el estímulo sea una escritura. Para poder hacer ésto es necesario utilizar una serie de parámetros a la hora de ejecutar la rutina (en el orden dado a continuación y separados por espacios): 

- Path del directorio o archivo que desea ser observado
- Usuario de el repositorio en la nube dónde se almacenarán los datos
- Contraseña de el repositorio en la nube dónde se almacenarán los datos
- El identificador de la cámara de la que se está grabando
- El identificador del usuario dueño de la cámara (Opcional)

Para tener acceso a la instancia de la base de datos que desea utilizar deberá crear un archivo .env que contenga la ruta al archivo de acceso. Dónde colocar éste archivo dependerá de la opción que escoja a la hora de correr el código:

- Si decide correr el código haciéndo uso del comando `go run main.go` desde el directorio main, deberá colocar el archivo dentro de ésta carpeta.
- Si decide buildear el proyecto generando un ejecutable, el archivo deberá ser generado en el mismo directorio en el que se encuentre el ejecutable.

La variable que debe contener el archivo se denomina ***FIRESTORE***. En forma general puede definirse como `FIRESTORE = <path>/<to>/<file>`.

## Instalación

Dentro del directorio raíz del proyecto encontrará la carpeta **main**. Dentro de ella, abra una consola. Puede realizar dos acciones en éste punto:

1) Utilizar el comando `go run main.go` para iniciar la ejecución del código. Recuerde ingresar los parámetros mencionados en el inicio de éste documento en el orden indicado. En caso contrario ocurrirá un error. 
2) Utilizar el comando `go build` para generar un ejecutable que luego será corrido.
