# API

# IP: wlcmtunknwndth.keenetic.pro:63342

## Auth

Для авторизации между страницами и элементами сайта используются jwt-токены. 
Токен один -- Access. Хранится, как Cookie под именем "access". 
Предусмотрена операция Refresh. Access выдается при успешных /register и /login и 
по умолчанию действует 4 минуты. Остальные запросы лишь проверют, есть ли 
токен в Cookie, и обновляют(Refresh) его при необходимости, если TimeToLive 
еще не вытек. Пользователи хранятся в базе данных. Структура User:

```Go
type User struct{
    Username string `json:"username"`
    Password string `json:"password"`
	Age      uint8  `json:"age"`
	Gender   bool   `json:"gender"` // True -- женщина, False -- Мужчина 
    isAdmin  boolean //структура не десериализуется.
```

Роль администратора нельзя задать через запрос. 
Роль задаётся только через базу данных. При входе по правам администратора 
информация присваивается через запрос от ф-ции к БД, а после кодируется
в jwt-токен, поэтому права не пропадают.

### POST /login

{
    "username": "string_value",
    "password": "string_value"
}

### POST /register

{
    "username": "string_value",
    "password": "string_value"
}

### POST /logout

empty body {}. В данном случае лишь стирается access-токен.

### DELETE /delete_user
Для удаления пользователя должен быть jwt-токен администратора.

{ "username": "string_value" }

## EVENTS

Event JSON: 

```JSON
{
  "id": uint,
  "price": uint,
  "restrictions": uint,
  "date": "timestamp as string",
  "feature": "deaf",
  "city": "moscow",
  "address": "Malay Ordinka, 3",
  "name": "Mayhem",
  "img_path": "/path/to/event/image/folder",
  "description": "chainsaw gutsfuck"
}
```

"restriction" is an age restriction, where num means lower bound
"img_path" -- картинки для ивента будут в папке ./data/events/<id>, пронумерованной от 1
"feature" -- для каждой особенности будет свой номер в бд, но для фронта:
    1. disability - "проблемы с мобильностью"
    2. deaf - "проблемы с слухом"
    3. blind - "проблемы c зрением"
    4. neuro - "нейроотличия"


### GET /event?id=<uint>     (Работает)

```JSON
{
    "id": uint,
    "price": uint,
    "restrictions": uint,
    "date": "timestamp as string",
    "feature": "deaf",
    "city":"moscow",
    "address":"Malaya Ordinka, 3",
    "name": "Mayhem",
    "img_path": "/path/to/event/image/folder",
    "description": "chainsaw gutsfuck"
}
```

Заголовки:

200 -- OK
400 -- Bad request (ошибка с параметром запроса)
401 -- Unauthorized (пользователь не авторизован)
403 -- Not enough permissions (пользователь не имеет прав)
500 -- Internal server error (внутренняя ошибка)

### GET /events?filter=["deaf", "blind" or other features]&ordering=["date", "price"]&order=["ascending", "descending"]&location=['moscow', etc] (Требует доработки)

```JSON
{
  "events": [
    {
      "id": uint,
      "price": uint,
      "restrictions": uint,
      "date": "timestamp as string",
      "feature": "deaf",
      "city":"moscow",
      "address":"Malaya Ordinka, 3",
      "name": "Mayhem",
      "img_path": "/path/to/event/image/folder",
      "description": "chainsaw gutsfuck"
    }
   ]       
}
```

### POST /create_event (РАБОТАЕТ)

```JSON
{
    "price": uint,
    "restrictions": uint,
    "date": "timestamp as string",
    "feature": "deaf",
    "city":"moscow",
    "address":"Malaya Ordinka, 3",
    "name": "Mayhem",
    "description": "chainsaw gutsfuck"
}
```

Img_Path представлен в виде /data/events/<id> . Предполагаю, что фотки при
загрузке будем просто нумеровать и доставать по очереди(1.jpeg, 2.jpeg и так далее)

При создании возвращается id нового ивента.

в виде ```Event created: id: 1```

id присваивается автоматически для того, чтобы не возникало
проблем

Заголовки:

201 -- Event created
400 -- Bad request(неправильный json)
401 -- Unauthorized (пользователь не авторизован)
403 -- Not enough permissions (пользователь не имеет прав)
500 -- Internal server error(ошибка на сервере)

### PATCH /patch_event (РАБОТАЕТ)

```JSON
{
    "id": uint,
    "price": uint,
    "restrictions": uint,
    "date": "timestamp as string",
    "feature": "deaf",
    "city":"moscow",
    "address":"Malaya Ordinka, 3",
    "name": "Mayhem",
    "description": "chainsaw gutsfuck"
}
```

200 -- EventPatched
400 -- Bad request(неправильный json)
401 -- Unauthorized (пользователь не авторизован)
403 -- Not enough permissions (пользователь не имеет прав)
500 -- Internal server error(ошибка на сервере)

### POST /delete_event?id=<id> (РАБОТАЕТ)

```JSON
{
    "price": uint,
    "restrictions": uint,
    "date": "timestamp as string",
    "feature": "deaf",
    "city":"moscow",
    "address":"Malaya Ordinka, 3",
    "name": "Mayhem",
    "description": "chainsaw gutsfuck"
}
```

Заголовки:

200 -- Deleted
400 -- Bad Request (неправильный параметр)
401 -- Unauthorized (пользователь не авторизован)
403 -- Not enough permissions (пользователь не имеет прав)
500 -- Internal server error (внутреняя ошибка)
