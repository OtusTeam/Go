center.icon[![otus main](img/main.png)]

---


class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Интерфейсы <br> в Go

### Александр Давыдов

---

class: top white
background-image: url(img/sound.svg)
background-size: 130%
.top.icon[![otus main](img/logo.png)]

.sound-top[
  # Как меня слышно и видно?
]

.sound-bottom[
  ## > Напишите в чат
  ### **+** если все хорошо
  ### **-** если есть проблемы cо звуком или с видео
  ### !проверить запись!
]

---

# О чем будем говорить:

* Мьютексы
* Условные переменные
* Гарантировано одноразовое выполнение
* Pool и WaitGroup
* Модель памяти в Go
* Race-детектор

---

# WaitGroup

Что выведет эта программа?

```
type Dog struct { name string; walkDuration time.Duration }

func (d Dog) Walk() {
	fmt.Printf("%s is taking a walk\n", d.name)
	time.Sleep(d.walkDuration)
	fmt.Printf("%s is going home\n", d.name)
}

func walkTheDogs(dogs []Dog) {
	for _, d := range dogs { go d.Walk() }
	fmt.Println("everybody's home")
}


func main() {
	dogs := []Dog{{"vasya", time.Second}, {"john", time.Second*3}}
	walkTheDogs(dogs)
}
```


---

# Waitgroup

```
type WaitGroup struct {
        // неэкспортируемые поля
}

func (wg *WaitGroup) Add(delta int) - инерементирует счетчик WaitGroup на 1

func (wg *WaitGroup) Done() - декрементит счетчик на 1

func (wg *WaitGroup) Wait() - блокируется, пока счетчик WaitGroup не обнулится.
```

---


# Waitgroup

```
type Dog struct { name string; walkDuration time.Duration }

func (d Dog) Walk(wg *sync.WaitGroup) {
	fmt.Printf("%s is taking a walk\n", d.name)
	time.Sleep(d.walkDuration)
	fmt.Printf("%s is going home\n", d.name)
	wg.Done()
}


func main() {
	dogs := []Dog{{"vasya", time.Second}, {"john", time.Second*3}}
	var wg sync.WaitGroup

	for _, d := range dogs {
		wg.Add(1)
		go d.Walk(&wg)
	}

	wg.Wait()
	fmt.Println("everybody's home")
}
```

---

# Waitgroup

```
type httpPkg struct{}

func (httpPkg) Get(url string) {}

var http httpPkg

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			http.Get(url)
		}(url)
	}
	wg.Wait()
}
```


---

# Задачка

https://play.golang.org/p/m16jnq3kO2O

использовать WaitGroup чтобы выпустить собак одновременно
и дождаться их возвращения

---

# Waitgroup

fun fact: аргумент Add может быть отрицательным

```
// Done decrements the WaitGroup counter by one.
func (wg *WaitGroup) Done() {
	wg.Add(-1)
}
```


---

# Mutex

```
var i int // i == 0

func worker(wg *sync.WaitGroup) {
	i = i + 1
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()

	fmt.Println("value of i after 1000 operations is", i)
}
```

https://play.golang.org/p/MQNepChxiEa

---

# Mutex

```
➜  hello10 GOMAXPROCS=4 go run main.go
value of i after 1000 operations is 995
➜  hello10 GOMAXPROCS=4 go run main.go 
value of i after 1000 operations is 999
➜  hello10 GOMAXPROCS=4 go run main.go
value of i after 1000 operations is 992
➜  hello10 GOMAXPROCS=2 go run main.go
value of i after 1000 operations is 994
```

```
runtime.GOMAXPROCS(2)
```

---

# Mutex

```
hello10 go run -race grtn.go
==================
WARNING: DATA RACE
Read at 0x00000121e868 by goroutine 7:
  main.worker()
      /Users/alexander.davydov/h
```


---

# Mutex

Что могло пойти не так?
<br><br>
i = i + 1:

- достать значение i
- инкрементировать
- записать новое значение

G1 starts first when i is 0, run first 2 steps and i is now 1. But before G1 updates value of i in step 3, new goroutine G2 is scheduled and it runs all steps. But in case of G2, value of i is still 0 hence after it executes step 3, i will be 1. Now G1 is again scheduled to finish step 3 and updates value of i which is 1from step 2. In a perfect world where goroutines are scheduled after completing all 3 steps, successful operations of 2 goroutines would have produced the value of i to be 2 but that’s not the case here. Hence, we can pretty much speculate why our program did not yield value of i to be 1000.

So far we learned that goroutines are cooperatively scheduled. Until unless a goroutine blocks with one of the conditions mentioned in concurrency lesson, another goroutine won’t take its place. And since i = i + 1 is not blocking, why Go scheduler schedules another goroutine?

---

# Mutex

```
var i int // i == 0

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // acquire lock
	i = i + 1
	m.Unlock() // release lock
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &m)
	}

	wg.Wait()

	fmt.Println("value of i after 1000 operations is", i)
}
```

https://play.golang.org/p/xVFAX_0Uig8

---

# sync.Mutex

Вывод: не стоит принимать во внимание работу планировщика Go: синхронизировать работу горутин надо самому.

