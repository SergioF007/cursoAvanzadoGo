# cursoAvanzadoGo

## ****Cómo aprender Go avanzado: concurrencia, patrones de diseño y net****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/a32877f3-0ed7-48c1-858e-56a56335d26b/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/cbfbd63e-ba2c-495f-aec9-f6af93ebcb40/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/42f2795b-b932-4c7a-a4fb-5d475128570a/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/4652d7d4-8d6d-470f-a688-dea18592a32a/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/f7f7dd35-92dd-41f5-b8d4-5c5f50bd6c42/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/540a1444-0ff5-4b7a-839f-92b1882c7fa4/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d27fa65d-0dd7-43f6-b9df-235fe4a0ebc0/Untitled.png)

## ****Race condition: el problema de depositar y retirar****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/a47afa21-5b1b-41b2-9e21-95a988e4c0ca/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/beb4fd4d-7eef-4337-ba64-64cd6b5100ed/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/044581e2-6c18-45fe-90e0-2a5990f29604/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/df8be90f-be85-4c29-bd43-0807296224b3/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/2a852780-3124-4f37-b56b-91bd9dd8b9fb/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/94ec89fd-e6db-4445-9a1b-3f60b6ce5954/Untitled.png)

```go
package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	b := balance
	balance = b + amount
}

func Balance() int {
	b := balance
	return b
}

func main() {

	// vamos a representar diferentes depositos que estan ocurriendo en el balance en el mismo periodo de tiempo.

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg)

	}

	wg.Wait()
	fmt.Println(Balance())

}
```

## ****Sync Mutex: Lock y Unlock****

**para poder ver si lo que estamos creando tiene riegos de condición de carrera  ya que estamos implementado diferentes subrutinas (goroutine).  que están accediendo a las mismas variables.** 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/40ddc9af-73c1-4638-a5bc-c38eac9c195e/Untitled.png)

**Estos son las advertencias que me indica:** 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/900ef980-272c-4c92-8718-1e2902ba2fc0/Untitled.png)

```go
package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	lock.Lock() // en el momento donde alguien este escribiendo datos en esta funcion, vamos a bloquiar el acceso a las mismas dado a que esta en un proceso.
	b := balance
	balance = b + amount
	lock.Unlock() // y se desbloque despues de ejecutar lo que ocurrer en las dos lineas anteriores.

}

func Balance() int {
	b := balance
	return b
}

func main() {

	// vamos a representar diferentes depositos que estan ocurriendo en el balance en el mismo periodo de tiempo.

	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)

	}

	wg.Wait()
	fmt.Println(Balance())

}
```

corregimos los errores de carrera utilizando el **concepto de candado**:   con **lock.Lock()** y **lock.Unlock()**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/aa129288-12e5-4804-8ed4-6da0dc99a503/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/37aba1f1-8134-405e-aeb4-3a6d0200f99d/Untitled.png)

### Nota: Es verdad que el lock genera un bloqueo, sin embargo si puedes notar, son bloqueos que podemos establecer en lineas específicas del código, (la idea no es bloquear toda la gorutine). Es decir, no es lo mismo ejecutar e una sola rutina dos depositos que crear 2 goroutines para cada deposito y bloquear solamente las lineas en donde se modifica la información que comparten entre ellas. Sigue siendo mucho mas eficiente hacer uso de la concurrencia en este tipo de escenarios. Lo que sucede es que por la sencilles del ejemplo no se percibe tanto la ventaja de usar las gorutines y el mutex al mismo tiempo, puesto que la idea de Nestor es que entendamos con un ejemplo sencillo como podemos enfrentar esta condición por medio del Lock() y Unlock()

## ****Mutex de lectura y escritura****

Solo para aclarar las diferencias entre RWLock y no usar nada:

- Lock bloquea lecturas (con RLock) y escrituras (con Lock) de otras goroutines
- Unlock permite nuevas lecturas (con Rlock) y/o otra escritura (con Lock)
- RLock bloquea escrituras (Lock) pero no bloquea lecturas (RLock)
- RUnlock permite nuevas escrituras (y también lecturas, pero por la naturaleza de RLock, estas no se vieron bloqueadas nunca)

En esencia, RLock de RWLock garantiza una secuencia de lecturas en donde el valor que lees no se verá alterado por nuevos escritores, a diferencia de no usar nada.

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/78be9436-9c78-4485-96cf-51d3fe973f4c/Untitled.png)

