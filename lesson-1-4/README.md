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

# Ошибки


- Ошибка - тип, реализующий интерфейс error
- Функции возвращают ошибки как обычные значения
- По конвенции, ошибка - последнее возвращаемое функцией значение
- Ошибки обрабатываются проверкой значения (и/или передаются выше через обычный return)

```
type error interface {
    Error() string
}
```


```
func Marshal(v interface{}) ([]byte, error) {
   e := &encodeState{}
   err := e.marshal(v, encOpts{escapeHTML: true})
   if err != nil {
      return nil, err
   }
   return e.Bytes(), nil
}
```

---


# Интерфейс error



```
package errors

func New(text string) error {
   return &errorString{text}
}

type errorString struct {
   s string
}

func (e *errorString) Error() string {
   return e.s
}
```

---

# Обработка ошибок: sentinel values
<br>

```
package io


// ErrShortWrite means that a write accepted fewer bytes than requested
// but failed to return an explicit error.
var ErrShortWrite = errors.New("short write")

// ErrShortBuffer means that a read required a longer buffer than was provided.
var ErrShortBuffer = errors.New("short buffer")
```

Ошибки в таком случае - часть публичного API, это наименее гибкая стратегия:

```
if err == io.EOF {
	...
}
```

---

# Обработка ошибки как значения

```
scanner := bufio.NewScanner(input)
for scanner.Scan() {
    token := scanner.Text()
    // process token
}
if err := scanner.Err(); err != nil {
    // process the error
}
```

vs

```
func (s *Scanner) Scan() (token []byte, error)

scanner := bufio.NewScanner(input)
for {
    token, err := scanner.Scan()
    if err != nil {
        return err // or maybe break
    }
    // process token
}
```

---

# Идиоматичная проверка ошибок

<br>
В целом ок:

```
func (router HttpRouter) parse(reader *bufio.Reader) (Request, Response) {
  requestText, err := readCRLFLine(reader) //string, err Response
  if err != nil {
    //No input, or it doesn't end in CRLF
    return nil, err
  }

  requestLine, err := parseRequestLine(requestText) //RequestLine, err Response
  if err != nil {
    return nil, err
  }

  if request := router.routeRequest(requestLine); request != nil {
    return request, nil
  }

  //Valid request, but no route to handle it
  return nil, requestLine.NotImplemented()
}
```

---


# Проверка ошибок

Теперь мы можем использовать проверку типов ошибок:

```
for try := 0; try < 2; try++ {
    file, err = os.Create(filename)
    if err == nil {
        return
    }
    if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
        deleteTempFiles()  // Recover some space.
        continue
    }
    return
}
```

---

# Проверка ошибок: типы

```

    // PathError records an error and the operation and
    // file path that caused it.
    type PathError struct {
        Op string    // "open", "unlink", etc.
        Path string  // The associated file.
        Err error    // Returned by the system call.
    }
    
    func (e *PathError) Error() string {
        return e.Op + " " + e.Path + ": " + e.Err.Error()
    }

```

```
    open /etc/passwx: no such file or directory
```

---

# Проверка ошибок: типы

```
err := something()
switch err := err.(type) {
	case nil:
	    // call succeeded, nothing to do
	case *PathError:
        fmt.Println(“invalid config path:”, err.Path)
	default:
		// unknown error
}
```

--

# Проверка ошибок: интерфейсы

```
package net

type Error interface {
    error
    Timeout() bool   // Is the error a timeout?
    Temporary() bool // Is the error temporary?
}
```

Проверяем поведение, а не тип:

```
if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
    time.Sleep(1e9)
    continue
}
if err != nil {
    log.Fatal(err)
}
```

---


# Антипаттерны проверки ошибок:

```
if err.Error() == "smth" { // строковое представление - для людей
```

```
func Write(w io.Writer, buf []byte) {
        w.Write(buf) // забыли проверить ошибку
}
```

```
func Write(w io.Writer, buf []byte) error {
        _, err := w.Write(buf)
        if err != nil {
                // логируем ошибку вероятно несколько раз
				// на разных уровнях абстракции
                log.Println("unable to write:", err)
 
                // unannotated error returned to caller
                return err
        }
        return nil
}
```

