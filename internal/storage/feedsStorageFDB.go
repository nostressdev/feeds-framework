package storage

import (
	"fmt"
	go_time "time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type ConfigFeedsFDB struct {
	DB       fdb.Database
	Subspace subspace.Subspace
}

type FeedsStorageFDB struct {
	*ConfigFeedsFDB
	FeedsSubspace               subspace.Subspace
	ActivitiesSubspace          subspace.Subspace
	FeedActivitiesSubspace      subspace.Subspace
	ActivityFeedsSubspace       subspace.Subspace
	ForeignIDActivitiesSubspace subspace.Subspace
	CollectionsSubspace         subspace.Subspace
	CollectionObjectsSubspace   subspace.Subspace
}

func NewFeedsFDB(config *ConfigFeedsFDB) FeedsStorage {
	return &FeedsStorageFDB{
		ConfigFeedsFDB:              config,
		FeedsSubspace:               config.Subspace.Sub("feeds"),
		ActivitiesSubspace:          config.Subspace.Sub("activities"),
		FeedActivitiesSubspace:      config.Subspace.Sub("feed_activities"),
		ActivityFeedsSubspace:       config.Subspace.Sub("activity_feeds"),
		ForeignIDActivitiesSubspace: config.Subspace.Sub("foreign_id_activities"),
		CollectionsSubspace:         config.Subspace.Sub("collections"),
		CollectionObjectsSubspace:   config.Subspace.Sub("collection_objects"),
	}
}

func (s *FeedsStorageFDB) AddActivity(feedID, objectID, userID, activityType string, time int64, redirectTo []string, extraData *anypb.Any) (*proto.Activity, error) {
	if time == 0 {
		time = go_time.Now().UnixNano()
	}
	id, err := UUIDFromTimestamp(uint64(time))
	if err != nil {
		return nil, err
	}
	activity := &proto.Activity{
		StringId:        id,
		ForeignObjectId: objectID,
		CreatedAt:       time,
		UserId:          userID,
		ActivityType:    activityType,
		ExtraData:       extraData,
	}
	_, err = s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := tr.Get(s.FeedsSubspace.Sub(feedID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such feed")
		}
		bytes, err = protobuf.Marshal(activity)
		if err != nil {
			return nil, err
		}
		tr.Set(s.ActivitiesSubspace.Sub(activity.StringId), bytes)
		tr.Set(s.FeedActivitiesSubspace.Sub(feedID, activity.StringId), []byte(activity.StringId))
		tr.Set(s.ActivityFeedsSubspace.Sub(activity.StringId, feedID), []byte(feedID))
		if activity.ForeignObjectId != "" {
			tr.Set(s.ForeignIDActivitiesSubspace.Sub(objectID, activity.StringId), []byte(activity.StringId))
		}
		for _, redirectFeedID := range redirectTo {
			tr.Set(s.FeedActivitiesSubspace.Sub(redirectFeedID, activity.StringId), []byte(activity.StringId))
			tr.Set(s.ActivityFeedsSubspace.Sub(activity.StringId, redirectFeedID), []byte(redirectFeedID))
		}
		return nil, nil
	})
	return activity, err
}

func (s *FeedsStorageFDB) GetActivity(activityID string) (*proto.Activity, error) {
	activity := &proto.Activity{}
	_, err := s.DB.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		bytes, err := tr.Get(s.ActivitiesSubspace.Sub(activityID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such activity")
		}
		err = protobuf.Unmarshal(bytes, activity)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	return activity, err
}

func (s *FeedsStorageFDB) GetActivityByObjectID(objectID string) (*proto.Activity, error) {
	var activity *proto.Activity = &proto.Activity{}
	_, err := s.DB.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		begin, end := s.ForeignIDActivitiesSubspace.Sub(objectID).FDBRangeKeys()
		iter := tr.GetRange(fdb.KeyRange{Begin: begin, End: end}, fdb.RangeOptions{}).Iterator()
		for iter.Advance() {
			if activity.StringId != "" {
				return nil, nerrors.BadRequest.New("object has more than 1 connected activity")
			}
			kv, err := iter.Get()
			if err != nil {
				return nil, err
			}
			activityID := string(kv.Value)
			bytes, err := tr.Get(s.ActivitiesSubspace.Sub(activityID)).Get()
			if err != nil {
				return nil, err
			}
			if bytes == nil {
				return nil, nerrors.BadRequest.New("no such activity")
			}
			err = protobuf.Unmarshal(bytes, activity)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}
		if activity.StringId == "" {
			return nil, nerrors.BadRequest.New("no such activity")
		}
		return nil, nil
	})
	return activity, err
}

