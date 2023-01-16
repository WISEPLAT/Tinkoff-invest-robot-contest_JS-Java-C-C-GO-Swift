# TinkoffInvestBot
Tinkoff Invest robot manager

# Система управления счетами и роботами в Тинькофф Инвестициях 

Для работы приложения требуются токен авторизации на торговой платформе Тинькофф Инвестиции и токен для доступа к телеграм-боту, которые указываются в файле настроек _settings.ini_. Файл настроек может быть закодирован, если запустить программу с ключом _-e_ и передать пароль.

## Начало работы

## Зависимости

Для сборки проекта необходимо установить некоторые пакеты. Вы можете сделать это, выполнив следующие команды:

```sh
sudo apt-get install g++ make qt5-default binutils cmake libssl-dev libboost-system-dev zlib1g-dev libcurl4-openssl-dev
```

### Сборка

Клонируйте репозиторий:

```bash
git clone https://github.com/samoilovv/TinkoffInvestBot.git
cd TinkoffInvestBot
git submodule update --init --recursive
``` 

Перейдите в директорию проекта и выполните следующие команды:

```bash
mkdir build && cd build
cmake ..
make
``` 

### Описание

Благодаря данному приложению вы можете управлять своими инвестициями как вручную, так и с помощью различных торговых роботов, прямо из своего мессенджера Телеграм. С помощью системы плагинов вы можете самостоятельно создавать (или заказывать у сторонних разработчиков) неограниченное количество дополнительных роботов и получать управление через мессенджер. Просто скопируйте .so или .dll файл в папку robots и новый робот автоматически появится в списке меню чата управления. Создаваемые роботы должны поддерживать интерфейс, описанный в файте _robotinterface.h_, тогда все его основные функции станут доступны через телеграм.

## Скриншоты

![alt text](example1.png "Commands list")