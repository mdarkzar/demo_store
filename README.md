**Это демонстрационный мини склад, который имеет:**
 
* Авторизацию
* Создание/удаление продукта
* Отображение уведомлений

API:

* Golang 1.2.0
* База данных: Postgresql 15.3
* Организация кода: чистая архитектура (собственная версия реализации на языке Golang)

Frontend:

* Vue 3
* Pinia
* Nginx

Собирается всё в Docker.

Запуск:

* make
* [http://localhost](http://localhost)

Тесты:

* После запуска, сделать make test

[Демо](https://demo-store.darkzar.uz)

* Логин любой и пароль - любой, действует авторегистрация, но логин должен начинаться с demo, например demo1, demo2 и тд.

[Мобильная версия для Android](https://github.com/mdarkzar/demo_store/raw/main/mobile/demo_store_v1.apk)