como podemos ver en este ejemplo, que como solo necesitamos en un deposito garantizar que para que comience una acción de deposito, la acciones que ejercen cambios en estas variables terminen su lógica antes de que comience de nuevo una acción de deposito. Por lo tanto esta función no la necesito modificar. el uso de **Lock() y Unlock() es correcto.** 

Pero para el caso de Balance su lógica es distinta,  Por lo que al bloquearlos con **Lock() y Unlock()** no permito ni la **Escritura y Lectura.** Quiere decir que mientras se este ejercicio alguna acción o cambios sobre su variable de lectura. No podrá ingresar a la info,  para leerla mientras alguien se encuentre ejerciendo un deposito. Por lo tanto no estaríamos resolviendo el erro de dejarla como estaba sin el uso de **Lock() y Unlock() en Balance().** Por que si no puedes a acceder a cambios de información actual y que el estado final dependa de que complete cierto bloque de código mas no solo una parte entonces puedo tomar un dato que no corresponda. Debido que no le estaríamos dando restricciones a Balance. Por lo tanto aplicamos restricciones respecto a las acciones    

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/100000a9-d9c6-45e6-98ae-559bb0d88dfd/Untitled.png)

Le estamos dando permisos de lectura mientras su permiso de escritura se encuentra bloqueado. **Por lo que estamos permitiendo tener diferentes lectores sin que sean bloqueados.** 

## ****Sistema de caché sin concurrencia****

```go
package main

import (
	"fmt"
	"time"
)

func Fibonacci(n int) int {

	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)

}

type Memory struct {
	f     Function
	cache map[int]FunctionResult // el cahce me va almacenar todas los Key  y sus resultados que se calculen en la funcion Function
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

// Constructor del Memory
func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

// Metodo Get de Memory
// va a recibirnos el entero y no va a devolver si ya existe un resultado para este entero en el cache
// si no existe lo va a calcular.

func (m *Memory) GetMemory(key int) (interface{}, error) {
	result, exists := m.cache[key] // consulta en el cache

	if !exists {
		result.value, result.err = m.f(key) // calcula el fibonacci
		m.cache[key] = result               // lo almaceno en el cache
	}

	return result.value, result.err
}

// usamos el interface{} para indicar que vamos a tener un respueta genrica
// osea que no sabemos con exactitud el tipo de resultado que vamos a obtener
func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {

	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 38}

	for _, n := range fibo {

		start := time.Now()
		value, err := cache.GetMemory(n)

		if err != nil {

			fmt.Println(err)

		}
		fmt.Printf("%d, %s, %d\n", n, time.Since(start), value)

	}

}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/c87492ee-94e8-4799-9572-dcdc39112536/Untitled.png)

vemos como al volver a ejecutar el fibonacci al input 42, como es un Key que ya esta almacenado en cache, es muy rápido la respuesta. por lo que ya ha sido calculada y alcenada en cache. 

## ****Sistema de caché con concurrencia****

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func Fibonacci(n int) int {

	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)

}

type Memory struct {
	f     Function
	cache map[int]FunctionResult // el cahce me va almacenar todas los Key  y sus resultados que se calculen en la funcion Function
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

// Constructor del Memory
func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

// Metodo Get de Memory
// va a recibirnos el entero y no va a devolver si ya existe un resultado para este entero en el cache
// si no existe lo va a calcular.

func (m *Memory) GetMemory(key int) (interface{}, error) {
	result, exists := m.cache[key] // consulta en el cache

	if !exists {
		result.value, result.err = m.f(key) // calcula el fibonacci
		m.cache[key] = result               // lo almaceno en el cache
	}

	return result.value, result.err
}

// usamos el interface{} para indicar que vamos a tener un respueta genrica
// osea que no sabemos con exactitud el tipo de resultado que vamos a obtener
func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {

	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 38}

	var wg sync.WaitGroup

	for _, n := range fibo {

		wg.Add(1)

		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, err := cache.GetMemory(index)

			if err != nil {

				fmt.Println(err)

			}
			fmt.Printf("%d, %s, %d\n", index, time.Since(start), value)
		}(n)

	}

	wg.Wait()

}
```

