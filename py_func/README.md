# Информация о рантайме Python 3.8

## Библиотеки и инструменты

В рантайме используются следующие библиотеки и инструменты:

* Python 3.8.6
* pip 19.3.1

## Обработчик функции

Для корректной работы в функции должен быть [обработчик](https://developers.sber.ru/docs/ru/platform-v/functions/reference/develop-py/handler) с функцией `handler` принимающей на вход `flask.Request` и возвращающей Response object. Укажите файл обработчика в [конфигурации функции](https://developers.sber.ru/docs/ru/platform-v/functions/how-to/configure-function#среда-исполнения).

## Работа с зависимостями

Внешние зависимости могут быть определены в [requirements.txt](./requirements.txt) проекта функции. Подробнее смотрите в [документации](https://developers.sber.ru/docs/ru/platform-v/functions/reference/develop-py/dependencies).

## Тестирование и локальная отладка

Unit-тесты для функций можно составлять так же, как и для других приложений на Python. Более подробные инструкции по составлению тестов приведены в [документации](https://developers.sber.ru/docs/ru/platform-v/functions/reference/develop-py/testing). 

При сборке функции в SmartMarket unit-тесты запускаться не будут.

Для запуска и тестирования функций при локальной разработке доступен SDK. Подробные инструкции по работе с ним смотрите в [документации](https://developers.sber.ru/docs/ru/platform-v/functions/reference/develop-py/local-dev).
