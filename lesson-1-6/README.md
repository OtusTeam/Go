.center.icon[![otus main](img/main.png)]

---


class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Функции и ошибки <br> в Go

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

* Определение и реализация интерфейсов
* Внутренняя структура интерфейсов
* Интерфейсы как "универсальный" тип
* Определение типа значения интерфейса
* Опасный и безопасный type cast
* Конструкций switch 
* Слайсы и словари с интерфейсами
* Где мои generic-и?


---

# Определение и реализация

Интерфейс - это набор сигнатур методов, которые должен имплементировать объект, чтобы "реализовать интерфейс".

```
type Ducker interface {
	Talk() string
	Walk()
	Swim()
}

type Dog struct {
	name string
}

func (d Dog) Talk() string {
	return "AGGGRRRR"
}

func (d Dog) Walk() {
}

func (d Dog) Swim() {
}

```


---

# Тип может реализовать несколько интерфейсов:

```
type Hound interface {
	Hunt()
}
type Poodle interface {
	Bark()
}

type GoldenRetriever struct{name string}

func (GoldenRetriever) Hunt() { fmt.Println("hunt") }
func (GoldenRetriever) Bark() { fmt.Println("bark") }


func f1(i Hound) { i.Hunt() }
func f2(i Poodle) { i.Bark() }


func main() {
	t := GoldenRetriever{"jack"}
	f1(t) // "hunt"
	f2(t) // "bark"
}
```

---

# Тип может реализовать несколько интерфейсов:

```
type Hound interface {
	Hunt()
}
type Poodle interface {
	Bark()
}

type GoldenRetriever struct{name string}

func (GoldenRetriever) Hunt() { fmt.Println("hunt") }
func (GoldenRetriever) Bark() { fmt.Println("bark") }


func f1(i Hound) { i.Hunt() }
func f2(i Poodle) { i.Bark() }


func main() {
	t := GoldenRetriever{"jack"}
	f1(t) // "hunt"
	f2(t) // "bark"
}
```

---

# Одному интерфейсу могут соответствовать много типов

```
type Poodle interface {
	Bark()
}

type ScandinavianClip struct{name string}
func (ScandinavianClip) Bark() { fmt.Println("bark") }


type ToyPoodle struct{name string}
func (ToyPoodle) Bark() { fmt.Println("bark") }


func main() {
	var t, sc Poodle

	t = ToyPoodle{"jack"}
	sc = ScandinavianClip{"jones"}

	t.Bark() // "bark"
	sc.Bark() // "bark"
}
```

---

# Интерфейсы


In Go we've two concepts related to interfaces:

- Интерфейс — набор методов, которые надо реализовать, чтобы удовлетворить интерфейсу. Ключевое слово interface.

```
type Stringer interface {
    String() string
}
```

- Тип интерфейс — переменная типа интерфейс, которая содержит значение типа, который реализует интерфейс.

```
var s Stringer
```

---

# Интерфейсы

Интерфейс - это набор методов:

```
type Dog interface {
	Bark()
	Eat()
	Name() string
	Weight(pounds bool) float64
}
```

---

# Интерфейсы

Интерфейс - это набор методов:

```
type Dog interface {
	Bark()
	Eat()
	Name() string
	Weight(pounds bool) float64
}
```

----


# Интерфейсы

Если мы объявили переменную как интерфейс - мы можем использовать только методы этого интерфейса

```
type Loud interface {
	Bark()
}

type Dog struct {
	name string
}

func (d Dog) Bark() { println("agrr") }
func (d Dog) Eat() {}


func main() {
	var dog Loud = Dog{"joe"}
	dog.Bark()
	dog.Eat() //dog.Eat undefined (type Loud has no field or method Eat)
}
```

---

```
type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
    return len(s)
}
func (s Sequence) Less(i, j int) bool {
    return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
    copy := make(Sequence, 0, len(s))
    return append(copy, s...)
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
    s = s.Copy() // Make a copy; don't overwrite argument.
    sort.Sort(s)
    str := "["
    for i, elem := range s { // Loop is O(N²); will fix that in next example.
        if i > 0 {
            str += " "
        }
        str += fmt.Sprint(elem)
    }
    return str + "]"
}
```

---

# Интерфейсы 

позволяют динамически определять соответствие, что-то подобное происходит в пакете fmt:


```
type Stringer interface {
    String() string
}

func ToString(any interface{}) string {
    if v, ok := any.(Stringer); ok {
        return v.String()
    }
    switch v := any.(type) {
    case int:
        return strconv.Itoa(v)
    case float:
        return strconv.Ftoa(v, 'g', -1)
    }
    return "???"
}
```

---

# Интерфейсы: когда использовать?

- если можете использовать конкретный тип - используйте конкретный тип
- если можете не использовать interface{} - не используйте

---

# Интерфейс под капотом

```
type iface struct { // 16 bytes on a 64bit arch
    tab  *itab
    data unsafe.Pointer
}
```

- tab holds the address of an itab object, which embeds the datastructures that describe both the type of the interface as well as the type of the data it points to.
- data is a raw (i.e. unsafe) pointer to the value held by the interface.


---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
