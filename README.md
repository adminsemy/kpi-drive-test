## Тестовое задание для backend Golang kpi-drive

### Описание задачи
[https://docs.google.com/document/d/1UIK3JQBAhxjH1gW_DsVBjwjsZjbrPU9HMuFcBlfafug/edit#heading=h.vqwuk7lrs14m](https://docs.google.com/document/d/1UIK3JQBAhxjH1gW_DsVBjwjsZjbrPU9HMuFcBlfafug/edit#heading=h.vqwuk7lrs14m)

### Запуск программы

Необходимо клонировать себе текущий репозиторий

    git clone https://github.com/adminsemy/kpi-drive-test

Затем перейти в папку kpi-drive-test и запустить программу

    go run cmd/kpi-drive/main.go -buf=10 -resp=10

Либо скомилировать программу

    go build cmd/kpi-drive/main.go

А затем запустить программу

    ./main -buf=10 -resp=10

Для программы используется следующий параметры для командной строки

- -buf - размер буфера для данных. Если данных больше, чем размер буфера, то эти данные будут игнорироваться. Если параметр не задан будет по умолчанию 1000
- -resp - количество данных. Задается сколько данных нам пришло (эдакая эмуляция запросов от пользователей). Если параметр не задан будет по умолчанию 10

### Тестирование

Все тестирование проходит в командной строке и терминале - задаются нужные параметры буфера и данных и в терминале выводятся сообщения о том, какие данных добавились в буфер и о том, какие данные сохранились