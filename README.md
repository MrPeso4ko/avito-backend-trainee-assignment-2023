# avito-backend-trainee-assignment-2023

Задание выполнено Сальниковым П.Д.

### Инструкции по запуску:

- Для работы сервиса нужна работающая БД postgres. Для локального запуска её можно запустить в `docker`
- Прописать переменные окружения в файле `config/.env`
  (пример заполненного файла в `config/.env.example`):

```
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DB=postgres
POSTGRES_USER=segments_manager
POSTGRES_PASSWORD=1234

API_HOST=0.0.0.0
API_PORT=9053

MODE=prod
LOG_LEVEL=info
```

- прописать в файле `local/docker-compose.yml` на 11 строке тот же порт, что и прописан в переменной API_PORT,
  если он отличается от 9053
- запустить контейнер командой

`docker compose -f local/docker-compose.yml up`

### Список методов API:

### Создание сегмента:

```
POST /segments
```

Payload запроса (JSON):

```json
{
  "segment_name": "<название сегмента>"
}
```

Пример ответа:

```json
{
  "error": false,
  "message": "Segment created successfully"
}
```

### Удаление сегмента:

```
DELETE /segments
```

Payload запроса (JSON):

```json
{
  "segment_name": "<название сегмента>"
}
```

Пример ответа:

```json
{
  "error": false,
  "message": "Segment deleted"
}
```

### Запрос сегментов пользователя:

```
GET /users?user_id=<id пользователя>
```

Пример ответа:

```json
{
  "error": false,
  "segments": [
    {
      "segment_name": "SEGMENT1"
    },
    {
      "segment_name": "SEGMENT2"
    }
  ]
}
```

### Добавление пользователя в сегменты:

```
PUT /users
```

Payload запроса (JSON):

```json
{
  "segments": [
    {
      "segment_name": "<название сегмента...>"
    },
    ...
    // перечисление сегментов в таком же формате 
  ]
}
```

Пример ответа:

```json
{
  "error": false,
  "message": "Added user to segments"
}
```

### Удаление пользователя из сегментов:

```
DELETE /users
```

Payload запроса (JSON):

```json
{
  "segments": [
    {
      "segment_name": "<название сегмента...>"
    },
    ...
    // перечисление сегментов в таком же формате 
  ]
}
```

Пример ответа:

```json
{
  "error": false,
  "message": "Removed user from segments"
}
```