func (s *FeedsStorageFDB) UpdateActivity(activityID string, extraData *anypb.Any) (*proto.Activity, error) {
	activity := &proto.Activity{}
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := tr.Get(s.ActivitiesSubspace.Sub(activityID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such activity")
		}
		err = protobuf.Unmarshal(bytes, activity)
		if err != nil {
			return nil, err
		}
		activity.ExtraData = extraData

		bytes, err = protobuf.Marshal(activity)
		if err != nil {
			return nil, err
		}
		tr.Set(s.FeedsSubspace.Sub(activity.Id), bytes)
		return nil, nil
	})
	return activity, err
}

func (s *FeedsStorageFDB) DeleteActivity(activityID string) error {
	activity := &proto.Activity{}
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		fmt.Println(1)
		bytes, err := tr.Get(s.ActivitiesSubspace.Sub(activityID)).Get()
		if err != nil {
			return nil, err
		}
		fmt.Println(2)
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such activity")
		}
		fmt.Println(3)
		err = protobuf.Unmarshal(bytes, activity)
		if err != nil {
			return nil, err
		}
		fmt.Println(4)
		if activity.ForeignObjectId != "" {
			tr.Clear(s.ForeignIDActivitiesSubspace.Sub(activity.ForeignObjectId, activityID))
		}
		fmt.Println(5)
		tr.Clear(s.ActivitiesSubspace.Sub(activityID))
		fmt.Println(6)
		begin, end := s.ActivityFeedsSubspace.Sub(activityID).FDBRangeKeys()
		fmt.Println(6)
		iter := tr.GetRange(fdb.KeyRange{Begin: begin, End: end}, fdb.RangeOptions{}).Iterator()
		for iter.Advance() {
			fmt.Println(7)
			kv, err := iter.Get()
			if err != nil {
				return nil, err
			}
			feedID := string(kv.Value)
			tr.Clear(s.FeedActivitiesSubspace.Sub(feedID, activityID))
		}
		fmt.Println(8)
		tr.Clear(s.ActivityFeedsSubspace.Sub(activityID))
		return nil, nil
	})
	return err
}

func (s *FeedsStorageFDB) CreateFeed(userID string, extraData *anypb.Any) (*proto.Feed, error) {
	id, err := UUIDFromTimestamp(GetUnixTimestampNow())
	if err != nil {
		return nil, err
	}
	feed := &proto.Feed{
		Id:        id,
		UserId:    userID,
		ExtraData: extraData,
	}
	_, err = s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := protobuf.Marshal(feed)
		if err != nil {
			return nil, err
		}
		tr.Set(s.FeedsSubspace.Sub(feed.Id), bytes)
		return nil, nil
	})
	return feed, err
}

func (s *FeedsStorageFDB) GetFeed(feedID string) (*proto.Feed, error) {
	feed := &proto.Feed{}
	_, err := s.DB.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		bytes, err := tr.Get(s.FeedsSubspace.Sub(feedID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such feed")
		}
		err = protobuf.Unmarshal(bytes, feed)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	return feed, err
}

func (s *FeedsStorageFDB) GetFeedActivities(feedID string, limit int64, offsetID string) ([]*proto.Activity, error) {
	activities := make([]*proto.Activity, 0)

	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := tr.Get(s.FeedsSubspace.Sub(feedID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such feed")
		}
		begin, end := s.FeedActivitiesSubspace.Sub(feedID).FDBRangeKeySelectors()
		if offsetID != "" {
			end = fdb.LastLessThan(s.FeedActivitiesSubspace.Sub(feedID, offsetID))
		}
		iter := tr.GetRange(fdb.SelectorRange{Begin: begin, End: end}, fdb.RangeOptions{Limit: int(limit), Reverse: true}).Iterator()
		for iter.Advance() {
			kv, err := iter.Get()
			if err != nil {
				return nil, err
			}
			activityID := string(kv.Value)
			bytes, err := tr.Get(s.ActivitiesSubspace.Sub(activityID)).Get()
			if err != nil {
				return nil, err
			}
			activity := &proto.Activity{}
			err = protobuf.Unmarshal(bytes, activity)
			if err != nil {
				return nil, err
			}
			activities = append(activities, activity)
		}
		return nil, nil
	})
	return activities, err
}

