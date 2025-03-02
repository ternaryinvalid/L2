# Задания по программированию на Go

В этом репозитории собраны решения различных задач на языке программирования Go. Задания охватывают паттерны проектирования, работу с внешними библиотеками, взаимодействие с операционной системой, а также реализацию утилит и серверов. Решения ориентированы на улучшение навыков работы с конкурентностью, каналами, структурированием данных и взаимодействием с ОС.

## Выполненные задания

### Паттерны проектирования
1. **Паттерн «Фасад»**  
   Объяснение паттерна, пример использования для упрощения сложных интерфейсов.
   
2. **Паттерн «Строитель»**  
   Применение для пошагового создания объектов с разными конфигурациями.

3. **Паттерн «Посетитель»**  
   Пример использования для добавления новых операций к объектам без изменения их классов.

4. **Паттерн «Комманда»**  
   Применение для инкапсуляции запросов как объектов.

5. **Паттерн «Цепочка вызовов»**  
   Пример организации цепочек обработки запросов с динамическим выбором обработчиков.

6. **Паттерн «Фабричный метод»**  
   Реализация для создания объектов через абстракцию, скрывающую детали создания.

7. **Паттерн «Стратегия»**  
   Применение для выбора алгоритмов на основе контекста выполнения.

8. **Паттерн «Состояние»**  
   Пример реализации для изменения поведения объекта в зависимости от его состояния.

### Задачи на разработку

1. **Программа с использованием NTP библиотеки**  
   Программа, печатающая точное время с использованием библиотеки `github.com/beevik/ntp`, обрабатывает ошибки и возвращает ненулевой код ошибки при сбоях.

2. **Функция для распаковки строки с повторяющимися символами**  
   Написана функция, которая распаковывает строки с повторяющимися символами/рунами, а также поддерживает escape-последовательности и корректную обработку ошибок.

3. **Утилита sort**  
   Утилита для сортировки строк в файле, поддерживает различные ключи для сортировки по числовым значениям, по колонкам и в обратном порядке.

4. **Поиск анаграмм по словарю**  
   Функция для поиска множества анаграмм в словаре, с фильтрацией и сортировкой.

5. **Утилита grep**  
   Реализация фильтрации строк с поддержкой множества ключей, таких как `-A`, `-B`, `-C`, `-c`, `-i`, и других.

6. **Утилита cut**  
   Утилита для извлечения колонок из строк с использованием пользовательских разделителей и поддержкой различных ключей.

7. **Or канал**  
   Реализована функция, которая объединяет несколько каналов `done` в один канал, закрывающийся, когда любой из исходных каналов закроется.

8. **Шелл-утилита**  
   Написан простой UNIX-шеел с поддержкой команд `cd`, `pwd`, `echo`, `kill`, `ps`, а также конвейеров команд с использованием `fork/exec`.

9. **Утилита wget**  
   Утилита для скачивания сайтов целиком, реализована с использованием HTTP-запросов.

10. **Простейший Telnet-клиент**  
   Написан telnet-клиент для подключения к серверу по TCP и взаимодействия с ним через стандартный ввод/вывод.

11. **HTTP-сервер для работы с календарем**  
   Реализован HTTP-сервер с REST API для создания, обновления, удаления событий, а также получения событий по дням, неделям и месяцам. Включает middleware для логирования и обработку ошибок с соответствующими HTTP-кодами.

### Инструкции по запуску

1. **Клонировать репозиторий:**
   ```bash
   git clone https://github.com/username/repository-name.git
   cd repository-name
