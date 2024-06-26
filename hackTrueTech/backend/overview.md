# Обзор бэкенда проекта для хакатона True Tech Hack 
# От МТС

## Применяемые технологии:
1. Postgresql
2. NATS
3. Golang
4. Go-cache

## Уровни: 

1. Postgresql

В данной БД хранятся все необходимые данные: список событий 
и данные пользователей. Сама БД работает в контейнере.
Подробнее о хранении информации можно узнать из файла 
api.md

2. Файловый Сервер

Файловый сервер выполняет функцию хранилища фотографий события, которые используются со стороны frontend'а.

3. NATS

В данном случае NATS выполняется функцию брокера сообщений, что
позволяет масштабировать сервис, так как все операции
с БД выполняются только через брокер, поэтому при 
необходимости можно добавить достаточное кол-во
копий БД, чтобы убрать зависимость от единственного образа.
NATS также работает в контейнере.

4. Сервис Кэширования

Сервис кэширования в данном случае позволяет оптимизровать SLO, чтобы пользователь не замечал задержек при получении большого кол-ва событий.
Кэш восстанавливается из БД после падения сервиса при его очередном запуске, а также каждые 5 минут делается backup.  

5. HTTP-сервер

HTTP-сервер выполняет основные функции: передача
событий и авторизация пользователей. HTTP-сервер
предусматривает работу с CORS, что позволяет работать с 
браузером настоящего пользователя. HTTP-сервер также работает
в контейнере. Для частых запросых одного и того же события предусмотрено кэширование, чтобы не нагружать Базу данных.

## Основная задача: поиск событий по определенным ограничениям.

Для решения данной задачи мы создали две дополнительные таблицы: одна таблица является индексной, то есть хранит id события и массив из 
предусмотренных ограничений, а вторая содержит список возможных ограничений. В данном случае сервис может быть внедрен в реальную Базу Данных MTC Live так, чтобы не пришлось изменять уже имеющиеся
таблицы. Таким образом, достаточно добавить наши две таблицы для того, чтобы можно было уже изменить реальные события с учетом, добавляя туда
информацию о доступности для разных категорий граждан. Обе таблицы в нашем случае индексируются по B-tree, поэтому среднее время ожидания запроса
получения событий ~70 мс.

## Возможность настройки.

Наш сервер позволяет изменить используемую БД(в том числе и возможна смена типа БД), как и брокер, что позволяет
подстраивать нашу систему к уже имеющемуся стеку применяемых технологий.

## Переносимость.

В данном случае все уровни контейнеризированы, что
позволяет развернуть сервер там, где работает 
докер. 

## Масштабируемость.

Как уже было сказано ранее, благодаря брокеру 
сервис возможно масштабировать по кол-ву копий БД, а также
повторять такие операции, как создание пользователей и 
события, для каждой из копий.





