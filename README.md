**Test task**

*Start*

``docker-compose up``

*Using*

`curl http://localhost:3333/api/v1/places/mow`
Вернет массив

`curl http://localhost:3333/api/v1/places/mow?locale=ru`
Вернет локализованный массив

`curl http://localhost:3333/admin/drop`
Вернет 401

`curl http://admin:admin@localhost:3333/admin/drop`
Уронит кэш