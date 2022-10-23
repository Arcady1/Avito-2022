Стек:
1. JSON формат как при отправке запроса, так и при получении результата.
2. БД: PostgreSQL.
3. Использовать docker-compose.

Optional:
1. Покрыть код тестами.
2. Реализовать сценарий разрезервирования денег, если услугу применить не удалось.

Описание:
1. Метод начисления средств на баланс.
    Принимает id пользователя и сколько средств зачислить. 
**2. Метод резервирования средств с основного баланса на отдельном счете.
    Принимает id пользователя, ИД услуги, ИД заказа, стоимость.**
**3. Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии. 
    Принимает id пользователя, ИД услуги, ИД заказа, сумму.**
4. Метод получения баланса пользователя. 
    Принимает id пользователя.

В конце:
1. Написать Readme файл с инструкцией по запуску и примерами запросов/ответов.


FLOW
Резервирование денег:
   1. Пользователь захотел что-то купить -> запрос на резервирование денег.
   2. Зарезервированная сумма списывается с баланса и сохраняется в orders.
Признание выручки:
   1. Если указанная сумма больше, чем заразервированная, то происходит разрезервирование денег:
      - order.status -> cancled
      - order.amount -> + account.balance
      - order.amount -> 0
   2. Если указанная сумма равна или меньше, чем заразервированная, то происходит списание денег:
      - order.status -> succeed
      - order.amount -> - INPUT_amount
      - order.amount -> + account.balance

Указать в README
   1. Как поступил с FLOW признания выручки.