---

# Антипаттерны проверки ошибок:

Никогда не проверяйте error.Error

```
func Write(w io.Writer, buf []byte) {
        w.Write(buf) // забыли проверить ошибку
}
```

---

# Stacktrace: github.com/pkg/errors
```
_, err := ioutil.ReadAll(r)
if err != nil {
        return errors.Wrap(err, "read failed")
}
```
```
package main

import "fmt"
import "github.com/pkg/errors"

func main() {
        err := errors.New("error")
        err = errors.Wrap(err, "open failed")
        err = errors.Wrap(err, "read config failed")

        fmt.Println(err) // read config failed: open failed: error
		fmt.Printf("%+v\n", err) // напечатает stacktrace

}
```

---

# Stacktrace: github.com/pkg/errors

Чтобы проверить, соответствует ли ошибка значению/типу, ее надо развернуть:

```
// Cause unwraps an annotated error.
func Cause(err error) error
```
```
	err1 := errors.New("im an error")
	err2 := errors.Wrap(err1, "some context")
	print(err1 == err2)				   // false
	print(err1 == errors.Cause(err2))  // true
```
---

# Stacktrace: github.com/pkg/errors
```
switch err := errors.Cause(err).(type) {
case *MyError:
        // handle specifically
default:
        // unknown error
}
```
```
// IsTemporary returns true if err is temporary.
func IsTemporary(err error) bool {
        te, ok := errors.Cause(err).(temporary)
        return ok && te.Temporary()
}
```

---

# Defer, Panic и Recover: defer

<br>
defer позволяет назначить выполнение вызова функции непосредственно
перед выходом из вызывающей функции

```
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
```

---

# Defer, Panic и Recover: defer

Аргументы отложенного вызова функции вычисляются тогда, когда вычисляется команда defer.

```
func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}
```

```
0
```

---

# Defer, Panic и Recover: defer


Отложенные вызовы функций выполняются в порядке LIFO: последний отложенный вызов будет вызван первым — после того, как объемлющая функция завершит выполнение.

```
func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}
```

```
3210
```

---
# Defer, Panic и Recover: defer

Отложенные функции могут читать и устанавливать именованные возвращаемые значения объемлющей функции.

```
func c() (i int) {
    defer func() { i++ }()
    return 1
}
```

Эта функция вернет 2

---

# Panic и Recover

Panic — это встроенная функция, которая останавливает обычный поток управления и начинает паниковать. Когда функция F вызывает panic, выполнение F останавливается, все отложенные вызовы в F выполняются нормально, затем F возвращает управление вызывающей функции. Для вызывающей функции вызов F ведёт себя как вызов panic. Процесс продолжается вверх по стеку, пока все функции в текущей го-процедуре не завершат выполнение, после чего аварийно останавливается программа. Паника может быть вызвана прямым вызовом panic, а также вследствие ошибок времени выполнения, таких как доступ вне границ массива.

<br>
<br>

Recover — это встроенная функция, которая восстанавливает контроль над паникующей го-процедурой. Recover полезна только внутри отложенного вызова функции. Во время нормального выполнения, recover возвращает nil и не имеет других эффектов. Если же текущая го-процедура паникует, то вызов recover возвращает значение, которое было передано panic и восстанавливает нормальное выполнение.

---

# Panic and recover

Паниковать стоит только в случае, если ошибку обработать нельзя, например:

```
var user = os.Getenv("USER")

func init() {
    if user == "" {
        panic("no value for $USER")
    }
}
```

---

# Panic and recover

<br><br>
"поймать" панику можно с помощью recover: вызов recover останавливает выполнение отложенных функций
и возвращает аргумент, переданнй panic
<br>

```
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
```

---

# Panic and recover


пример из encoding/json:

```
// jsonError is an error wrapper type for internal use only.
// Panics with errors are wrapped in jsonError so that the top-level recover
// can distinguish intentional panics from this package.
type jsonError struct{ error }

func (e *encodeState) marshal(v interface{}, opts encOpts) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if je, ok := r.(jsonError); ok {
				err = je.error
			} else {
				panic(r)
			}
		}
	}()
	e.reflectValue(reflect.ValueOf(v), opts)
	return nil
}
```

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
