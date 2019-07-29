.center.icon[![otus main](img/main.png)]

---


class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Кодогенерация Go

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


# План занятия

.big-list[
* Посмотрим, где нам может помочь генерация кода
* Детально рассмотрим Protocol Buffers
* Поговорим о тестировании
]

---


# Go generate

```
package main

import (
	"fmt"
)

//go:generate echo "Hello, world!"

func main() {
	fmt.Println("run any unix command in go:generate")
}
```

```
> go generate
Hello, world!
```

---

# Зачем?

 - генерировать структуры на основе JSON
 - генерировать заглушки для интерфейсов (mocks для тестов)
 - protobufs: генерировать кода из описания протокола (.proto)
 - yacc: генерация .go файлов из  yacc (.y)
 - Unicode: generating tables from UnicodeData.txt
 - HTML: embedding .html files into Go source code
 - bindata: вставка бинарных данных JPEGs в код на Go в виде byte array

 - string methods: generating String() string methods for types used as enumerated constants
 - macros: generating customized implementations given generalized packages, such as sort.Ints from ints


---

# Цикл разработки пакета с go generate

```
	% edit …
	% go generate
	% go test

	% git add *.go  # коммитим сгенерированный код
	% git commit
```

---


# иногда достаточно ldflags

```
package main

import (
	"fmt"
)

var VersionString = "unset"

func main() {
	fmt.Println("Version:", VersionString)
}
```

```
go run -ldflags '-X main.VersionString=1.0' main.go
```


---


# Принципы go generate


- go generate запускаеися разработчиком программы/пакета, а не пользователем
- инструментария для go generate находится у создателя пакета
- генерация кода не должна происходить автоматически во время go build, go get, но вызываться эксплицитно
- инструменты генерации кода "невидимы" для пользователя, и могут быть недоступны для него
- go generate работает только с .go-файлами, как часть тулкита go 


- не забывайте добавлять disclaimer

```
/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/gogen/unmarshalmap
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/
```

https://docs.google.com/document/d/1V03LUfjSADDooDMhe-_K59EgpTEm3V8uvQRuNMAEnjg/edit

---

# Yacc

As outlined above, we define a custom command

	//go:generate -command yacc go tool yacc

and then anywhere in main.go (say) we write

	//go:generate yacc -o foo.go foo.y

---

# Binary data

A tool that converts binary files into byte arrays that can be compiled into Go binaries would work similarly. Again, in the Go source we write something like

//go:generate bindata -o jpegs.go pic1.jpg pic2.jpg pic3.jpg

This is also demonstrates another reason the annotations are in Go source: there is no easy way to inject them into binary files.


---

Sort

