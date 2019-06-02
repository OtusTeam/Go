.center.icon[![otus main](img/main.png)]

---


class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Слайсы и словари <br> в Go

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
]

---

# Небольшой тест

.left-text[
Пожалуйста, пройдите небольшой тест. 
<br><br>
Возможно вы уже многое знаете про слайсы и словари в Go =)
<br><br>
[]()
]

.right-image[
![](img/gopher9.png)
]

---


# Массивы

Массив - нумерованая последовательность элементов фиксированной длинны.
Массив располагается последовательно в памяти и не меняет своей длинны.

```
var arr [256]int        // фиксированная длинна

var arr [10][10]string  // может быть многомерным 

arr := [10]int{1,2,3,4,5}
```

Длинна массива - часть типа, т.е. массивы разной длинны это разные типы данных.

---


# Операции над массивами

Все ожидаемо

```
arr[3] = 1  // индексация

len(arr)    // длинна массива

arr[3:5]    // получение слайса
```

---


# Слайсы

Слайсы - это те же "массивы", но переменной длинны.
<br><br>
Создание слайсов:

```
var s []int   // s будет инициализирован Zero Value = nil

s := []int{}  // c помощью литерала слайса

s := make([]int, 5) // с помощью функции make

s := make([]int, 5, 10)

```
---


# Добавление элементов в слайс

Добавить новые элементы в слайс можно с помощью функции `append`

```
s[i] = 1               // работает если i < len(s)

s[len(s) + 10] = 1     // случится panic

s = append(s, 1)       // добавляет 1 в конец слайса

s = append(s, 1, 2, 3) // добавляет 1, 2, 3 в конец слайса

s = append(s, s2...)   // добавляет содержимое слайса s2 в конец s
```

---


# Получение под-слайсов (нарезка)

`s[i:j]` - возвращает под-слайс, с `i`-ого элемента включительно, по `j`-ый не влючительно.
Длинна нового слайса будет `j-i`.

```
s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

s2 := s[:]    // копия s (shallow) 

s2 := s[3:5]  // []int{3,4}

s2 := s[3:]   // []int{3, 4, 5, 6, 7, 8, 9}

s2 := s[:5]   // []int{0, 1, 2, 3, 4}
```

---


# Как это реализовано ?

```

// runtime/slice.go

type slice struct {
  array unsafe.Pointer
  len   int
  cap   int
}
```

```
l := len(s)  // len - вернуть длинну слайса

c := cap(s)  // cap - вернуть емкость слайса
```

Отличное описание: [https://blog.golang.org/go-slices-usage-and-internals](https://blog.golang.org/go-slices-usage-and-internals)

---


# Получение под-слайса (нарезка)

```
s := []byte{1,2,3,4,5}

s2 := s[2:5]
```

.left-image[
![](img/slice.png)
]

.right-image[
![](img/slice2.png)
]

---


# Авто-увеличение слайса

Если `len < cap` - увеличивается `len` <br><br>
Если `len = cap` - увеличивается `cap`, выделяется новый кусок памяти, данные копируются.

```
arr := []int{1}

for i := 0; i < 100; i++ {

  fmt.Printf("len: %d \tcap %d  \tptr %0x\n",
             len(arr), cap(arr), &arr[0])

  arr = append(arr, i)

}
```

Попробуйте на [https://play.golang.org/p/g7cjWi_dF9F](https://play.golang.org/p/g7cjWi_dF9F)

---


# Неочевидные следствия

При копировании слайса (а так же получени под-слайса и передаче в функцию) копируется только заголовок. Область памяти остается общей. Но только до тех пор пока один из слайсов не "вырастет" (произведет реаллокацию)

```
arr := []int{1, 2}

arr2 := arr
arr2[0] = 42
fmt.Println(arr[0]) // ?

arr2 = append(arr2, 3, 4, 5, 6, 7, 8, 9, 0)
arr2[0] = 1
fmt.Println(arr[0]) // ?
```

---


# Неочевидные следствия

![img/share.png](img/share.png)

<br><br>
Попробуйте на [https://play.golang.org/p/d-QBZnH5Jd6](https://play.golang.org/p/d-QBZnH5Jd6)

---


# Правила работы со слайсами

- Если хотите написать функцию *изменяющую* слайс, сделайте так что бы он возвращала новый слайс.
  Не изменяйте слайсы, которые передали вам как аргументы, т.к. это shalow копии исходых слайсов.
```
func AppendUniq(slice []int, slice2 []int) []int {
  ...
}
s = AppendUniq(s, s2)
```

- Если хотите получить полную копию, используйте функцию `copy`
```
s := []int{1,2,3}
s2 := make([]int, len(s))
copy(s2, s2)
```

---

# Сортировка

Для сортировки используется пакет `sort`

```
import sort

s := []int{3, 2, 1}
sort.Ints(s)

s := []string{"hello", "cruel", "world"}
sort.Strings(s)

// а что если нужно сортировать свои типы ?
s := []User{ 
  {"vasya", 19},
  {"petya", 18},
}
sort.Slice(s, func(i, j) bool {
  return s[i].Age < s[j].Age
})
```

---


# Словари (map)

---



# Небольшой тест

.left-text[
Проверим что мы узнали за этот урок
<br><br>
[]()
]

.right-image[
![](img/gopher9.png)
]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!