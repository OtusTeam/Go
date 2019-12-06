.center.icon[![otus main](img/main.png)]

---


class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Плюсы и минусы языка Go

### Антон Телышев

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


  ### !включить запись!
]


---

# План занятия

.left-text[
<br>
* Почему появился Go?
* Поговорим о достоинствах и недостатках языка
* Рассмотрим программу на Go
]

.right-image[
  ![](img/gopher3.png)
]


---

# Преподаватель

###Антон Телышев

-- Ведущий разработчик в “Центр недвижимости от Сбербанка” (ДомКлик) <b>(Golang)</b><br><br>

-- Окончил магистратуру МГТУ им Н.Э. Баумана (кафедра “Компьютерные системы и сети”)<br><br>

-- Разрабатывал системы мониторинга и внутренние сервисы в Mail.Ru Group <b>(Python/Golang)</b><br><br>

-- Разрабатывал Kaspersky Fraud Prevention Cloud <b>(Python/C++)</b><br><br>

-- Окончил Технопарк Mail.Ru (системный архитектор), где впоследствии преподавал "Подготовительную программу по С++"<br><br>

-- Telegram: @antonboom


---

# Кто уже пишет на Go?

- по работе
- для себя
- не пишу


---

class: bottom
background-image: url(img/rating-1.png)
background-size: 40%

# Рейтинг StackOverflow (самый популярный)

https://proglib.io/p/stack-overflow-2018/


---

class: bottom
background-image: url(img/rating-2.png)
background-size: 40%

# Рейтинг StackOverflow (самый нужный)

https://proglib.io/p/stack-overflow-2018/


---

# Где используется Go

- Веб (backend)
- Системные утилиты
- Devops
- Сетевое программирование

Go — это язык для создания систем. Он замечательно подходит для облачных решений (веб-серверов, кэшей), микросервисов и распределенных систем


---

# Где не используется Go

- Разработка игр
- Научные вычисления
- Машинное обучение
- Встраевыемые устройства


---

# Что написано на Go

- Grafana
- Docker
- Consul
- Kubernetes
- Prometheus

Кто использует Go?<br>
https://github.com/golang/go/wiki/GoUsers


---

# Немного истории

1956-1958 LISP<br>
1959 Cobol<br>
1964 Basic<br>
1970 Pascal<br>
1970 C<br>
1978 SQL<br>
1983 C++<br>
1991 Python<br>
1995 Java<br>
1995 PHP<br>
<b>2009 Go<br></b>
2010 Rust


---

class: bottom
background-image: url(img/proc.png)
background-size: 80%

# Развитие процессоров


---

# Golang

<b>Go (часто также Golang)</b> — компилируемый многопоточный язык программирования, разработанный внутри компании
Google.
<br><br>
Разработка Go началась в сентябре 2007 года, его непосредственным проектированием занимались <b>Роберт
Гризмер</b>, <b>Роб Пайк</b> и <b>Кен Томпсон</b>, занимавшиеся до этого проектом разработки операционной системы Inferno.
<br><br>
Официально язык был представлен в ноябре 2009 года.
<br><br>
https://github.com/golang/go


---

# Проблемы Google, подтолкнувшие к Go


* Медленная сборка (вплоть до часа)

* Неконтролируемые зависимости

* Каждый программист использует свое подмножество языка

* Трудность в чтении чужого кода

* Сложности деплоя (инструменты автоматизации, межъязыковые сборки и пр.)

<br><br>
https://talks.golang.org/2012/splash.article


---

# Требования к Go

* Возможность работы на больших масштабах: крупные команды разработчиков, большое количество зависимостей


* Должен быть знакомым программистам Google, а значит - Си-подобным


* Должен быть современным:
  - использование возможностей многоядерных машин "из коробки"
  - встроенные библиотеки для работы с сетью и пр.


<br><br>
https://talks.golang.org/2012/splash.article


---

# Характеристики Go

-- Императивный

-- Компилируемый в нативный код

-- Статически типизируемый

-- Нет классов, но есть структуры с методами

-- Есть интерфейсы

-- Нет наследования, но есть встраивание

-- Функции - объекты первого класса

-- Есть замыкания

-- Функции могут возвращать больше 1 значения

-- Есть указатели, но нет адресной арифметики

-- Обширные возможности для конкурентности

-- Сборка в 1 бинарный файл

-- Набор стандартных инструментов


---

.center[
  ## Достоинства Go
  ![](img/gopher11.png) 
]


---

# Строгая статическая явная типизация

- Типобезопасность
- Обнаружение многих ошибок уже на этапе компиляции


---

# Скорость компиляции

На простых программах ощущение работы с интерпретируемым языком


---

# Высокая скорость исполнения

Код на Go компилируется напрямую в машинный код, который зависит от выбранной ОС и архитектуры процессора машины (GOOS, GOARCH)


---

# Переносимость

- Бинарные файлы переносимы в рамках одной ОС и архитектуры
- Возможность кросс-компиляции (GOOS, GOARCH, CGOENABLED)


---

# Concurrency