func (s *FeedsStorageFDB) UpdateFeed(feedID string, extraData *anypb.Any) (*proto.Feed, error) {
	feed := &proto.Feed{}
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := tr.Get(s.FeedsSubspace.Sub(feedID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such feed")
		}
		err = protobuf.Unmarshal(bytes, feed)
		if err != nil {
			return nil, err
		}
		feed.ExtraData = extraData

		bytes, err = protobuf.Marshal(feed)
		if err != nil {
			return nil, err
		}
		tr.Set(s.FeedsSubspace.Sub(feed.Id), bytes)
		return nil, nil
	})
	return feed, err
}

func (s *FeedsStorageFDB) DeleteFeed(feedID string) error {
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := tr.Get(s.FeedsSubspace.Sub(feedID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such feed")
		}
		tr.Clear(s.FeedsSubspace.Sub(feedID))
		begin, end := s.FeedActivitiesSubspace.Sub(feedID).FDBRangeKeys()
		iter := tr.GetRange(fdb.KeyRange{Begin: begin, End: end}, fdb.RangeOptions{}).Iterator()
		for iter.Advance() {
			kv, err := iter.Get()
			if err != nil {
				return nil, err
			}
			activityID := string(kv.Value)
			tr.Clear(s.ActivityFeedsSubspace.Sub(activityID, feedID))
		}
		tr.ClearRange(s.FeedActivitiesSubspace.Sub(feedID))
		return nil, nil
	})
	return err
}

func (s *FeedsStorageFDB) CreateCollection(collectionID string, deletingType proto.DeletingType) (*proto.Collection, error) {
	var err error
	if collectionID == "" {
		collectionID, err = UUIDFromTimestamp(GetUnixTimestampNow())
		if err != nil {
			return nil, err
		}
	}
	collection := &proto.Collection{
		Id:           collectionID,
		DeletingType: deletingType,
	}
	_, err = s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := protobuf.Marshal(collection)
		if err != nil {
			return nil, err
		}
		tr.Set(s.CollectionsSubspace.Sub(collection.Id), bytes)
		return nil, nil
	})
	return collection, err
}

func (s *FeedsStorageFDB) CreateObject(collectionID, objectID string, data *anypb.Any) (*proto.Object, error) {
	var err error
	if objectID == "" {
		objectID, err = UUIDFromTimestamp(GetUnixTimestampNow())
		if err != nil {
			return nil, err
		}
	}
	object := &proto.Object{
		Id:   objectID,
		Data: data,
	}
	_, err = s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := tr.Get(s.CollectionsSubspace.Sub(collectionID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such collection")
		}
		bytes, err = protobuf.Marshal(object)
		if err != nil {
			return nil, err
		}
		tr.Set(s.CollectionObjectsSubspace.Sub(collectionID, objectID), bytes)
		return nil, nil
	})
	return object, err
}

func (s *FeedsStorageFDB) GetObject(collectionID, objectID string) (*proto.Object, error) {
	object := &proto.Object{}
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := tr.Get(s.CollectionsSubspace.Sub(collectionID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such collection")
		}
		bytes, err = tr.Get(s.CollectionObjectsSubspace.Sub(collectionID, objectID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such object")
		}
		err = protobuf.Unmarshal(bytes, object)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	return object, err
}


func (s *FeedsStorageFDB) UpdateObject(collectionID, objectID string, data *anypb.Any) (*proto.Object, error) {
	object := &proto.Object{
		Id:   objectID,
		Data: data,
	}
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := tr.Get(s.CollectionsSubspace.Sub(collectionID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such collection")
		}
		bytes, err = tr.Get(s.CollectionObjectsSubspace.Sub(collectionID, objectID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such object")
		}
		bytes, err = protobuf.Marshal(object)
		if err != nil {
			return nil, err
		}
		tr.Set(s.CollectionObjectsSubspace.Sub(collectionID, objectID), bytes)
		return nil, nil
	})
	return object, err
}

func (s *FeedsStorageFDB) DeleteObject(collectionID, objectID string) error {
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := tr.Get(s.CollectionsSubspace.Sub(collectionID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such collection")
		}
		bytes, err = tr.Get(s.CollectionObjectsSubspace.Sub(collectionID, objectID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such object")
		}
		tr.Clear(s.CollectionObjectsSubspace.Sub(collectionID, objectID))
		return nil, nil
	})
	return err
}