---

# sync.Mutex: паттерны использования

помещайте мьютекс выше тех полей, доступ к которым он будет защищать:

```
var sum struct {
    sync.Mutex     // <-- этот мьютекс защищает
    i int          // <-- поле под ним
}
```


держите блокировку не дольше, чем требуется:

```
func doSomething(){
    mu.Lock()
    item := cache["myKey"]
    http.Get() // какой-нибудь дорогой IO-вызов
    mu.Unlock()
}
```

---

# sync.Mutex: паттерны использования

используйте defer, чтобы разблокировать мьютекс там где у функции есть несколько точек выхода:

```
func doSomething() {
	mu.Lock()
	defer mu.Unlock()
    err := ...
	if err != nil {
		//log error
		return // <-- разблокировка произойдет здесь
	}

        err = ...
	if err != nil {
		//log error
		return // <-- или тут
	}
	return // <-- и тут тоже
}
```
---

# sync.Mutex: паттерны использования

но надо быть аккуратным:

```
func doSomething(){
    for {
        mu.Lock()
        defer mu.Unlock()
         
        // какой-нибудь интересный код
        // <-- defer будет выполнен не тут, а при выходе из функции
     }
}
// И поэтому в коде выше будет дедлок!
```

---

# sync.Pool

```
type Dog struct { name string }

func (d *Dog) Bark() { fmt.Printf("%s", d.name) }

var dogPack = sync.Pool{
	New: func() interface{} { return &Dog{} },
}

func main() {
	dog := dogPack.Get().(*Dog)
	dog.name = "ivan"
	dog.Bark()
	dogPack.Put(dog)
}
```

---

# sync.Pool

```
var gdog *Dog

func BenchmarkWithPool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dog := dogPool.Get().(*Dog)
			dog.name = "ivan"
			dogPool.Put(dog)
		}
	})
}

func BenchmarkWithoutPool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dog := &Dog{name:"ivan"}
			gdog = dog
		}
	})
}
```

---

# sync.Pool

```
go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/nyddle/dogpool
BenchmarkWithPool       100000000      17.5 ns/op     0 B/op     0 allocs/op
BenchmarkWithoutPool    50000000       26.0 ns/op    16 B/op     1 allocs/op
PASS
ok      github.com/nyddle/dogpool 3.109s
```

---

# sync.Pool

! Любой элемент, хранящийся в пуле, может быть удален автоматически в любое время без уведомления

runtime/mgc.go:

```
func gcStart(trigger gcTrigger) {
  [...]
  // clearpools before we start the GC
  clearpools()
```

Пример использования: https://golang.org/src/fmt/print.go#L109
Подробнее про устройство в 1.13: https://dev-gang.ru/article/go-ponjat-dizain-syncpool-cpvecztx8e/


---

# sync.Pool

 A Pool is a set of temporary objects that may be individually saved and retrieved.

Any item stored in the Pool may be removed automatically at any time without notification. If the Pool holds the only reference when this happens, the item might be deallocated.

A Pool is safe for use by multiple goroutines simultaneously.
Pool безпасен

Pool's purpose is to cache allocated but unused items for later reuse, relieving pressure on the garbage collector. That is, it makes it easy to build efficient, thread-safe free lists. However, it is not suitable for all free lists.

An appropriate use of a Pool is to manage a group of temporary items silently shared among and potentially reused by concurrent independent clients of a package. Pool provides a way to amortize allocation overhead across many clients.

An example of good use of a Pool is in the fmt package, which maintains a dynamically-sized store of temporary output buffers. The store scales under load (when many goroutines are actively printing) and shrinks when quiescent.

On the other hand, a free list maintained as part of a short-lived object is not a suitable use for a Pool, since the overhead does not amortize well in that scenario. It is more efficient to have such objects implement their own free list.

Pool нельзя копировать после первого использования.


---

# sync.Once

```
type Once struct {
	m    Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
```

---

# sync.Once

// Do вызывает функцию f только в том случае, если это первый вызов Do для 
// этого экземпляра Once. Другими словами, если у нас есть var once Once и 
// once.Do(f) будет вызываться несколько раз, f выполнится только в 
// момент первого вызова, даже если f будет иметь каждый раз другое значение.
// Для вызова нескольких функций таким способом нужно несколько
// экземпляров Once.
//
// Do предназначен для инициализации, которая должна выполняться единожды
// Так как f ничего не возвращает, может быть необходимым использовать
// замыкание для передачи параметров в функцию, выполняемую Do:
//  config.once.Do(func() { config.init(filename) })
//
// Поскольку ни один вызов к Do не завершится пока не произойдет 
// первый вызов f, то f может заблокировать последующие вызовы
// Do и получится дедлок.
//
// Если f паникует, то Do считает это обычным вызовом и, при последующих
// вызовах, Do не будет вызывать f.
//


---

# sync.Once

```
type OldDog struct {
	name string
	die sync.Once
}

func (d *OldDog) Die() {
	d.die.Do(func() { println("bye!") })
}


func main() {

	d := OldDog{name:"bob"}
	d.Die()
	d.Die()
	d.Die()
}
```

```
bye!
```

---

# sync.Once

package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
