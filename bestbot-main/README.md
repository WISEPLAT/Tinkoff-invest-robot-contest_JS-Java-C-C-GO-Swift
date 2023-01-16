# Торговая система "BestBot v.9000"
Данная сборка является примером торгового робота в рамках конкурса Тинькоф Инвестиций

https://github.com/Tinkoff/invest-robot-contest

# Зависимости
Данная сборка была протестирована в ОС Debian 11 и в NodeJs 16, но предполагается, что она будет работать и в других системах с установленным nodejs.
Так же могут потребоваться стандартные библиотеки:
```
sudo apt-get install gcc g++ make
sudo apt-get install build-essential
sudo apt install sqlite3
sudo apt-get install python
```



Используется NodeJS SDK для API Тинькофф Инвестиций GRPC собственной разработки: https://github.com/betslus1/unofficial-tinkoff-invest-api_v2-lazy-sdk-NODEJS

# Установка

### Скопировать

автоматически

`git clone --recurse-submodules git@github.com:betslus1/bestbot.git`

вручную

`Скопировать данный репозитарий кнопкой Download ZIP`

`Скопировать репозитарий https://github.com/betslus1/unofficial-tinkoff-invest-api_v2-lazy-sdk-NODEJS в папку lib`

### Установить зависимости
1) В консоли открыть папку проекта например `cd bestbot`
2) Установить библиотеки и скомпилировать командой `npm i`

# Настройка системы
Настройка системы осуществляется в файле `options.js`

В первую очередь необходимо прописать:
- token от системы Тинькоф Инвестиций.
- tradeAccount Номер аккаунта, где будут проходить торги
- Логин/Пароль для доступа к системе
- Токен для уведомления в telegram а так же ID вашего аккаунта (Подробности в файле настроек)
- Прочие настройки задокументированы в файле options.js (Так же вы можете добавлять свои настройки)

# Запуск торговой системы
`node app.js`

OR

`node app.js [token]`

# Функциональные возможности
- Получение/обработка Market-data в реальном времени по протоколу GRPC с серверов Тинькоф Инвестиций, в том числе стаканов/свечей/служебных сообщений
- Контроль балансов и заявок
- Возможность работы в различных торговых контурах (Реальный счет, Тестовый счет, Симуляция)
- Обработка ошибок, логирование работы торговой системы на каждом этапе, возможность просмаривать логи в удобном виде
- Уведомление в Telegram об различных событиях и ошибках
- Учет вашх сделок по методу FIFO, а так же рассчет сводной информации по инструментам.
- Возможность контролировать работу торговой системы через визуальный интерфейс
- Возможность интеграции собственных торговых стратегий и их тестирование
- Кэширование маркет даты
- Открытие/Закрытие тестовых счетов. Зачисление балансов.
- Удобная система визуализации работы системы в реальном времени
- Отслеживание статуса работы биржи
- Адаптация под мобильные экраны телефона веб версии


# Схема взаимодействия модулей приложения
<img src="https://habrastorage.org/webt/aq/zn/lj/aqznljwebankysnyhhxbmct82ho.png" />


# Торговая стратегия
### Примеры торговых стратегий
- Стратегия Random: "Покупаем в случайное время, Продаем в случайное время". Данная торговая стратегия разработана для тестирования системы (максимализация количества сделок)
- Стратегия RSI: Пример торговой стратегии, в которой робот принимает решение о покупке и продаже на основе рассчетных индикаторов

### Разработка своей стратегии
 Входные параметры: 
  - lastCandles - последние свечи (глубина задается в настройках)
  - currentOrder - список текущих выставленных заявок
  - currentBalance - текущий баланс
  
 Выходные параметры: 
  - logs - логи работы стратегии
  - commands - список приказов к торговой системе о выставлении сделок
  - indicators - рассчетные индикаторы на основе которых принималось решение

Разработка:
 - Разместите свою торговую стратегию в папке `/strategy`
 - Дополнительно можете указать описание индикаторов на основе которых принимаются решения
 - В вашей торговой стратегии должна быть основная функция обработки "Step". Функция Step принимает Входные параметры, проводит все нобходимые действия в рамках стратегии и возвращает в торговую систему принятое решение

Сокращенный пример торговой стратегии (Полноценный пример смотрите в примерах торговых стратегий)
```
module.exports.indicators = { 'RSI':{ 'min':0, 'max':100, 'buy':40, 'sell':60 }};
  
module.exports.step       = function (lastCandles, currentOrder, currentBalance) {
   let indicators['RSI'] = calc_RSI(lastCandles);
   if(indicators['RSI'] > module.exports.indicators['RSI']?.buy){
     commands.push({'type':'Buy', 'price':price.Buy, 'quantity': quantity.Buy});
     logs.push('Покупаем: ${price.Buy} x ${quantity.Buy}');
   }
   if(indicators['RSI'] < module.exports.indicators['RSI']?.sell){
     commands.push({'type':'Sell', 'price':price.Sell, 'quantity': quantity.sell});
     logs.push('Продаем: ${price.Buy} x ${quantity.Buy}');
   }
   return {logs, commands, indicators};
 }
```

# Интерфейс 
### Все параметры в обоих интерфейсах обновляются в реальном времени во всех режимах
 - `lib/consoleUI.js` - Набор модулей для управления консольным интерфейсом
 - `lib/Web.js` - Набор модулей для управления HTTP версией интерфейса по протоколам HTTP и WebSocket
 - Страницы Веб-версии находятся в папке `/views`, используется шаблонизатор EJS + express

### Описание консольного интерфейса
<img src="https://habrastorage.org/webt/e1/tw/lo/e1twloa5sf_1jtgomhhmzbi4p-y.gif" />

