Домашнее задание №3
Студент - Цокало Жан

Наговнокодил конечно, но вроде всё работает, как по ТЗ. Даже красиво работает.
Неплохо было бы конечно причесать код, возможно раскидать его по разным пакетам, но уже времени нет и очень много сил отняло это ДЗ.
Где-то в половине случаев стабильно куда-то теряется один успешный запрос вне зависимости от параметров запуска утилиты. Так и не смог понять почему. Пробовал блокровать счётчик успешных запросов через mutex, но результат остался таким же.
HTTP-сервер и утилиту запускал на разных машинах. Сервер на ноутбуке (он слабее), утилиту на стационарном компьютере (он мощнее) в надежде, что процессор ноутбука не выдержит нагрузки и перестанет отвечать на часть запросов. Однако добиться потери пакетов так и не удалось. Больше 1700 запросов в секунду отправлять не удалось при любом количестве потоков. Скорее всего это связано с тем что процессор 2-х ядерный. Скриншоты по ссылкам:
https://mega.nz/file/0RB22Jha#HDB5bmXJVWv4OgvPDz16ggwn-FBcpwr_j4_w8TOjgnA
https://mega.nz/file/lYYGWDhL#vKAF6nHnyQv8PQhBVGmkQX3jhjkyB0JOIXEj07sdZWw
Был бы рад получить обратную связь по следующим вопросам. В каких местах кода целесообразнее было использовать указатели или передавать указатели в функции вместо значений?
Какие переменые в каком месте кода оптимальнее было бы объявить?
Может нужно было структуры по-другому сформировать и их экземпляры, например, целиком передавать в функции?
Функцию sendHTTPRequest как оптимальнее было сделать, как сейчас сделано или как метод к структуре Worker?