### Nota: Vemos que a aplicar concurrencia no esta haciendo uso del cache como podemos ver en los tiempos de respuesta para el caso del input 42, es como si se estuviese re calculando su resultado, por lo que esta tardando un tiempo similar en ambos casos. Así que vamos analizar nuestras advertencias de **condición de carrera para que ver que nos dice.**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/1983c68e-0cbf-4fba-a858-689ac5fa0d7e/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/2bb16c14-cc02-4829-a430-9aa2d24fa84b/Untitled.png)

al corregir los posibles errores de carrera vemos que de igual forma los tiempos de respuestas no son óptimos, por lo que en este caso no seria recomendable aplicar la concurrencia, ya que sin ella obteníamos respuestas de tiempo mas cortos. 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/b390e2d2-d31e-4cb7-bc09-6f22520de294/Untitled.png)

## ****Reutilización de computación intensiva****

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func ExpensiveFibonacci(n int) int {

	fmt.Printf("Calculate Expensive Fibonacci for %d\n", n)
	time.Sleep(5 * time.Second)

	return n
}

type Service struct {
	InProgress map[int]bool       //almacena los numeros de los caules vamos a calcular la serie fibonacci he idicar en que estado estan
	IdPending  map[int][]chan int // vamos a mapiar las llaves de tipo entero a un slice de canales
	Lock       sync.RWMutex
}

func (s *Service) Work(job int) {

	s.Lock.RLock()
	exists := s.InProgress[job] // esto nos va a indicar si el job esta siendo procesado o no
	if exists {

		s.Lock.RUnlock()
		response := make(chan int)
		defer close(response)

		s.Lock.Lock()
		s.IdPending[job] = append(s.IdPending[job], response)
		// asi que por el canal de response en por donde se le va a cuminicar
		// al worket exists que ya se a terminado de calcular la serie de fibonacci
		s.Lock.Unlock()

		fmt.Printf("Waiting for Response job: %d\n", job)
		resp := <-response
		fmt.Printf("Response Done, received %d\n", resp)

		return
	}

	// cuando aun no esta en progreso, Generamos el bloqueo para empezar el proceso
	s.Lock.RUnlock()

	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()

	fmt.Printf("Calculate Fibonacci for %d\n", job)
	result := ExpensiveFibonacci(job)

	// ya que calculamos la serie fibonacci guadamos su resultado y lo almacenamos en result
	// para consultar los datos vamos, bloquiamos y habilitamos la lectura
	s.Lock.RLock()
	// traemos pos Workers que estaba esperando los resultados que la funcion a calculado y si existen o no
	pendingWorkers, exists := s.IdPending[job]
	s.Lock.RUnlock()

	if exists {
		for _, pendingWorker := range pendingWorkers {
			pendingWorker <- result // le notificamos a todos los workers que estan pendiente que su resultado ya a sido calculado y se les esta enviando
		}

		fmt.Printf("Result sent - all peding workers ready job: %d\n", job)

	}

	s.Lock.Lock()
	// los estoy configurando a su estado inicial
	s.InProgress[job] = false // va hacer falso por que ya fue calculado
	s.IdPending[job] = make([]chan int, 0)
	s.Lock.Unlock()

}

// Creamos el constructor
func NewService() *Service {

	return &Service{
		InProgress: make(map[int]bool),
		IdPending:  make(map[int][]chan int),
	}
}

func main() {

	service := NewService()
	jobs := []int{3, 4, 5, 5, 4, 8, 8, 8}
	var wg sync.WaitGroup
	wg.Add(len(jobs)) // le agrego al contador la longitud de los jobs que tenemos
	for _, n := range jobs {
		go func(job int) {

			defer wg.Done()
			service.Work(job)

		}(n)
	}

	wg.Wait()

}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/57a68d69-32cd-4ff1-ad2a-fd3604f3abec/Untitled.png)

revisa que todas la función ya hayan sido ejecutadas y sino tambien se pregunta si esa que pretendo ejecutar ya esta en proceso de ejecución y si es ha si tengo que generar una respuesta cuando la ejecución ya halla sigo terminada y me la envié como parámetro.  Por lo tanto no comenzaría la ejecución por que ya una acción que se a empezado para hallar ese calculo.    

### Nota: Estudiarlo para dar mis conclusiones

### Tarea: Implementar estos nuevos conceptos al ejercicio de ****Sistema de caché con concurrencia para mejorar velocidades de respuesta.****

