# Тестовое задание на позицию стажёра-бэкендера
В реализации используются:
* Gin для маршрутизации запросов
* MySQL база данных
* Gorm для работы с БД из приложения

## **Основное задание**
Запуск проекта<br>
```bash
git clone https://github.com/stepann0/balance-microservice.git
cd balance-microservice
```
Освободите порты `:8080` и `:3306`, затем
```bash
docker-compose up
```
Запустите новую сессию терминала, затем:
```bash
cd balance-microservice
```
**Чтобы "прогнать" все нижеприведенные запросы**
```bash
bash requests.txt
```
Создание пользователей:<br>
```bash
curl -X PUT http://localhost:8080/ --data '{"name": "Alina", "balance": 5000}'
```
```json
{
  "account": {
    "id": 1,
    "name": "Alina",
    "balance": 5000
  }
}
```
Начисление средств пользователям:<br>
```bash
curl -X POST http://localhost:8080/increase --data '{"id":1, "amount": 5000}'
```
```json
{
  "amount": 5000,
  "balance": 10000,
  "status": "OK"
}
```
Резервирование средств с основного баланса:<br>
```bash
curl -X POST http://localhost:8080/reserve --data '{"user_id":1,"service_id":2,"amount": 2000}'
```
```json
{
  "reserve_amount": 2000,
  "user_balance": 8000
}
```
Признание выручки:<br>
```bash
curl -X PUT -i http://localhost:8080/accept --data '{"user_id":1,"service_id":2,"amount": 2000}'
```
```
HTTP/1.1 200 OK
Date: Tue, 01 Nov 2022 09:52:17 GMT
Content-Length: 0
```
Отклонение платежа:<br>
```bash
curl -X POST http://localhost:8080/reserve --data '{"user_id":1,"service_id":2,"amount": 2000}'
curl -X PUT http://localhost:8080/decline --data '{"user_id":1,"service_id":2,"amount": 2000}'
```
```
HTTP/1.1 200 OK
Date: Tue, 01 Nov 2022 10:59:19 GMT
Content-Length: 0
```
Получение баланса пользователя:<br>
```bash
curl -X GET -i http://localhost:8080/1
```
```json
{
  "account": {
    "id": 1,
    "name": "Alina",
    "balance": 8000
  }
}
```
## **Дополнительные задания**
Месячный отчет в бухгалтерию:<br>
```bash
curl -X GET -i http://localhost:8080/report/11/2022
```
```
консультация;2100.000
ремонт;2399.000
покупка;1597.000
доставка;6398.000
```
