# feeds-framework

Главные используемые сущности - Activity и Feed(набор Activity).
Также есть SortedFeed, которая поддерживает упорядочивание по некоторому ключу, который задается математическим выражением(поле key_formula) с ипользованием параметров события, причем даже тех, что указаны в extra_data.

В БД храним так:

feeds:\<feedID> -> feedData

sorted_feeds:\<feedID> -> feedData

activities:\<activityID> -> activityData

Далее для удобного получения событий из лент будем хранить индекс:

feed_activities:\<feedID>:\<activityID> -> activityID

sorted_feed_activities:\<feedID>:\<activityKey> -> activityID

Добавить одно и то же событие(с одним внутренним ID) в несколько лент можно с помощью параметра redirect_to в запросе AddActivity, событие добавляется в перечисленные ленты.