## ****¿Qué son los patrones de diseño?****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/c9beb191-1cf7-42ea-8cbf-ba140669a762/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/78833936-70a0-4e48-a7fd-00797fc2bbe7/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/c4e6d84f-329b-4942-97c3-01e1581796e0/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/3824758c-685d-46ca-9fb5-38c0ded5e9d4/Untitled.png)

## ****Factory****

Este patrón de diseño no se preocupa por cuantos subtipos hay  ! puedo tener cualquier cantidad de subclases que están re utilizando atributos y métodos en común. 

```go
package main

import "fmt"

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

// Vamos a aplicar composicion para representar una herencia de un clase que tiene relacion con Computer

type Laptop struct {
	Computer
}

// ya que estamos construyendo la herencia con composicion, y asi solo me permite compartir los atributos
// para representar la herencia de metodos en go vamos a crear un funcion que me retorne la interfaz que tambien implementad Computer.
// Por lo que tenemos que aplicar composicion para poder representar la herencia de metodos
func newLaptop() IProduct {

	return &Laptop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

// Creamos otra subclase

type Desktop struct {
	Computer
}

func newDesktop() IProduct {

	return &Desktop{
		Computer: Computer{
			name:  "Desktop Computer",
			stock: 35,
		},
	}
}

func GetComputerFactory(computerType string) (IProduct, error) {

	if computerType == "laptop" {
		return newLaptop(), nil
	}

	if computerType == "desktop" {
		return newDesktop(), nil
	}

	return nil, fmt.Errorf("Invalid computer type")
}

// Creo la funcion que me va imprimir la informacion de la supclase de interes.
func printNameAndStock(p IProduct) {
	fmt.Printf("Product name: %s, with stock %d\n", p.getName(), p.getStock())
}

func main() {

	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")

	printNameAndStock(laptop)
	printNameAndStock(desktop)
}
```

### Nota:  Comparativa con Java y Go en el aspecto de la herencia.

Como Java en el caso particular de la herencia, No soporta el tipo de herencia múltiple, esto quiere decir que una clase Hija solo puede heredar de una clase padre. En comparacion con Go, como no maneja manaje la herencia y en su lugar usa la composición para presentar  la herencia, Soporta que un struc al implementar hacer uso de la composición se pueda componer de distintos tipos de struc . 

**Esta la consulta realizada a chapGPT:** 

pero en go al representar la herencia por composición estaría permitiendo una similitud con la herencia múltiple ?

**Respuesta de Chat:** 

Exacto, en Go no hay herencia en el sentido tradicional de la programación orientada a objetos, pero la composición y la incrustación de tipos permiten que los diseñadores de software en Go alcancen efectos similares a la herencia múltiple. Al "incrustar" un tipo dentro de otro, efectivamente se importan todos los métodos del tipo incrustado, lo que da una especie de "herencia de comportamiento".

Por ejemplo:

```go

go
Copy code
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

// ReaderWriter tiene tanto Read como Write métodos
type ReaderWriter interface {
    Reader
    Writer
}

// O también podrías hacer esto con structs en lugar de interfaces
type Engine struct {
    // ...
}

func (e *Engine) Start() {
    // ...
}

type Car struct {
    Engine  // esto es incrustación
}

// Ahora puedes hacer
c := &Car{}
c.Start()  // Este método se hereda de Engine

```

Así que aunque Go no tiene "herencia" en el sentido tradicional, las capacidades de composición y de incrustación de tipos permiten un alto grado de reutilización de código, que es uno de los objetivos principales de la herencia en la programación orientada a objetos.

## ****Singleton****

Es bastante común verlo en diferentes situaciones. 

EJ: Caso de conexión con una DB es el caso de usa donde es mas implementado. 

Donde tenemos una conexión a una DB y la estamos instanciando en un clase. 

Por que tenemos nuestra clase de DB   —>  Se crea el objeto de conexión.  Pero que paso si volvemos a crear otro objeto de conexión ?  tendríamos un total de dos instancias de conexión.

Pero este comportamiento no siempre lo estamos buscando y en su lugar hay veces que solo necesitamos manejar una sola instancia de conexión, **Como es el caso de la conexión a un DB donde es mas recurrente implementar el Singlenton**. 

 

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type DataBase struct{}

// Vamos a crear la funcion de conexion de la DataBase
func (DataBase) CreateSingleConnection() {
	fmt.Println("Creating Singleton for DataBase")
	time.Sleep(2 * time.Second)
	fmt.Println("Creation Done")
}