- Механизмы реализации конкурентности имеют строгое теоретическое основание (http://www.cs.cmu.edu/~crary/819-f09/Hoare78.pdf)


- Сотни тысяч горутин на одной машине:
  - маленький стек горутины
  - дешёвое переключение контекста
  - мультиплексирование горутин по ядрам ОС


---

# Интерфейсы

Позволяют создавать систему на основе слабосвязанных либо совершенно не связанных компонентов.

```golang
package io

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

- не нужно явно указывать, что структура реализует интерфейс
- код полагается на абстракцию, не на конкретную реализацию
- хорошая архитектура для юнит-тестирования


---

# Сборщик мусора

- не нужно заботиться о висячих указателях и освобождении динамически выделенной памяти

- escape analysis


---

# Набор библиотек

### Богатая стандартная библиотека:
-- net/http 

-- database/sql

-- encoding/json

-- html/templates 

-- io/ioutil<br>
и пр. <br><br>

### Сообщество не дремлет:
-- https://github.com/avelino/awesome-go

-- https://go.dev/
<br><br>

### Playground:
-- https://play.golang.org

-- https://goplay.space


---

# Встроенные инструменты

<br>
```
$ go build
$ godoc
$ go generate
$ go fmt
$ go run
$ go tool cover
$ go tool pprof
$ go tool trace
$ go vet
$ go test [-bench] [-coverprofile] [-race]
```

Сторонние инструменты:

-- https://github.com/golangci/golangci-lint

-- https://github.com/alecthomas/gometalinter

-- https://github.com/golang/lint

-- https://github.com/TrueFurby/go-callvis

-- https://github.com/kyoh86/richgo

.right-image[
  ![](img/jake.gif)
]


---

# Встроенные инструменты

<style type="text/css">
  .tools {
    display: flex;
    flex-wrap: row;
    align-content: flex-start;
  }
  .cover-trace {
    display: flex;
    flex-wrap: wrap;
  }
  .cover-trace img {
    max-width: 450px;
    max-height: 300px
  }
</style>


<div class="tools">
  <div class="cover-trace">
    <img src="img/cover.png">
    <img src="img/trace.png">
  </div> 
  <div class="pprof">
    <img height="500px" src="img/pprof.png">
  </div>
</div>


---

.center[
  ## Недостатки Go
  ![](img/gopher12.png) 
]


---

# Отсутствие дженериков

Обобщённое программирование (англ. generic programming) — парадигма программирования, заключающаяся в таком описании данных и алгоритмов, которое можно применять к различным типам данных, не меняя само это описание. 

<br><br><br>
### Зачем нужны дженерики в Go?
https://habr.com/ru/company/mailru/blog/462811/

<br><br><br>
### Варианты решения:
- interface{}
- кодогенерация


---

# Шаблонная обработка ошибок

.center-image[
  ![](img/err.png)
]


---

# Недостатки Go, а какие чувствуете вы?

- Часть вещей приходится писать руками
- Плюсовики с трудом переходят на Go (в отличие от питонистов)
- Встроенные range, len, cap, make и пр. только для builtin типов
- Не попишешь расширения для С++
- Определение области видимости по регистру
- "Молодой" компилятор
- Игнорирует достижения современного проектирования языков
- Нет перечислений (enums)
- Отсутствие препроцессинга


---

# Разберем код программы на Go


### Книга "Go in Action"
https://github.com/goinaction/code/tree/master/chapter2/sample


---

# Заповеди Роб Пайка

-- Don't communicate by sharing memory, share memory by communicating.

-- Concurrency is not parallelism.

-- The bigger the interface, the weaker the abstraction.

-- Make the zero value useful.

-- interface{} says nothing.

-- Gofmt's style is no one's favorite, yet gofmt is everyone's 
favorite.

-- A little copying is better than a little dependency.

-- Syscall must always be guarded with build tags.

-- Clear is better than clever.

-- Reflection is never clear.

-- Errors are values.

-- Don't just check errors, handle them gracefully.

-- Documentation is for users.

-- Don't panic.

<br>
https://go-proverbs.github.io/


---

# Опрос

.left-text[
  Заполните пожалуйста опрос
  <br><br>
  [https://otus.ru/polls/6302/](https://otus.ru/polls/6302/)
]

.right-image[
  ![](img/gopher7.png)
]

---

# Дополнительные материалы

<br>
<b>Красота Golang</b><br>
https://evilinside.ru/krasota-go-lang/

<br>
<b>Go: хороший, плохой, злой</b><br>
https://habr.com/ru/company/mailru/blog/353790/

<br>
<b>Плюсы и минусы Go для разработчиков на C++ (нужен VPN)</b><br>
https://www.slideshare.net/yandex/go-c-39549651

<br>
О плюсах и минусах Go<br>
https://habr.com/ru/post/229169/

<br>
Go: недостатки<br>
https://bolknote.ru/all/3258/

<br>
Почему язык Go?<br>
https://geekbrains.ru/posts/why_go

<br>
Чем хорош Go и зачем его изучать?<br>
https://proglib.io/p/language-go

<br>
Плюсы и минусы Go<br>
https://www.andmed.org/


---

class: bottom
background-image: url(img/rating-3.png)
background-size: 48%

# P.S. Рейтинг зарплат (Мой круг)

https://habr.com/ru/company/moikrug/blog/420391/


---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
