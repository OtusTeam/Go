.center.icon[![otus main](img/main.png)]

---


class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Структуры в Go

### Дмитрий Смаль

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

# Небольшой тест

.left-text[
Пожалуйста, пройдите небольшой тест. 
<br><br>
Возможно вы уже многое знаете про структуры в Go =)
<br><br>
[https://forms.gle/xLLab1NXH9NLKJij8](https://forms.gle/xLLab1NXH9NLKJij8)
]

.right-image[
![](img/gopher9.png)
]

---


# Структуры

Структуры - фиксированный набор именованных переменных.
Переменные размещаются рядом в памяти и обычно используются совместно.

```
struct{}  // пустая структура, не занимает памяти

type User struct { // структура с именованными полями
  Id      int64
  Name    string
  Age     int
  friends []int64  // приватный элемент
}
```

---

# Литералы структур

```
u1 :=  User{}  // Zero Value для типа User

u2 := &User{}  // Тоже, но указатель

u3 := User{1, "Vasya", 23}  // По номерам полей

u4 := User{
  Id:       1,
  Name:     "Vasya",
  friends:  []int64{1, 2, 3},
}

```

---


# Анонимные типы и структуры

Анонимные типы задаются литералом, у такого типа нет имени. 

Типичный сценарий использования: когда структура нужно только внутри одной функции. 

```
var wordCounts []struct{w string; n int}
```

```
var resp struct {
  Ok        bool `json:"ok"`
  Total     int  `json:"total"`
  Documents []struct{
    Id    int    `json:"id"`
    Title string `json:"title"`
  } `json:"documents"`
}
json.Unmarshal(data, &resp)
fmt.Println(resp.Documents[0].Title)
```
[https://play.golang.org/p/iDFhk57b3vf](https://play.golang.org/p/iDFhk57b3vf)

---

# Размер и выравнивание структур

Узнать размер любого типа (без внутренних структур) можно с помощью `unsafe.Sizeof`

```
unsafe.SizeOf(1)    // 8 на моей машине
unsafe.SizeOf("A")  // 16 (длина + указатель)
var x struct {
  a byte     // 1
  b bool     // 1
  c uint64   // 8
}
unsafe.SizeOf(x)    // 16 !
```
![img/aling.png](img/align.png)

Узнать смещение поля в структуре можно с помощью `unsafe.Offsetof`
---

# Указатели

Указатель - это адрес некоторого значения в памяти. Указатели строго типизированы.
Zero Value для указателя - nil.

```
x := 1         // Тип int
xPtr := &x     // Тип *int
var p *int     // Тип *int,  значение nil
```

---

# Получение адреса

Можно получать адрес не только переменной, но и поля структуры или элемента массива или слайса.
Получение адреса осуществляется с помощью оператора `&`.
```
var x struct {
  a int
  b string
  c [10]rune
}
bPtr := &x.b
c3Ptr := &x.c[2]
```

Но не значения в словаре!
```
dict := map[string]string{"a": "b"}
valPtr := &dict["a"]  // не скомпилируется
```

Так же нельзя (и не нужно) получить указатель на функцию.


---

# Разыменование указателей

Разыменование осуществляется с помощью оператора `*`
```
a := "qwe"  // Тип string
aPtr := &a  // Тип *string
b := *aPtr  // Тип string, значение "qwe"

var n *int  // nil
nv := *n    // panic
```

В случае указателей на *структуры* в можете обращаться к полям структуры без разыменования
```
p := struct{x, y int }{1, 3} // структура
pPtr= &p                     // указатель
fmt.Println(pPtr.x)
fmt.Println(pPtr.y)
```

---


# Копирование указателей и структур

При присвоении переменных типа структура - данные копируются.
```
a := struct{x, y int}{0, 0}
b := a
a.x = 1
fmt.Println(b.x) // 0
```

При присвоении указателей - копируется только адрес данных.
```
a := new(struct{x, y int})
b := a
a.x = 1
fmt.Println(b.x) // 1
```

---

# Определение методов 

В Go можно определять методы у именованых типов (кроме интерфейсов)

```
type User struct {
  Id      int64
  Name    string
  Age     int
  friends []int64
}

func (u User) IsOk() bool {
  for _, fid := range u.friends {
    if u.Id == fid {
      return true
    }
  }
  return false
}

u := User{}
fmt.Println(u.IsOk())

```

---

# Методы типа и указателя на тип

Методы объявленные над типом получают копию объекта, поэтому не могут его изменять!
```
func (u User) HappyBirthday() {
  u.Age++   // это изменение будет потеряно
}
```

Методы объявленные над указателем на тип - могут
```
func (u *User) HappyBirthday() {
  u.Age++   // OK
}
```

Метод типа можно вызывать у значения и у указателя. <br>
Метод указателя можно вызывать у указателя и у значения, если оно адресуемо.
---

# Экспортируемые и приватные элементы

Элементы структур, начинающиеся со строчной буквы - приватные, они будут видны
только в том же пакете, где и структура. Элементы, начинающиеся с заглавной - публичны, 
они будут видны везде.

```
type User struct {
  Id      int64
  Name    string   // экспортируемый элемент
  Age     int
  friends []int64  // приватный элемент
}
```

Не совсем очевидное следствие: пакеты стандартной библиотеки, например `encoding/json` тоже не могут :)<br><br>
Доступ к приватным элементам (на чтение!) все же можно получить с помощью пакета `reflect`

---

# Функции-конструкторы