- Левая верхняя часть окна - Лог работы стратегии
- Нижняя левая часть окна - Лог работы торговой системы и глобальные события
- Правая часть окна - От верхнего к нижнему: Информация о инcтрументах и балансах, текущие заявки, свечи, стакан
- Обновление содержимого окон происходит с помощью функции `consoleUI.render('имя окна','содержимое')`

### Описание HTTP версии интерфейса
По умолчанию веб-версия доступна по адресу http://ваш_ip_адрес:333/ (demo:demo). Порт и логин/пароль меняются через options.js.
<img src="https://habrastorage.org/webt/75/w8/i0/75w8i0ngcpnwk80e4eav_2otuaa.png" />
В целом интерфейс использует то же самое расположение основных окон что и консольная версия. Добавлены функции управления, такие как выбор торгового контура, зачисления баланса, выбор инструмента и т.д.

### Процесс симуляции с визуализацией в реальном времени
<img src="https://habrastorage.org/webt/gh/y6/wq/ghy6wqgykkp3xg5g_riwm49kvac.gif" />

Обновление содержимого окон происходит с помощью функции `web.render('имя окна','содержимое')` по протоколу WebSocket

### Поддержка мобильных устройств
<img src="https://habrastorage.org/webt/fa/wf/vy/fawfvyzzqlac7epjcocz3_wts48.png" />

# Торговые контура
- Боевой режим - работа на реальном счете
- Песочница - работа в песочнице брокера
- Бэктестинг - работа на историческиз данных для проверки работы стратегии. Тестовый контур реализован в файле `lib/backtest.js` и использует основные методы основного контура для получения максимальной правдоподобности и достижения полиморфизма кода.

В контурах Песочница и Бектестинг доступны пополнения счета.
В контуре Бэктестинг доступна выгрузка исторических данных и выбор скорости работы симуляции. Визулизация в реальном времени.

# Логирование
Логирование реализовано через глобальную функцию "log" и описано в файле `lib/logger.js`
Параметры функции:
- Модуль <String> - модуль программы к которой относится регистрируемое событие
- Сообщение <String>- текст сообщения о событии
- Тип сообщения - тип регистрируемого события (debug - скрытое сохранение события, global - отображаем событие в консоли, error - отображаем и отправляем уведомление, увеличиваем счетчик ошибок)
- Вывод в глобальном логе <bool> - Дополнительный флаг отображение в глобальном логе
- Отправка уведомления <bool> - Отправка сообщения о событии в Telegram

 <img src="https://habrastorage.org/webt/cq/mr/6g/cqmr6gk0nc1ywuvbvwvfewwru8s.png" />
 
 # Статистика
 Статистика считается по методу FIFO, для этого ежеминутно запрашивается информация о завершенных сделках. Каждому трейду покупки ставится в соответствие трейд продажи и высчитывается уплаченная комиссия, прибыль. Так же доступно отображение сводной информации по каждому инструменту.
 <img src="https://habrastorage.org/webt/yc/wb/8b/ycwb8b2p0gm3iotc2_uda0fmwr0.png" />
 
 # Система хранения 
 Система хранения подключается в файле `lib/db.js`, по умолчанию используется Sqlite3,  но можно переопределить методы `query` и `select` для использования другой СУБД. Исторические свечи хранятся в папке `/cache`. Для очистки БД достаточно удалить файл `data/db.sqlite`, после перезапуска приложения его структура автоматически восстановится
 
 
# Отказ от ответственности
 Данная система поставляется, как есть. Автор не несет ответсвенности за возможные убытки. Торги на бирже являются высокорисковыми операциями требующими опрделенных навыков и опыта. Не является торговой рекомендацией.
 
 # Номинация "Самое оригинальное использование API Тинькофф Инвестиций"
 - Тип: Игра
 - Жанр: Симуляция, Аркада
 - Системные требования: Компьютер средней мощности
 - Дополнительные программы: Браузер
 - Дистрибутив: Поставляется в рамках торговой системы "bestBot", вкладка "DirectMod". 

 ### Описание
 Недалекое будущее. Ученые научились переносить разум в информационные системы. Ваш разум был перенесен и в результате ошибки затерялся где-то в глубинах API Тинькофф инвестиций. Теперь вы не человек, а биржевая заявка на бирже. Вас окружают трейдеры, которые хотят Вас поймать. Но Вы понимаете, что жизнь биржевой заявки продолжается, пока она не сработает с встречными заявками. Вы вынуждены скитаться по биржевым стаканам избегая заявок других трейдеров...

 <img src="https://habrastorage.org/webt/ni/oh/8p/nioh8pjd-xxxbutpzag6pfv_voc.gif" />
 
- В случае пересечения с биржевым стаканом на бирже срабатывает сделка. Если пересечение в верхней части стакана то рыночная продажа, в нижней части стакана - рыночная покупка.
- За каждый стакан в котором вас не смогли купить, вы получаете 5 копеек.
- Используются реальные рыночные данные транслируемые в реальном времени из API Тинькофф-инвестиций
- В случае повышения цены на рынке, окно пролета становится выше, или ниже в случае понижения цены
- Размер окна пролета измеряется текущим спредом на рынке
- Данная игра может рассматриваться как пример обработки стакана в режиме реального времени.
 
 ### Условные обозначения
* Amount - Текущий размер ваших денежных средств
* Shares - Текущее количество ваших ценных бумаг
* Profit - Ваша прибыль или убыток относительно стартового баланса

PS. Сервис доступен во время работы бирж. Уровень сложности определяется выбором торгового инструмента. Чем выше ликвидность и ниже спред, тем сложнее будет играть.