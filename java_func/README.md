# Информация о рантайме Java 11

## Библиотеки и инструменты

Рантайм Java 11 использует следующие библиотеки и инструменты:

* Maven 3.6.3
* Java SDK - OpenJDK 11

## Обработчик функции

Для корректной работы функции, в ней должен быть [обработчик](https://developers.sber.ru/docs/ru/platform-v/functions/reference/develop-java/handler), имплементирующий интерфейс `ru.sber.platformv.faas.api.HttpFunction`.

Укажите класс и метод обработчика в [конфигурации функции](https://developers.sber.ru/docs/ru/platform-v/functions/how-to/configure-function#среда-исполнения).

Файлы для имплементации дополнительной бизнес-логики функции должны быть расположены в папке `src/main/java/`.

## Работа с зависимостями

Внешние зависимости могут быть определены в [pom.xml](./pom.xml) проекта функции. Локальные зависимости следует поместить в директорию `libs` в корне проекта функции, и так же определить их в `pom.xml`, указав параметры `<scope>system</scope>` и `<systemPath>путь-до-.jar-файла</systemPath>`. 

Подробнее смотрите в [документации](https://developers.sber.ru/docs/ru/platform-v/functions/reference/develop-java/dependencies).

## Тестирование и локальная отладка

Для тестирования при локальной разработке функций вы можете воспользоваться SDK. Подробные инструкции по его настройке и использованию расположены в [документации](https://developers.sber.ru/docs/ru/platform-v/functions/reference/develop-java/local-dev).

## Взаимодействие с DataSpace

Для работы с DataSpace вам понадобится DataSpace SDK для подключенной к вашей функции БД. Подробнее о получении и работе с SDK - в [документации на DataSpace](https://developers.sber.ru/docs/ru/platform-v/dataspace/graphql/sdk).

Запросы к DataSpace строятся при помощи GraphQL. Подробные инструкции приведены [в документации DataSpace](https://developers.sber.ru/docs/ru/platform-v/dataspace/graphql/overview).

Вы также можете ознакомиться с [примером использования DataSpace в функции](https://developers.sber.ru/docs/ru/platform-v/example/example-dataspace-sdk).