В Go принят подход Zero Value: постарайтесь сделать так, что бы
ваш тип работал без инициализации, т.е. что было возможным сделать:
```
var someVar YourType
someVar.doSomeJob()
```

Если ваш тип содержит словари, каналы или инициализация обязательна - скройте
ее от пользователя, создав функции-конструкторы:

```
func  NewYourType() (*YourType) {
  // ...  
}
func NewYourTypeWithOption(option int) (*YourType) {
  // ...
}
```

Если опций и настроек много - можно использовать функции-кастомизаторы:
[https://github.com/samuel/go-zookeeper/blob/master/zk/conn.go#L175](https://github.com/samuel/go-zookeeper/blob/master/zk/conn.go#L175)

---

# Задачка

.left-code[
Реализовать тип `IntStack`, который содержит стэк целых чисел. 
У него должны быть методы `Push(i int)` и `Pop() int`.

<br><br>

[https://play.golang.org/p/yoaEnO1Bct1](https://play.golang.org/p/yoaEnO1Bct1)
]

.right-image[
![](img/gopher9.png)
]

---

# Встроенные структуры

В Go есть возможность "встраивать" типы внутрь структур. При этом у элемента структуры НЕ задается имя.

```
type LinkStorage struct {
  sync.Mutex                  // только тип!
  storage map[string]string   // тип и имя
}
```

Как обращаться к элементам встроенных типов
```
storage := LinkStorage{}
storage.Mutex.Lock()     // имя типа используется 
storage.Mutex.Unlock()   // как имя элемента структуры
```

---

# Продвижение методов

При встраивании методы встроенных стркутур можно вызывать у ваших типов!

```
// вместо
storage.Mutex.Lock()
// можно просто
storage.Lock()
```

Как следствие: если тип `A` реализует некоторый интрефейс `I` и тип `B` встраивает `A`,
то он автоматически реализует интерфейс `I`. <br>

Например `LinkStorage` теперь реализует интрефейс `sync.Locker`.

---

# Но, это не наследование

При вызове "продвинутых" методов, встроенный тип не имеет ни какой информации настоящем объекте.
```
type Base struct {}
func (b Base) Name() string {
  return "Base"
}
func (b Base) Say() {
  fmt.Println(b.Name())
}

type Child struct {
  Base
  Name string
}
func (c Chind) Name() string {
  return "Child"
}

c := Child{}
c.Say() // Base   увы =(
```
---

# Тэги элементов структуры

К элементам структуры можно добавлять метаинформацию - тэги. <br>
Тэг это просто литерал строки, но есть соглашение о структуре такой строки:
```
`key:"value"  key1:"value1,value11"`
```
Например
```
type User struct {
  Id      int64    `json:"-"`    // игнорировать в encode/json
  Name    string   `json:"name"`
  Age     int      `json:"user_age" db:"how_old:`
  friends []int64 
}
```

Получить информацию о тэгах можно через `reflect`
```
u := User{}
ut := reflect.TypeOf(u)
ageField := ut.FieldByName("Age")
jsonSettings := ageField.Get("json")  // "user_age"
```

---

# Использование тэгов для JSON сериализации

Для работы с JSON используется пакет `encoding/json`

```
// Можно задать имя поля в JSON документе
Field int `json:"myName"`

// Не выводить в JSON поля у которых Zero Value
Author *User `json:"author,omitempty"`

// Использовать имя поля Author, но не выводить Zero Value
Author *User `json:",omitempty"`

// Игнорировать это поле при сериализации / десереализации
Field int `json:"-"`
```

---

# Использование тэгов для работы с СУБД

Зависит от пакета для работы с СУБД.<br>
Например, для `github.com/jmoiron/sqlx`
```
var user User
row := db.QueryRow("SELECT * FROM users WHERE id=?", 10)
err = row.Scan(&user)
```

Для ORM библиотеки GORM `github.com/jinzhu/gorm` фич намного больше
```
type User struct {
  gorm.Model
  Name         string
  Email        string  `gorm:"type:varchar(100);unique_index"`
  Role         string  `gorm:"size:255"` // set field size to 255
  MemberNumber *string `gorm:"unique;not null"` // set member number to unique and not null
  Num          int     `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
  Address      string  `gorm:"index:addr"` // create index with name `addr` for address
  IgnoreMe     int     `gorm:"-"` // ignore this field
}
```
---






# Небольшой тест

.left-text[
Проверим что мы узнали за этот урок
<br><br>
[https://forms.gle/xLLab1NXH9NLKJij8](https://forms.gle/xLLab1NXH9NLKJij8)
]

.right-image[
![](img/gopher9.png)
]

---

# Домашнее задание

Реализовать двусвязанный список. <br>
Что такое двусвязный список: [https://en.wikipedia.org/wiki/Doubly_linked_list](https://en.wikipedia.org/wiki/Doubly_linked_list) <br><br>

Ожидаемые типы (псевдокод):

```
List      // тип контейнер
  Len()   // длинна списка
  First() // первый Item
  Last()  // последний Item
  PushFront(v interface{}) // добавить значение в начало
  PushBack(v interface{})  // добавить значение в конец

Item   // элемент списка
  Value() interface{}  // возвращает значение
  Nex() *Item          // следующий Item
  Prev() *Item         // предыдущий
  Remove()             // удалить Item из списка
```
<br><br>

(*) Желательно написать тесты

---

# Опрос

.left-text[
Заполните пожалуйста опрос
<br><br>
[https://otus.ru/polls/3645/](https://otus.ru/polls/3645/)
]

.right-image[
![](img/gopher7.png)
]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!


---