var db *DataBase
var lock sync.Mutex

func getDataBaseIntance() *DataBase {
	lock.Lock()
	defer lock.Unlock()
	// estamos validadno si ya hay una instancia creada, por lo que si no la hay, sera igual a nil y toca crearla
	if db == nil {
		fmt.Println("Creating DB Connection")
		db = &DataBase{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("DB Already Created")
	}

	return db

}

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDataBaseIntance()
		}()
	}
	wg.Wait()
}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/af0eb1a1-0aa9-4073-924f-49ea7d97940c/Untitled.png)

## ****Adapter****

Sirve cuando una interface esta definida de alguna manera y luego tengo un struct que no implementa la interfaz como se quisiera, por que hay un incompatibilidad. 

Lo que vamos hacer es un intermedio que nos ayude adaptar esta interfaz al comportamiento que se busca. 

**En el siguientes ejemplo vemos el error que se me produce si intento implementar un método heredado que no esta adaptado a la funcionalidad que intento incorporar.**  

```go
package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using Cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using BankAccount %d\n", bankAccount)
}

func main() {

	cash := &CashPayment{}
	ProcessPayment(cash)
	// probamos

	bank := &BankPayment{}
	ProcessPayment(bank)

}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/b4d625cc-6f43-4a99-aa73-f3c047570465/Untitled.png)

Asi que para solucionar esto es que podemos implementar **el patron de diseño Adapter** 

```go
package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using Cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

// Lo anterior es el primer compartamiento del metodo que esta implementado
// lo que hicimos es que mediante una interfaz implementamos  el metodo Pay del struc CashPayment

type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using BankAccount %d\n", bankAccount)
}

// Vamos a crear el adaptador del metodo Pay
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

// Vamos hacer que BankPaymentAdapter implemete el Pay de la manera correcta
// como vemos esta funcion espera es el adaptador. que esta construido con su objeto principal
// y el valor que va a resivir el metodo que se pretende adaptar
func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount) // aquie estariamos creando la sobreescritura del metodo
}

func main() {

	cash := &CashPayment{}
	ProcessPayment(cash)
	// probamos

	//bank := &BankPayment{}
	//ProcessPayment(bank)

	// Instanciamos y creamos el adaptador .
	bpa := &BankPaymentAdapter{
		bankAccount: 5,
		BankPayment: &BankPayment{},
	}

	ProcessPayment(bpa)

}
```

### En conclusison: Este patron no permite implementar el tipo de polimorfismo con sobre escritura de metodos.

## ****Observer****

Utilizo observadores para evitar estar preguntado de manera recurrente si el evento de interés ya ha ocurrido. Por lo tanto vamos a garantizar que el evento notifique cuando su acción ya halla concluido. 

```go
package main

import "fmt"

type Topic interface {
	register(observer Observer)
	// esto notifica a los observadores que ha ocurrido el evento que se estaba
	//esperando que ocurriera
	broadCast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

// Item -> No Disponible en primera instacia
// Item -> Avise -> HAY ITEM

// este Item struct es el objeto que representa el estado del observable
type Item struct {
	observers []Observer // es la variable que se crea para que este observando la disponibilidad del producto
	name      string
	available bool // esta variable indica cuando hay o no hay disponinlidad del item
}

// constructor de Item
func NewItem(name string) *Item {

	return &Item{
		name: name,
	}
}

// esta funcion me va a actualiar la disponibilidad del Item para poder ejecutar el
func (i *Item) UpdateAvailable() {
	fmt.Printf("Items %s is available\n", i.name)
	i.available = true
	i.broadCast()
}

// implementamos los metodos de la interfaces para que el struc de Item
func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i *Item) broadCast() {

	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

// creamos un notificador de tipo email, este va hacer un tipo de observer.
type EmailClient struct {
	id string
}

func (eC *EmailClient) updateValue(value string) {

	fmt.Printf("Sending Email - %s available from client %s\n", value, eC.id)

}

func (eC *EmailClient) getId() string {
	return eC.id
}

func main() {

	nvidiaItem := NewItem("RTX 3080")

	// vamos a crear dos observadores
	firstObserver := &EmailClient{
		id: "12ab",
	}

	secunObserver := &EmailClient{
		id: "34dc",
	}

	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secunObserver)
	nvidiaItem.UpdateAvailable()
}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/03d15d24-b1db-412b-969a-8e7e2bd8a061/Untitled.png)

