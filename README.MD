1) На данной VM ссылка admin уже занята, поэтому , пожалуйста,
   используйте префикc tdm. 
   
   Например, GET /tdm/admin/redirects/{id}
             GET /tdm/redirect?url=some/url

2) База данных: AWS Postgres (free version)
   Думаю NoSql тут могла быть лучше, но я пока знаю только Postgres

3) DBFunction.txt
   Файл хранит функцию для метода GET /tdm/redirect

4) CreateTableScript.txt 
   Файл с описанием таблицы

5) Кеш пополняется/изменяется только для метода GET /tdm/redirect