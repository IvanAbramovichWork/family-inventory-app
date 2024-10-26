# Используем официальный образ PostgreSQL
FROM postgres:14

# Установим переменные окружения для базы данных
ENV POSTGRES_DB=mydatabase
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword

# Скопируем SQL скрипт с таблицами в контейнер
COPY init.sql /docker-entrypoint-initdb.d/