## ****Strategy****

Consiste en definir una familia de algoritmos o de funciones similares, en structs separados.  De tal manera que haya una clase base que las que este utilizando y las pueda utilizar de manera indiferente. 

este condigo utiliza principios SOLID. 

```go
package main

import "fmt"

type PasswordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

type HashAlgorithm interface {
	Hast(p *PasswordProtector)
}

// creamos el constructor de PasswordProtector
func NewPassrordProtector(user string, passwordName string, hash HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: hash,
	}
}

// vamos hacer que HashAlgorithm sea intercambiable
func (p *PasswordProtector) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash

}

func (p *PasswordProtector) Hast() {
	p.hashAlgorithm.Hast(p)
}

type SHA struct{}

func (SHA) Hast(p *PasswordProtector) {
	fmt.Printf("Hashing using SHA for %s\n", p.passwordName)
}

type MD5 struct{}

func (MD5) Hast(p *PasswordProtector) {
	fmt.Printf("Hashing using MD5 fot %s\n", p.passwordName)
}

func main() {

	sha := &SHA{}
	md5 := &MD5{}

	passwordProtector := NewPassrordProtector("Nestor", "gmail pasword", sha)
	passwordProtector.Hast()
	passwordProtector.SetHashAlgorithm(md5)
	passwordProtector.Hast()
}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/7744d121-3bfe-47b9-b694-a1d35f914c6b/Untitled.png)

## ****Escaneado de puertos sin concurrencia****

Se planteo un ejercicio que pretende recorrer los puertos del 0 al 99, para saber cuales están en uso en tu servidor local.

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	for port := 0; port < 100; port++ {
		// revisar los puerdots en este rango: 1 ,2 ..4, 99
		// sitio: 1, sitio: 2, .. sitio: 99
		// si me logro conectar a uno de estos, quiere decir que el puerto esta abierto.
		// 1 -> Open,
		// Si la conexion falla
		// 2 -> Closed    quiere decir que el puerto esta cerrado

		//Vamos a realizar un conexion tcp a cada uno de estos sitios con su repectivo puerto
		// para esta conexion primero idicamos el protocolo que vamos a implemetar
		// vamos a mostrar cual es el puerto al cual nos vamos a conectar (tenemos el nombre del sitio y el puerto)
		conexion, err := net.Dial("tcp", fmt.Sprintf("%s: %d", "scanme.nmap.org", port))
		if err != nil {
			continue
		}
		conexion.Close()
		fmt.Printf("Port %d is open\n", port)

	}
}
```

### Nota: Vemos que su comportamiento en la respuesta es bastante lento y eso que solo estamos evaluando 100 ports.

## ****Escaneado de puertos con concurrencia****

```go
package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for port := 1; port < 65535; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conexion, err := net.Dial("tcp", fmt.Sprintf("%s: %d", "scanme.nmap.org", port))
			if err != nil {
				return
			}
			conexion.Close()
			fmt.Printf("Port %d is open\n", port)

		}(port)
	}
	wg.Wait()
}
```

Podemos ver que no me esta mostrando ningún puerto activo. Creo que se debe a que estamos ejecutándolo desde una ubuntu wsl y no me esta cogiendo los port del servicio. 

### Nota: Posibles respuestas dadas de chapGPT del porque puede ocurrir lo que esta pasando.

### Pregunta: estoy ejecutando un programa en go para identificar los puerto que se encuentra corriendo de un servicio, tenga en cuenta que tengo montado mi entorno de desarrollo en go en ubuntu wsl. Pero al ejecutar el programa no no me muestra los puertos que estan activos, este es el código: package main  - puese el codigo-   Pero en cambio si implemento expresiones regulares con la función checkEnviron par saber si por el ttl estamos en una máquina Linux o una caja windows como en el siguiente programa -puse el codigo del compañero-

Vamos a abordar este problema paso a paso:

