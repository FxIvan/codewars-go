| Method | Description |
| --- | --- |
| len() | Devuelve longitud del string |
| for i := 0; i < 10; i++ {} | for sobre elementos de un array o objectos |
| for i, r := range str{} | for sobre elementos de un array o objectos |
| for i, r := range str{} | for sobre elementos de un array o objectos |
| strings.Join(array, sep) | sep es el separador que se coloca entre los elementos en la cadena final.|
| !true | en GO tambien se puede utilizar "!" para indicar el valor contrario a un valor boleano |
| uri[len(uri)-3):] | exluye los primeros 3 caracteres de la cadena, por el 'array[condicion:']|
| proxy := method.NewProxyValue() | dentro de proxy hay funciona que cree |
| make([]string,len(arr)) | make() permite armar un array, el primer parametro es el tipo de array y el segundo la longitud de ese array  |
| copy(newArray,arr) | copy(arrayNuevo,arrayACopiar), se utiliza para copiar el array a una nueva variable, en este caso va hacer arrayNuevo |
| package bufio | El paquete bufio implementa E/S almacenadas en búfer. Envuelve un objeto io.Reader o io.Writer, creando otro objeto (Reader o Writer) que también implementa la interfaz pero proporciona almacenamiento en búfer y algo de ayuda para E/S textuales. Bufer es  es una región de la memoria utilizada para almacenar temporalmente datos mientras se transfieren entre dos lugares o se procesan|
| os.Stdin  | Se utiliza para leer datos proporcionados por el usuario o por otro programa a través de la consola o la terminal.|
| os.Args   | os.Args es una variable que proporciona acceso a los argumentos de la línea de comandos que se pasan al programa cuando se ejecuta. Es un slice de strings ([]string) que contiene el nombre del programa en su primera posición, seguido de los argumentos que el usuario haya pasado al ejecutarlo.|
| bufio.NewScanner | El escáner es particularmente útil para leer datos línea por línea o palabra por palabra (tokens) de forma sencilla y eficiente. |


