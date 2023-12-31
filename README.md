# Region Todo List
Этот микросервис реализует функциональность управления списком задач.

# Требования к запуску
### Предварительные требования
-  Установленный Docker
- Установленный Docker Compose
## Установка и Запуск приложения

1. Клонируйте репозиторий проекта на свой компьютер:
```shell
git clone https://github.com/ymoldabe/Todo_list
```
2. Перейдите в директорию проекта:
```shell
cd ./Todo_list
```
3. Запустите приложение с помощью Docker Compose:
```shell
make run
```
или
```shell
docker-compose up
```
4. Чтобы остановить приложение, выполните команду:
```shell
make stop
```
или
```
docker-compose down
```


### API Endpoints
#### ** Формат обмена данными JSON.**

## Создание задачи

1. Метод: POST
- URL: /api/todo-list/task
- Тело запроса:

```json
{
   "title": "Купить книгу",
   "activeAt": "2023-08-04"
}
```
- Создание новой задачи

## Обновление задачи

2. Метод: PUT
- URL: /api/todo-list/task:id
- Тело запроса:
```json
{
   "title": "Купить книгу - Высоконагруженные приложения",
   "activeAt": "2023-08-05"
}
```
- Обновление существующей задачи.

## Удаление задачи

3. Метод: DELETE
- URL: /api/todo-list/task:id

   - Удаление задачи.


## Пометить задачу выполненной

4. Метод: PUT
- URL: /api/todo-list/task:id/done

  -  Помечает задачу как выполненную.

## Список задач

5. Метод: GET
- URL: /api/todo-list/task

  -  Получает список задач.
