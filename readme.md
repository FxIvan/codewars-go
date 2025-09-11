| Method | Description |
| --- | --- |
| len() | Devuelve longitud del string |
| for i := 0; i < 10; i++ {} | for sobre elementos de un array o objectos |
| for i, r := range str{} | for sobre elementos de un array o objectos, i representa al indice y r el elemento que se itera |
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
| strings.Builder | Cuando concatenamos con "+", Go crea una nueva cadena en memoria cada vez. Por ejemplo, "hola" + " buen día" crea "hola buen día" en memoria; si después concatenamos "hola" + " buenas tardes", se genera una segunda cadena en memoria. Esto es ineficiente si necesitamos muchas concatenaciones (como en un libro), ya que crea múltiples copias en memoria. strings.Builder usa un buffer, que permite reutilizar el mismo espacio de memoria para añadir varios elementos sin crear una nueva cadena cada vez. |
|strings.HasPrefix("codecademy","code") -> true  strings.HasPrefix("GitHub","Hub") -> false | La función HasPrefix() devuelve un valor booleano que indica si una cadena dada comienza con un prefijo determinado. El método devuelve verdadero si la cadena comienza con el prefijo; de lo contrario, devuelve falso. |
|Agregar validaciones al inicio. Ej: 	if len(grafo) == 0 { return totalPeso // Retorna 0 si el grafo está vacío } | Evita errores durante la ejecucion |
|math.MinInt| Para no tener que inicializar con 0, Incorrecto: maxPeso := 0; Correcto: maxPeso := math.MinInt |


### Problema
```
| panic: runtime error: invalid memory address or nil pointer dereference [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x48a930] goroutine 1 [running]: github.com/codewars peteTheBaker.Cakes(0x496240?, 0xc00009ce10) /home/almendra/codewars/peteTheBaker/cackes.go:56 +0x210 main.main() /home/almendra/codewars/main.go:96 +0x205 exit status 2 
```
### Solucion
Se debe verificar que las variables sean diferente a nil si se utiliza punteros. Paso en el archivo cackes.go|


[NOTAS](https://docs.google.com/document/d/1fWRgd2q1P5VzxhLDA0Aw--kVTLPxS81V8kpMmKM_-wc/edit?usp=sharing)

