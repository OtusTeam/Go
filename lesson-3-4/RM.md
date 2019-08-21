.center.icon[![otus main][image-1]]

---

class: top white
background-image: url(img/sound.svg)
background-size: 130%
.top.icon[![otus main][image-2]]

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

class: white
background-image: url(img/message.svg)
.top.icon[![otus main][image-3]]

# gRPC

### Alexander Davydov

---

# План занятия

.big-list[
* Описание API с помощью Protobuf
* Генерация кода для GRPC клиента и сервера
* Реализация API
* Прямая и обратная совместимость API
* Представление о Clean Architecture
]

---

# Protocol buffers



---

# HTTP/2 vs HTTP

[https://imagekit.io/demo/http2-vs-http1][1]

# HTTP/2
   
- server push (vs long polling)
-  header compression
- multiplexing (several requests in one connection)

---

# Типы gRPC API



---

# Scalability

---

# Security (SSL)

---

# gRPC vs REST



# Итого: почему gRPC

---


# Ссылки

.big-list[
* [http://goog-perftools.sourceforge.net/doc/tcmalloc.html][2]
* [https://programmer.help/blogs/exploration-of-golang-source-code-3-realization-principle-of-gc.html][3]
* [https://blog.golang.org/ismmkeynote][4]
* !!! [https://about.sourcegraph.com/go/gophercon-2018-allocator-wrestling][5]
* [http://gchandbook.org][6]
]

---

# Опрос

.left-text[
Заполните пожалуйста опрос
<br><br>
[https://otus.ru/polls/4561/][7]

]

.right-image[
![][image-4]
]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main][image-5]]

# Спасибо за внимание!

[1]:	https://imagekit.io/demo/http2-vs-http1
[2]:	http://goog-perftools.sourceforge.net/doc/tcmalloc.html
[3]:	https://programmer.help/blogs/exploration-of-golang-source-code-3-realization-principle-of-gc.html
[4]:	https://blog.golang.org/ismmkeynote
[5]:	https://about.sourcegraph.com/go/gophercon-2018-allocator-wrestling
[6]:	http://gchandbook.org
[7]:	https://otus.ru/polls/4561/

[image-1]:	img/main.png
[image-2]:	img/logo.png
[image-3]:	img/logo.png
[image-4]:	img/gopher7.png
[image-5]:	img/logo.png