One could imagine a variant sort implementation that allows one to specify concrete types that have custom sorters, just by automatic rewriting of macro-like sort definition. To do this, we write a sort.go file that contains a complete implementation of sort on an explicit but undefined type spelled, say, TYPE. In that file we provide a build tag so it is never compiled (TYPE is not defined, so it won't compile) but is processed by go generate:

	// +build generate

Then we write an generator directive for each type for which we want a custom sort:

//go:generate rename TYPE=int
//go:generate rename TYPE=strings

or perhaps

	//go:generate rename TYPE=int TYPE=strings

The rename processor would be a simple wrapping of gofmt -r, perhaps written as a shell script.

There are many more possibilities, and it is a goal of this proposal to encourage experimentation with pre-build-time code generation.


---

# Вернемся к дженерикам

```
The generic dilemma is this: do you want slow programmers, 
slow compilers and bloated binaries, or slow execution times?
(c) Russ Cox
```

---

# Какие варианты:

- copy & paste (см. пакеты strings and bytes)
- интерфейсы
```
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```
- type assertions
- рефлексия
- go generate

---

# Generics!

```
go get github.com/cheekybits/genny
```

объявляем заглушки по типам:

```
type KeyType generic.Type
type ValueType generic.Type
```

пишем обычный код:

```
func SetValueTypeForKeyType(key KeyType, value ValueType) { /* ... */ }
```

---

# Генерация go структур из JSON

https://mholt.github.io/json-to-go/

```
go get github.com/ChimeraCoder/gojson/gojson
```

```
cat schema.json| gojson -name Person

package main

type Person struct {
        Age     int64    `json:"age"`
        Courses []string `json:"courses"`
        Name    string   `json:"name"`
}
```


---

# impl

```
go get -u github.com/josharian/impl
```

```
$ impl 'f *File' io.ReadWriteCloser
func (f *File) Read(p []byte) (n int, err error) {
	panic("not implemented")
}

func (f *File) Write(p []byte) (n int, err error) {
	panic("not implemented")
}

func (f *File) Close() error {
	panic("not implemented")
}
```

```
impl 's *Shortener' github.com/nyddle/shortener/service.Shortener
func (s *Shortener) Shorten(url string) string {
	panic("not implemented")
}

func (s *Shortener) Resolve(url string) string {
	panic("not implemented")
}

```

---

---

# TODO: implement shortener interface


---

# JSON Enums

```
go get github.com/campoy/jsonenums
```

```
//go:generate jsonenums -type=Status
type Status int

const (
	Pending Status = iota
	Sent
	Received
	Rejected
)
```

---

# Stringer

```
go get golang.org/x/tools/cmd/stringer
```

```
//go:generate stringer -type=MessageStatus
type MessageStatus int

const (
	Sent MessageStatus = iota
	Received
	Rejected
)
```

```
func main() {
	status := Sent
	fmt.Printf("Message is %s", status) // Message is Sent
}
```

---

# Protocol buffers


xml:
```
<person>
  <name>Elliot</name>
  <age>24</age>
</person>
```

json:
```
{
  "name": "Elliot",
  "age": 24
}
```

protobuf:
```
[10 6 69 108 108 105 111 116 16 24]
```

---

# Protocol buffers: запись и чтение

```
	course := &myotus.Course{
		Title:   "Golang",
		Teacher: []*myotus.Teacher{{Name: "Dmitry Smal", Id: 1}, {Name: "Alexander Davydov", Id: 2}},
	}

	out, err := proto.Marshal(course)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
```

```
	otusdb := &myotus.Otus{}
	if err := proto.Unmarshal(in, otusdb); err != nil {
		log.Fatalln("Failed to parse otus database:", err)
	}
```

---

# Protocol buffers: чтение


---

# Protocol buffers: типы данных




---

# Protocol buffers: go code

```
message Foo {}
```

```
type Foo struct {
}

// Reset sets the proto's state to default values.
func (m *Foo) Reset()         { *m = Foo{} }

// String returns a string representation of the proto.
func (m *Foo) String() string { return proto.CompactTextString(m) }

// ProtoMessage acts as a tag to make sure no one accidentally implements the
// proto.Message interface.
func (*Foo) ProtoMessage()    {}
```



---

# Protocol buffers

The process is the same as with yacc. Inside main.go, we write, for each protocol buffer file we have, a line like

	//go:generate protoc -go_out=. file.proto

Because of the way protoc works, we could generate multiple proto definitions into a single .pb.go file like this:

	//go:generate protoc -go_out=. file1.proto file2.proto

Since no globbing is provided, one cannot say *.proto, but this is intentional, for simplicity and clarity of dependency.

Caveat: The protoc program must be run at the root of the source tree; we would need to provide a -cd option to it or wrap it somehow.

---

# Protocol buffers: Packages and input paths

Каждый .proto файлик соответствует пакету

```
option go_package = "github.com/golang/protobuf/ptypes/any";
```

---

# Protocol buffers: generated code

A summary of the properties of the protocol buffer interface for a protocol buffer variable v:

Names are turned from camel_case to CamelCase for export.
There are no methods on v to set fields; just treat them as structure fields.
There are getters that return a field's value if set, and return the field's default value if unset. The getters work even if the receiver is a nil message.
The zero value for a struct is its correct initialization state. All desired fields must be set before marshaling.
A Reset() method will restore a protobuf struct to its zero state.
Non-repeated fields are pointers to the values; nil means unset. That is, optional or required field int32 f becomes F *int32.
Repeated fields are slices.
Helper functions are available to aid the setting of fields. Helpers for getting values are superseded by the GetFoo methods and their use is deprecated. msg.Foo = proto.String("hello") // set field
Constants are defined to hold the default values of all fields that have them. They have the form Default_StructName_FieldName. Because the getter methods handle defaulted values, direct use of these constants should be rare.
Enums are given type names and maps from names to values. Enum values are prefixed with the enum's type name. Enum types have a String method, and a Enum method to assist in message construction.
Nested groups and enums have type names prefixed with the name of the surrounding message type.
Extensions are given descriptor names that start with E_, followed by an underscore-delimited list of the nested messages that contain it (if any) followed by the CamelCased name of the extension field itself. HasExtension, ClearExtension, GetExtension and SetExtension are functions for manipulating extensions.
Oneof field sets are given a single field in their message, with distinguished wrapper types for each possible field value.
Marshal and Unmarshal are functions to encode and decode the wire format.


---

# Protocol buffers: обратная совместимость

- you must not change the tag numbers of any existing fields.
- you may delete fields.
- you may add new fields but you must use fresh tag numbers (i.e. tag numbers that were never used in this protocol buffer, not even by deleted fields).

---


# Домашнее задание

Реализовать утилиту envdir на Go.
<br><br>

Эта утилита позволяет запускать программы получая переменные окружения из определенной директории.
Пример использования:

```
go-envdir /path/to/env/dir some_prog
```

Если в директории /path/to/env/dir содержатся файлы
* `A_ENV` с содержимым `123`
* `B_VAR` с содержимым `another_val`
То программа `some_prog` должать быть запущена с переменными окружения `A_ENV=123 B_VAR=another_val`

---

# Опрос

.left-text[
Заполните пожалуйста опрос
<br><br>
[https://otus.ru/polls/3938/](https://otus.ru/polls/3938/)
]

.right-image[
![](img/gopher7.png)
]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
