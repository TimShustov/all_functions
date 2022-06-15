import logging

# Метод handler. Данный метод будет обрабатывать HTTP запросы поступающие к функции
# Принимает (flask.Request): HTTP request object.
# Возвращает текст ответа или любой набор значений, которые можно превратить в
# Response object используя
# `make_response <https://flask.palletsprojects.com/en/1.1.x/api/#flask.Flask.make_response>`.
def handler(request):
    # Логирование входящего запроса
    logging.info("Request received: %s\nMethod: %s", request.data.decode(), request.method)

    # Подготовка и возврат ответа на вызов
    return (
        "Hello from Python function!\n" +
        "You said: " + request.data.decode(),
        200,
        {"Content-Type": "text/plain"}
    )