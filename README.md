# effectiveMobTask
1 Сервис слушает очередь кафки FIO, в котором приходит информация с ФИО в
формате
{
"name": "Dmitriy",
"surname": "Ushakov",
"patronymic": "Vasilevich" // необязательно
}

2 В случае некорректного сообщения, обогатить его причиной ошибки (нет
обязательного поля, некорректный формат...) и отправить в очередь кафки
FIO_FAILED

3 Корректное сообщение обогатить
1 Возрастом - https://api.agify.io/?name=Dmitriy
2 Полом - https://api.genderize.io/?name=Dmitriy
3 Национальностью - https://api.nationalize.io/?name=Dmitriy

4 Обогащенное сообщение положить в БД postgres (структура БД должна быть создана
путем миграций)

5 Выставить rest методы
1 Для получения данных с различными фильтрами и пагинацией
2 Для добавления новых людей
3 Для удаления по идентификатору
4 Для изменения сущности

6 Выставить GraphQL методы аналогичные п.5

7 Предусмотреть кэширование данных в redis

8 Покрыть код логами

9 Покрыть бизнес-логику unit-тестами

10 Вынести все конфигурационные данные в .env

Выполнены:
Пункты 5, 8, 9, 10
База данных создаётся путем миграции, миграции прописаны в Makefile, для базы данных использовался image докер с postgres:15-alpine.
Создание контеинера так же присутствует в Makefile, нужно только загрузить image базы данных. 
  
