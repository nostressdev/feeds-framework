# feeds-framework

## Feeds and activities

Главные используемые сущности - Activity и Feed(набор Activity).

В БД храним так:

feeds:\<feedID> -> feedData

activities:\<activityID> -> activityData

Далее для удобного получения событий будем хранить индексы:

feed_activities:\<feedID>:\<activityID> -> activityID

foreign_id_activities:\<foreignObjectID>:\<activityID> -> activityID

Добавить одно и то же событие(с одним внутренним ID) в несколько лент можно с помощью параметра redirect_to в запросе AddActivity, событие добавляется в перечисленные ленты.

## Collections

Также мы поддерживаем коллекции объектов, на которые ссылаются события.
Сами объекты могут быть чем угодно.

В БД храним это все так:

collections:\<collectionID> -> collectionData

objects:\<collectionID>:\<objectID> -> objectData

При удалении объекта есть два варианта поведения:
1) удаляются все события, ссылающиеся на этот объект - CASCADE
2) во всех событиях ссылающихся на этот объект foreign_object_id заменяется на "", то есть они перестают ссылаться на этот объект - SET_NIL

То, какое поведение применяется указывается в коллекции и одинаково для всех объектов в этой коллекции.


## Краткий принцип работы методов относительно БД
* AddActivity: записываем данные в activities:\<activityID>, записываем данные в feed_activities:\<feedID>:\<activityID> и если есть objectID, то записываем данные в foreign_id_activities:\<objectID>:\<activityID>
* GetActivity: получаем данные в activities:\<activityID>
* GetActivityByObjectID: получаем данные в диапазоне foreign_id_activities:\<objectID>, если там больше 1 записи, то возвращаем ошибку
* UpdateActivity: перезаписываем данные в activities:\<activityID>
* DeleteActivity: удаляем данные в  activities:\<activityID>, удаляем данные в feed_activities:\<feedID>:\<activityID> и если есть objectID, то удаляем данные в foreign_id_activities:\<objectID>:\<activityID>
* CreateFeed: записывает данные в feeds:\<feedID>
* GetFeed: получает данные в feeds:\<feedID>
* GetFeedActivities: получает id событий по диапазону feed_activities:\<feedID> и для каждого получаем данные в activities:\<activityID>
* UpdateFeed: перезаписывает данные в feeds:\<feedID>
* DeleteFeed: удаляет данные в feeds:\<feedID> и удаляет диапазон feed_activities:\<feedID>
* CreateCollection: записывает данные в collections:\<collectionID>
* CreateObject: записывает данные в objects:\<collectionID>:\<objectID>
* UpdateObject: перезаписывает данные в objects:\<collectionID>:\<objectID>
* DeleteObject: удаляет данные из objects:\<collectionID>:\<objectID> и в зависимости от типа удаления в коллекции изменяет/удаляет события, которые находим по диапазону в индексе foreign_id_activities:\<objectID>
