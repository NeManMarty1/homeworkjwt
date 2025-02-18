-- Создание роли (пользователя) с паролем
CREATE ROLE your_user WITH LOGIN PASSWORD 'your_password';

-- Создание базы данных
CREATE DATABASE your_database;

-- Предоставление прав на базу данных пользователю
GRANT ALL PRIVILEGES ON DATABASE your_database TO your_user;

-- Создание таблиц (пример)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL
);
