# Диплом:
# Разработка и оптимизация приложений для электронной комерции

## Стек технологий

### Сервер

- Node.js - платформа
- MongoDB - NoSQL документориентированная база данных
- NestJS - Архитектурный фреймворк
- Redis - кэш
- MinIO - файловое хранилище
- Docker - система виртуализации

- JWT авторизация
- Ролевая модель (user, admin)

### Клиент

- VueJS 3.0 - фронтенд фреймворк
- ElementUI-Plus - библиотека компонентов на основе Vue 3.0
- VueRouter - 

## Список литературы

- https://nodejs.org/en/
- https://ru.wikipedia.org/wiki/MongoDB
- https://vuejs.org/
- https://element-plus.org/en-US/
- https://redis.io/
- https://min.io/
- https://www.docker.com/

### Backend go

- https://github.com/ilyakaznacheev/cleanenv - config
- https://github.com/mongodb/mongo-go-driver - mongo
- https://github.com/gin-gonic/gin - gin
- https://github.com/gin-contrib/cors - cors
- https://github.com/golang-jwt/jwt - jwt

## Сущности

- User
- Product
- Order
- Cart

## Модель данных

```
type User = {
    id: string;
    name: string;
    phone: string;
    address: string;
    email: string;
    role: admin | user
    password: string;
} 

type Product = {
    id: string;
    name: string;
    description: string;
    price: number;
    images: string[];
    count: number;
}

type Order = {
    id: string;
    user_id: string;
    products: string[];
}

type Cart = {
    id: string;
    user_id: string;
    products: string[];
}
```

## Методы

Auth:
- POST /signup - регистрация пользователя ✅
- POST /signin - вход в аккаунт ✅

User:
- GET  /user/ - получить пользователя ✅

Product:
- GET    /product/ - получение списка продуктов
- GET    /product/:product_id - получение детальной карточки продукта
- POST   /product/ - создание нового продукта
- DELETE /product/:product_id - удаление продукта

Order:
- GET /order/ - получение списка заказов пользователя
- GET /order/:order_id - детальная карточка заказа
- POST /order/ - создание заказа

Cart:
- GET /cart/ - корзина пользователя
