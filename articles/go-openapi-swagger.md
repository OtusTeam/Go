# Использование Swagger в Go

В предыдущей статье я сделал короткое введение в спецификацию Open API и показали некоторые инструменты для работы с ней. Сейчас я более подробно рассмотрим go-swagger - утилиту для генерации Go кода из swagger файлов.

## go-swagger - кодогенрация для Go

[go-swagger](https://github.com/go-swagger/go-swagger) это инструмент для Go разработчиков, позволяющий автоматически генерировать Go код по swagger файлам. Он опирается на различные библиотеки из [проекта go-openapi](https://github.com/go-openapi) для работы с форматом swagger.

Я некоторое время следил за проектом. Он развивается очень активно, коммиты добавляются в master-ветку каждые несколько дней. Основные контрибьюторы очень быстро реагируют на возникающие проблемы. Проект ввыпускается релизами с конкретными версиями и поставляется в виде исполняемых файлов или docker-контейнеров.

## Пример

В первую очередь установить команду `swagger` [по инструкции с github](https://github.com/go-swagger/go-swagger/blob/master/docs/install.md). Далее мы будем использовать пример `swagger.yaml` файла из предыдущего поста.

## Создание REST сервера

Используйте следущие bash команды что бы создать и запустить сервер для вашего swagger файла. Единственные требование - это наличие swagger.yaml файла в текущей рабочей директории, и то что эта директория находится внутри `GOPATH`.

    $ # Validate the swagger file 
    $ swagger validate ./swagger.yaml
    The swagger spec at "./swagger.yaml" is valid against swagger specification 2.0
    $ # Generate server code
    $ swagger generate server
    $ # go get dependencies, alternatively you can use `dep init` or `dep ensure` to fix the dependencies.
    $ go get -u ./...
    $ # The structure of the generated code
    $ tree -L 1
    .
    ├── cmd
    ├── Makefile
    ├── models
    ├── restapi
    └── swagger.yaml
    $ # Run the server in a background process
    $ go run cmd/minimal-pet-store-example-server/main.go --port 8080 &
      09:40:12 Serving minimal pet store example at http://127.0.0.1:8080
    $ # go-swagger serves the swagger scheme on /swagger.json path:
    $ curl -s http://127.0.0.1:8080/swagger.json | head
      {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
    $ # Test list pets
    $ curl -i http://127.0.0.1:8080/api/pets
    HTTP/1.1 501 Not Implemented
    Content-Type: application/json
    Content-Length: 50

    "operation pet.List has not yet been implemented"
    $ # Test enforcement of scheme - create a pet without a required property name.
    $ curl -i http://127.0.0.1:8080/api/pets \
        -H 'content-type: application/json' \
        -d '{"kind":"cat"}'
    HTTP/1.1 422 Unprocessable Entity
    Content-Type: application/json
    Content-Length: 49

    {"code":602,"message":"name in body is required"}

Все должно заработать автоматически! go-swagger создал следущие директории

* `cmd` - команда для запуска сервера, обработка параметров командной строки и конфигурации.
* `restapi` - логика маршрутизации запросов на основе секции `paths` в swagger файле.
* `models` - модели из раздела `definitions` в swagger файле.

## Создание REST клиента

Давайте попробуем создать клиент к нашему серверу

    $ swagger generate client

И напишем небольшую программу, которая использует сгенерированный клиент

    package main

    import (
        "context"
        "flag"
        "fmt"
        "log"

        "github.com/posener/swagger-example/client"
        "github.com/posener/swagger-example/client/pet"
    )

    var kind = flag.String("kind", "", "filter by kind")

    func main() {
        flag.Parse()
        c := client.Default
        params := &pet.ListParams{Context: context.Background()}
        if *kind != "" {
            params.Kind = kind
        }
        pets, err := c.Pet.List(params)
        if err != nil {
            log.Fatal(err)
        }
        for _, p := range pets.Payload {
            fmt.Printf("\t%d Kind=%v Name=%v\n", p.ID, p.Kind, *p.Name)
        }
    }

Когда мы запустим ее, мы получим ошибку HTTP 501:

    $ go run main.go 
      15:57:53 unknown error (status 501): {resp:0xc4204c2000}
    exit status 1

## Реализация метода API

Как вы можете видеть 
