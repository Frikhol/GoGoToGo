# test-project

1. Написать сам скелет GRPC сервера. Сделать логгер, конфиг, подготовить место для всех слоев. Сделать проверочную ручку health которая ничего не принимает и ничего не отдает, нужна для проверки работоспособности сервера

2. Сделать мини-сервис редиректов. Сервис взаимодействует с Postgres, сохраняет данные туда. Сервис должен иметь 2 эндпоинта
    1. `POST Add` Добавляет ссылку в базу. Input:
        1. Link
        2. fakeLink
        3. Erase time
    2. `GET take` Принимает uri, выдает ссылку
3. Добавить кэширование через redis
4. Добавить grpc  gateway
5. Навесить метрики на ручки
6. Добавить ручку авторизации, сделать интерсептор на валидацию JWT
7. Ручка, которая принимает данные и кладет их в кафку. Также отдельный консюмер, который читает из Кафки и актами пишет в Кликхаус