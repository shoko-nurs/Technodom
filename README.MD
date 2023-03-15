1) Проверить код можно по ссылкам:
   
   Пользовательский метод
   http://technodom.shokonurs.codes:3334/redirects?url=z/x/c
   
   Методы управления данными

   а) Получить данные постранично
   http://technodom.shokonurs.codes:3334/admin/redirects?page=1&per_page=4
   
   b) Получить только 1 объект
   http://technodom.shokonurs.codes:3334/admin/redirects/20
   
   c) POST метод (с телом)
   http://technodom.shokonurs.codes:3334/admin/redirects
   
   d) PATCH метод (с телом)
   http://technodom.shokonurs.codes:3334/admin/redirects
   
   e) DELETE метод (с указанием метода)
   http://technodom.shokonurs.codes:3334/admin/redirects/20


2) База данных: AWS Postgres (free version)
   Думаю NoSql тут могла быть лучше, но я пока знаю только Postgres

3) DBFunction.txt
   Файл хранит функцию для метода GET /tdm/redirect

4) CreateTableScript.txt 
   Файл с описанием таблицы

5) Кеш пополняется/изменяется только для метода GET /tdm/redirect

6) Опциональные требования не выполнены