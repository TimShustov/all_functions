# Информация о рантайме NodeJS 16

## Библиотеки и инструменты

В рантайме используются следующие библиотеки и инструменты:

* NodeJS 16.14.0
* Npm 8.3.1

## Обработчик функции

Для корректной работы функция должна содержать [обработчик](https://developers.sber.ru/docs/ru/platform-v/functions/reference/develop-js/handler) с методом `handler`, принимающим объекты запроса **Request** и **Response** библиотеки Express.

Укажите файл обработчика в [конфигурации функции](https://developers.sber.ru/docs/ru/platform-v/functions/how-to/configure-function#среда-исполнения). Если в разработке используется typescript, то следует указать транспилированный файл.

Перед запуском функции, будет запущен скрипт `prepare-function`, указанный в `package.json`. В него можно добавить различные подготовительные команды, такие как `tsc` или `webpack`.

Если функции необходимо отправлять файлы (html, js, css, png и подобные), можно воспользоваться одним из следующих методов:
* использовать в обработчике метод `res.sendFile(path [, options] [, fn])`
* установить статические папки в массиве `staticContentFolders` в `package.json`, например: `"staticContentFolders": ["static", "public"])`, и загружать файлы через url, например: `http://localhost:8082/index.html`.

## Работа с зависимостями

Добавлять зависимости можно в блоке `dependencies` файла `package.json`. Подробнее смотрите в [документации](https://developers.sber.ru/docs/ru/platform-v/functions/reference/develop-js/dependencies).

## Тестирование и локальная отладка
Скрипт `npm run test` не запускается во время сборки функции. Поэтому, чтобы добавить тесты, укажите соответствующие команды для запуска
инструментов тестирования (jest, mocha и т.д.) в скрипте `prepare-function` файла [package.json]. Подробнее смотрите в [документации](https://developers.sber.ru/docs/ru/platform-v/functions/reference/develop-js/testing).

Для запуска и тестирования функций при локальной разработке доступен SDK. Подробные инструкции по работе с ним смотрите в [документации](https://developers.sber.ru/docs/ru/platform-v/functions/reference/develop-js/local-dev).

