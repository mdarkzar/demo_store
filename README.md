**Это демонстрационный мини склад, который имеет:**
 
* Авторизацию
* Создание/удаление продукта
* Отображение уведомлений
* Отдельный микросервис для сохранения массовых уведомлений

API:

* Golang 1.2.0
* База данных: Postgresql 15.3
* Организация кода: чистая архитектура (собственная версия реализации на языке Golang)
* AMQP: RabbitMQ 3.12 (для общения между микросервисами)
* Контейнер для AMQP: Protobuf

Frontend:

* Vue 3
* Pinia
* Nginx

Собирается всё в Docker.

Мобильный клиент:

* Flutter 3.10+
* Getx 4.6.5

Запуск:

* make
* [http://localhost](http://localhost)

Тесты:

* После запуска, сделать make test

[Демо](https://demo-store.darkzar.uz)

* Логин любой и пароль - любой, действует авторегистрация, но логин должен начинаться с demo, например demo1, demo2 и тд.

[Мобильная версия для Android](https://github.com/mdarkzar/demo_store/raw/main/mobile/demo_store_v1.apk)
