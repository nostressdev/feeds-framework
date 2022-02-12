# feeds-framework

## Feeds and activities

Главные используемые сущности - Activity и Feed(набор Activity).

В БД храним так:

feeds:\<feedID> -> feedData

activities:\<activityID> -> activityData

Далее для удобного получения событий будем хранить индексы:

feed_activities:\<feedID>:\<activityID> -> activityID

activity_feeds:\<activityID>:\<feedID> -> feedID

object_id_activities:\<foreignObjectID>:\<activityID> -> activityID

Добавить одно и то же событие(с одним внутренним ID) в несколько лент можно с помощью параметра redirect_to в запросе AddActivity, событие добавляется в перечисленные ленты.

## Collections

Также мы поддерживаем коллекции объектов, на которые ссылаются события.
Сами объекты могут быть чем угодно.

В БД храним это все так:

collections:\<collectionID> -> collectionData

collection_objects:\<collectionID>:\<objectID> -> objectData

При удалении объекта есть два варианта поведения:
1) удаляются все события, ссылающиеся на этот объект - CASCADE
2) во всех событиях ссылающихся на этот объект object_id заменяется на "", то есть они перестают ссылаться на этот объект - SET_NIL

То, какое поведение применяется указывается в коллекции и одинаково для всех объектов в этой коллекции.

## Reactions
Реакции это Activity с непустым linked_activity_id. 

Важное замечание: у Activity есть object_id и linked_activity_id, при работе с реакциями используются оба, причем в таком контексте: object_id указывает на верхнеуровневый объект, к котому привязана реакция, а linked_activity_id на непосредственного предка.

Пример: пост, комментарий к нему и ответ на комментарий. 

Пост - Activity c пустым linked_activity_id

Верхнеуровневый комментарий - Activity с linked_activity_id = post_id = object_id

Ответ на комментарий - Activity с linked_activity_id = prev_comment и object_id = post_id

Т.к. это Activity, то в БД они хранятся там же, в activities:\<activityID>, и так же будет храниться индекс object_id_activities:\<foreignObjectID>:\<activityID> -> activityID, для получения всех реакций на объект(например все комментарии к посту)

Реакции являются полноценными Acitiviy, поэтому все операции связанные с лентами будут работать так же, как и с обычными Activity

Так же будем хранить индекс со всеми непостредственными реакциями на activity:
activity_reactions:<activityID>:<reactionActivityID> -> reactionActivityID

## Grouping feeds
Группирующие ленты - ленты в которых события объединяются исходя из заданного формата ключа.

Формат ключа: шаблон как в https://pkg.go.dev/text/template, в который будет передан Activity, соответственно обращаться можно только к полям этого самого Activity.

Храним в БД так:

grouping_feeds:\<feedID> -> feedData

А так же индексы:

grouping_feed_activities:\<feedID>:\<createdFormatKey>:\<activityID> -> activityID

activity_grouping_feeds:\<activityID>:\<feedID> -> feedID

## Краткий принцип работы методов относительно БД
### Activities
* AddActivity(to simple feed): записываем данные в activities:\<activityID>, записываем данные в feed_activities:\<feedID>:\<activityID>, записываем данные в activity_feeds:\<activityID>:\<feedID> и если есть objectID, то записываем данные в object_id_activities:\<objectID>:\<activityID>
* AddActivity(to grouping feed): записываем данные в activities:\<activityID>, записываем данные в grouping_feed_activities:\<feedID>:\<createedFormatKey>:\<activityID>, записываем данные в activity_grouping_feeds:\<activityID>:\<feedID> и если есть objectID, то записываем данные в object_id_activities:\<objectID>:\<activityID>
* GetActivity: получаем данные в activities:\<activityID>
* GetActivityByObjectID: получаем данные в диапазоне object_id_activities:\<objectID>, если там больше 1 записи, то возвращаем ошибку
* UpdateActivity: перезаписываем данные в activities:\<activityID>
* DeleteActivity: удаляем данные в  activities:\<activityID>, если есть objectID, то удаляем данные в object_id_activities:\<objectID>:\<activityID>, удаляем feed_activities:\<feedID>:\<activityID> для всех feedID из диапазона  activity_feeds:\<activityID>, а затем очищаем и сам диапазон activity_feeds:\<activityID>
### Feeds
* CreateFeed: записывает данные в feeds:\<feedID>
* GetFeed: получает данные в feeds:\<feedID>
* GetFeedActivities: получает id событий по диапазону feed_activities:\<feedID> и для каждого получаем данные в activities:\<activityID>
* UpdateFeed: перезаписывает данные в feeds:\<feedID>
* DeleteFeed: удаляет данные в feeds:\<feedID>, удаляет данные в activity_feeds:\<activity_id>:\<feed_id> для всех activity_id из feed_activities:\<feedID> и наконец удаляет диапазон feed_activities:\<feedID>
### Collections
* CreateCollection: записывает данные в collections:\<collectionID>
* CreateObject: записывает данные в objects:\<collectionID>:\<objectID>
* GetObject: получает данные из objects:\<collectionID>:\<objectID>
* UpdateObject: перезаписывает данные в objects:\<collectionID>:\<objectID>
* DeleteObject: удаляет данные из objects:\<collectionID>:\<objectID> и в зависимости от типа удаления в коллекции изменяет/удаляет события, которые находим по диапазону в индексе object_id_activities:\<objectID>
### Reactions
* CreateReaction: записываем данные в activities:\<activityID>, записываем данные в activity_reactions:\<linkedActivityID>:\<reactionID> и если есть objectID, то записываем данные в object_id_activities:\<objectID>:\<activityID>
* AddReaction(to simple feed): записываем данные в feed_activities:\<feedID>:\<activityID>, записываем данные в activity_feeds:\<activityID>:\<feedID>
* AddReaction(to grouping feed): записываем данные в grouping_feed_activities:\<feedID>:\<createedFormatKey>:\<activityID>, записываем данные в activity_grouping_feeds:\<activityID>:\<feedID>
* GetReation: получаем данные из activities:\<activityID>
* GetActivityReactions: получаем все данные из диапазона activity_reactions:\<linkedActivityID>, и возможно рекурсивно далее
* UpdateReaction: перезаписываем данные в activities:\<activityID>
* DeleteReaction: удаляем данные в  activities:\<activityID>, если есть objectID, то удаляем данные в object_id_activities:\<objectID>:\<activityID>, удаляем feed_activities:\<feedID>:\<activityID> для всех feedID из диапазона  activity_feeds:\<activityID>, а затем очищаем и сам диапазон activity_feeds:\<activityID>, удаляем activity_reactions:\<linkedActivityID>:\<activityID>, для всех activity из activity_reactions:\<activityID> проставляем linkedActivityID="deleted" и после очищаем диапазон activity_reactions:\<activityID>
### GroupingFeeds
* CreateGroupingFeed: записывает данные в данные в grouping_feeds:\<feedID>
* GetGroupingFeed: получает данные в grouping_feeds:\<feedID>
* GetGroupingFeedActivities: получает данные из диапазона grouping_feed_activities:\<feedID>
* UpdateGroupingFeed: обновляет данные в grouping_feeds:\<feedID>
* DeleteGroupingFeed: удаляет данные в данные в grouping_feeds:\<feedID>, удаляет данные в activity_grouping_feeds:\<activityID>:\<feedID> для всех activityID из grouping_feed_activities:\<feedID> и наконец удаляет диапазон grouping_feed_activities:\<feedID>
