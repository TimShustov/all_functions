// Метод handler. Данный метод будет обрабатывать HTTP запросы поступающие к функции
// Принимает объекты запроса Request ssи Response библиотеки Express (http://expressjs.com)
exports.handler = (req, res) => {
    // Логирование входящего запроса
    console.log(`Request received: ${JSON.stringify(req.body)}\nMethod: ${req.method}`);

    // Подготовка и возврат ответа на вызов
    res.status(200).send(`Hello from NodeJS function!\nYou said: ${JSON.stringify(req.body)}`);
};