1. **Ambos códigos**: Tienes dos programas en Go que, en esencia, hacen un escaneo de puertos. El segundo también verifica el valor TTL para identificar el sistema operativo.
2. **WSL**: El Windows Subsystem for Linux (WSL) es una capa de compatibilidad que permite ejecutar ejecutables de Linux en Windows. Sin embargo, hay algunas limitaciones y diferencias entre WSL y una verdadera máquina Linux. Por ejemplo, hay diferencias en la gestión de redes entre WSL1 y WSL2, y esto podría afectar la capacidad de tu programa para escanear puertos.
3. **Razón del problema**: Es posible que la primera versión del código no funcione debido a la forma en que está configurado el entorno WSL o debido a problemas con el firewall o el control de acceso en Windows o en el host de destino. Si el segundo programa funciona y el primero no, es posible que el problema esté en la configuración del entorno o en las herramientas que estás usando.
4. **Posibles soluciones**:
    
    a. **Versión de WSL**: Si estás usando WSL1, considera la posibilidad de cambiar a WSL2, que tiene una mejor implementación de la pila de red y podría solucionar el problema. Puedes comprobar qué versión de WSL estás usando con el comando: **`wsl --list --verbose`**. Si ves que estás en la versión 1, puedes cambiar a la versión 2 siguiendo las **[instrucciones oficiales de Microsoft](https://docs.microsoft.com/en-us/windows/wsl/install)**.
    
    b. **Firewall y Control de Acceso**: Asegúrate de que no haya un firewall o una configuración de control de acceso en Windows, WSL o el host de destino que esté bloqueando el escaneo de puertos. Puede que necesites configurar excepciones para permitir que tu programa escanee puertos.
    
    c. **Herramientas y Configuración**: Usa herramientas como **`nmap`** en tu entorno WSL para comprobar si puedes escanear puertos desde allí. Si **`nmap`** tampoco puede detectar puertos abiertos, es probable que el problema esté en tu entorno o configuración y no en tu código.
    
    d. **Código**: Revisa tu código para asegurarte de que estás usando las bibliotecas y funciones correctas y de que no hay errores lógicos. Por ejemplo, en tu segundo programa, estás usando el valor TTL para identificar el sistema operativo, pero ten en cuenta que el TTL puede ser alterado en tránsito o configurado de manera diferente en diferentes sistemas, por lo que no siempre es un indicador fiable.
    
5. **Consejo adicional**: Al realizar escaneos de puertos, especialmente en hosts que no te pertenecen, asegúrate de tener el permiso adecuado. Escanear puertos sin permiso puede ser considerado una actividad maliciosa y podría tener consecuencias legales.

Espero que estas recomendaciones te ayuden a identificar y resolver el problema. ¡Buena suerte!

Pero siguiendo con las recomendaciones de un compañero, nos indica lo siguiente:

Le agregué expresiones regulares en la función *checkEnviron* par saber si por el *ttl* estamos en una máquina Linux o una caja windows xd

```go
package main

import (
	"flag"
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	host = flag.String("host", "scanme.nmap.org", "host to be scanned")
)

func checkEnviron() (string, error) {

	out, err := exec.Command("ping", "-c", "1", *host).Output()
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`ttl=(.?).[\S]`)
	ttl := fmt.Sprintf("%s", re.FindString(string(out)))
	ttl = strings.Split(ttl, "=")[1]
	ttlNum, err := strconv.Atoi(ttl)
	if err != nil {
		return "", err
	}
	if ttlNum <= 64 {
		return "\n\t[+] Linux system\n", nil
	} else if ttlNum >= 127 {
		return "\n\t[+] Windows system\n", nil
	} else {
		return "\n\t[-] the time to the life of the target system doesn't exists\n", nil
	}

}

func main() {

	flag.Parse()

	environ, _ := checkEnviron()
	fmt.Println(environ)

	var wg sync.WaitGroup
	//wg.Add(1000)
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, j))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("[+] port %d open!!\n", j)
		}(i)
	}
	wg.Wait()
}
```

 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/cdb0810d-4bd0-4c33-b96e-9b72fb879563/Untitled.png)

como vemos, ya me lee los puerto de este nuevo host que le indicamos, o si no le ingresamos nada va a evaluar el host por defecto que tenemos configurado. 

## Proyecto

## ****Netcat****

```go

```

## ****Servidor de chat en Go con net****

## Terminando el Chat

## ****Resumen del curso de Go avanzado****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e11402e6-df0d-4361-8ae8-dea3230fceb8/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d028d508-8e0b-499b-a2cb-b6e038af4cb4/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e8d6a2d9-9d54-4c27-8f78-0754f41ceef9/Untitled.png)
