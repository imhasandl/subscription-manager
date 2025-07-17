
## Запуск проекта

Рекомендуется запускам проект в отделном docker контейнере.

### Запуск с помощью Docker Compose

1.  Установите Docker и Docker Compose.
2.  Создайте файл `.env` в корневой директории проекта и заполните его следующими значениями либо сами установите конфигурацию вашей базы данных:

    ```env
    POSTGRES_DB=subscription_db
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=postgres
    DB_PORT=5433
    DB_URL=postgres://postgres:postgres@db:5432/subscription_db?sslmode=disable
    PORT=8080
    ```
3.  Запустите проект с помощью Docker Compose:

    ```bash
    docker compose up --build

    либо

    docker-compose up --build
    ```

## Описание API

### GET /subscription/{id}

*   Принимает: `id` подписки в формате UUID в URL.
*   Возвращает: JSON представление подписки с указанным `id`.

---

### POST /subscription

*   Принимает: JSON с данными для создания новой подписки:

    ```json
    {
        "service_name": "Название сервиса",
        "price_rub": 1000,
        "user_id": "UUID пользователя",
        "start_date": "Дата начала подписки",
        "end_date": "Дата окончания подписки"
    }
    ```

*   Возвращает: JSON представление созданной подписки.

---

### PUT /subscription/{id}

*   Принимает: `id` подписки в формате UUID в URL и JSON с данными для обновления подписки:

    ```json
    {
        "service_name": "Новое название сервиса"
    }
    ```

*   Возвращает: JSON представление обновленной подписки.

---

### DELETE /subscription/{id}

*   Принимает: `id` подписки в формате UUID в URL.
*   Возвращает: JSON с сообщением об успешном удалении:

    ```json
    {
        "status": "success"
    }
    ```

---

### GET /subscription/sum

*   Принимает: JSON с датой начала для расчета суммы подписок:

    ```json
    {
        "start_date": "Дата начала"
    }
    ```

*   Возвращает: JSON с общей суммой подписок, начинающихся с указанной даты:

    ```json
    {
        "total_sum": 1000
    }
    ```



## Примечания
- В проекте используется методология DRY (Don't Repeat Yourself) через файл `utils/json.go`, который содержит переиспользуемые функции `RespondWithJSON` и `RespondWithError` для стандартизированной обработки HTTP-ответов и ошибок во всех обработчиках.
