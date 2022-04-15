package storage

import (
	"html/template"
	"log"
	"strings"
	go_time "time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/bwmarrin/snowflake"
	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type ConfigFeedsFDB struct {
	DB        fdb.Database
	Subspace  subspace.Subspace
	Generator *snowflake.Node
}

type FeedsStorageFDB struct {
	*ConfigFeedsFDB
	FeedsSubspace                  subspace.Subspace
	ActivitiesSubspace             subspace.Subspace
	FeedActivitiesSubspace         subspace.Subspace
	ActivityFeedsSubspace          subspace.Subspace
	ObjectIDActivitiesSubspace     subspace.Subspace
	CollectionsSubspace            subspace.Subspace
	CollectionObjectsSubspace      subspace.Subspace
	GroupingFeedsSubspace          subspace.Subspace
	GroupingFeedActivitiesSubspace subspace.Subspace
	ActivityGroupingFeedsSubspace  subspace.Subspace
	ActivityReactionsSubspace      subspace.Subspace
}

func NewFeedsFDB(config *ConfigFeedsFDB) FeedsStorage {
	return &FeedsStorageFDB{
		ConfigFeedsFDB:                 config,
		FeedsSubspace:                  config.Subspace.Sub("feeds"),
		ActivitiesSubspace:             config.Subspace.Sub("activities"),
		FeedActivitiesSubspace:         config.Subspace.Sub("feed_activities"),
		ActivityFeedsSubspace:          config.Subspace.Sub("activity_feeds"),
		ObjectIDActivitiesSubspace:     config.Subspace.Sub("object_id_activities"),
		CollectionsSubspace:            config.Subspace.Sub("collections"),
		CollectionObjectsSubspace:      config.Subspace.Sub("collection_objects"),
		GroupingFeedsSubspace:          config.Subspace.Sub("grouping_feeds"),
		GroupingFeedActivitiesSubspace: config.Subspace.Sub("grouping_feed_activities"),
		ActivityGroupingFeedsSubspace:  config.Subspace.Sub("activity_grouping_feeds"),
		ActivityReactionsSubspace:      config.Subspace.Sub("activity_reactions"),
	}
}

func (s *FeedsStorageFDB) AddActivity(feedID, objectID, userID, activityType string, time int64, redirectTo []string, extraData *anypb.Any) (*proto.Activity, error) {
	if time == 0 {
		time = go_time.Now().UnixNano()
	}
	snowflakeID := s.Generator.Generate()
	activity := &proto.Activity{
		Id:           snowflakeID.Int64(),
		StringId:     snowflakeID.String(),
		ObjectId:     objectID,
		CreatedAt:    time,
		UserId:       userID,
		ActivityType: activityType,
		ExtraData:    extraData,
	}
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := protobuf.Marshal(activity)
		if err != nil {
			return nil, err
		}
		err = s.addToFeed(tr, feedID, activity)
		if err != nil {
			return nil, err
		}
		tr.Set(s.ActivitiesSubspace.Sub(activity.StringId), bytes)
		if activity.ObjectId != "" {
			tr.Set(s.ObjectIDActivitiesSubspace.Sub(objectID, activity.StringId), []byte(activity.StringId))
		}
		for _, redirectFeedID := range redirectTo {
			err = s.addToFeed(tr, redirectFeedID, activity)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	})
	return activity, err
}

func (s *FeedsStorageFDB) AddExistingActivity(feedID, activityID string) (*proto.Activity, error) {
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
		return nil, s.addToFeed(tr, feedID, activity)
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
	activity := &proto.Activity{}
	_, err := s.DB.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		begin, end := s.ObjectIDActivitiesSubspace.Sub(objectID).FDBRangeKeys()
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
			log.Println(activityID)
			bytes, err := tr.Get(s.ActivitiesSubspace.Sub(activityID)).Get()
			if err != nil {
				return nil, err
			}
			if bytes == nil {
				return nil, nerrors.BadRequest.New("no such activity " + string(activityID))
			}
			err = protobuf.Unmarshal(bytes, activity)
			if err != nil {
				return nil, err
			}
		}
		if activity.StringId == "" {
			return nil, nerrors.BadRequest.New("no such activity " + string(objectID))
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
		tr.Set(s.ActivitiesSubspace.Sub(activity.StringId), bytes)
		return nil, nil
	})
	return activity, err
}

func (s *FeedsStorageFDB) DeleteActivity(activityID string) error {
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
		if activity.ObjectId != "" {
			tr.Clear(s.ObjectIDActivitiesSubspace.Sub(activity.ObjectId, activityID))
		}
		if err := s.deleteActivity(tr, activity); err != nil {
			return nil, err
		}
		begin, end := s.ObjectIDActivitiesSubspace.Sub(activityID).FDBRangeKeys()
		iter := tr.GetRange(fdb.KeyRange{Begin: begin, End: end}, fdb.RangeOptions{}).Iterator()
		for iter.Advance() {
			kv, err := iter.Get()
			if err != nil {
				return nil, err
			}
			bytes, err := tr.Get(s.ActivitiesSubspace.Sub(string(kv.Value))).Get()
			if err != nil {
				return nil, err
			}
			if bytes == nil {
				return nil, nerrors.Internal.New("no such reaction")
			}
			reaction := &proto.Activity{}
			if err := protobuf.Unmarshal(bytes, reaction); err != nil {
				return nil, err
			}
			reaction.ObjectId = "deleted"
			if reaction.LinkedActivityId == activityID {
				reaction.LinkedActivityId = "deleted"
			}
			bytes, err = protobuf.Marshal(reaction)
			if err != nil {
				return nil, err
			}
			tr.Set(s.ActivitiesSubspace.Sub(reaction.StringId), bytes)
		}
		tr.ClearRange(s.ActivityReactionsSubspace.Sub(activityID))
		tr.ClearRange(s.ObjectIDActivitiesSubspace.Sub(activityID))
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
			end = fdb.FirstGreaterOrEqual(s.FeedActivitiesSubspace.Sub(feedID, offsetID))
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
			log.Println(activityID)
			activities = append(activities, activity)
		}
		return nil, nil
	})
	log.Println(activities)
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

func (s *FeedsStorageFDB) CreateReaction(activityType, userID, linkedActivityID string, time int64, extraData *anypb.Any) (*proto.Activity, error) {
	if time == 0 {
		time = go_time.Now().UnixNano()
	}
	snowflakeID := s.Generator.Generate()
	reaction := &proto.Activity{
		Id:               snowflakeID.Int64(),
		StringId:         snowflakeID.String(),
		LinkedActivityId: linkedActivityID,
		CreatedAt:        time,
		UserId:           userID,
		ActivityType:     activityType,
		ExtraData:        extraData,
	}
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := tr.Get(s.ActivitiesSubspace.Sub(linkedActivityID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such linked activity")
		}
		activity := &proto.Activity{}
		err = protobuf.Unmarshal(bytes, activity)
		if err != nil {
			return nil, err
		}
		reaction.ObjectId = activity.LinkedActivityId
		if reaction.ObjectId == "" {
			reaction.ObjectId = reaction.LinkedActivityId
		}
		bytes, err = protobuf.Marshal(reaction)
		if err != nil {
			return nil, err
		}
		tr.Set(s.ActivitiesSubspace.Sub(reaction.StringId), bytes)
		tr.Set(s.ActivityReactionsSubspace.Sub(reaction.LinkedActivityId, reaction.StringId), []byte(reaction.StringId))
		tr.Set(s.ObjectIDActivitiesSubspace.Sub(reaction.ObjectId, reaction.StringId), []byte(reaction.StringId))
		return nil, nil
	})
	return reaction, err
}

func (s *FeedsStorageFDB) deleteActivity(tr fdb.Transaction, activity *proto.Activity) error {
	tr.Clear(s.ActivitiesSubspace.Sub(activity.StringId))
	// for simple feeds
	begin, end := s.ActivityFeedsSubspace.Sub(activity.StringId).FDBRangeKeys()
	iter := tr.GetRange(fdb.KeyRange{Begin: begin, End: end}, fdb.RangeOptions{}).Iterator()
	for iter.Advance() {
		kv, err := iter.Get()
		if err != nil {
			return err
		}
		log.Println("deleting in " + string(kv.Value))
		tr.Clear(s.FeedActivitiesSubspace.Sub(string(kv.Value), activity.StringId))
	}
	tr.Clear(s.ActivityFeedsSubspace.Sub(activity.StringId))

	// for grouping feeds
	begin, end = s.ActivityGroupingFeedsSubspace.Sub(activity.StringId).FDBRangeKeys()
	iter = tr.GetRange(fdb.KeyRange{Begin: begin, End: end}, fdb.RangeOptions{}).Iterator()
	for iter.Advance() {
		kv, err := iter.Get()
		if err != nil {
			return err
		}
		bytes, err := tr.Get(s.GroupingFeedsSubspace.Sub(string(kv.Value))).Get()
		if err != nil {
			return err
		}
		if bytes == nil {
			return nerrors.Internal.New("no such grouping feed")
		}
		feed := &proto.GroupingFeed{}
		err = protobuf.Unmarshal(bytes, feed)
		if err != nil {
			return err
		}
		templ, err := template.New("key").Parse(feed.KeyFormat)
		if err != nil {
			return err
		}
		builder := new(strings.Builder)
		err = templ.Execute(builder, activity)
		if err != nil {
			return err
		}
		key := builder.String()
		tr.Clear(s.GroupingFeedActivitiesSubspace.Sub(string(kv.Value), key, activity.StringId))
	}
	tr.Clear(s.ActivityGroupingFeedsSubspace.Sub(activity.StringId))
	return nil
}

func (s *FeedsStorageFDB) addToFeed(tr fdb.Transaction, feedID string, activity *proto.Activity) error {
	bytes, err := tr.Get(s.FeedsSubspace.Sub(feedID)).Get()
	if err != nil {
		return err
	}
	isSimpleFeed := !(bytes == nil)
	groupingFeedBytes, err := tr.Get(s.GroupingFeedsSubspace.Sub(feedID)).Get()
	if err != nil {
		return err
	}
	isGroupingFeed := !(groupingFeedBytes == nil)
	if !isSimpleFeed && !isGroupingFeed {
		return nerrors.BadRequest.New("no such feed")
	}

	if isSimpleFeed {
		tr.Set(s.FeedActivitiesSubspace.Sub(feedID, activity.StringId), []byte(activity.StringId))
		tr.Set(s.ActivityFeedsSubspace.Sub(activity.StringId, feedID), []byte(feedID))
	} else {
		groupingFeed := &proto.GroupingFeed{}
		err = protobuf.Unmarshal(groupingFeedBytes, groupingFeed)
		if err != nil {
			return err
		}

		templ, err := template.New("key").Parse(groupingFeed.KeyFormat)
		if err != nil {
			return err
		}
		builder := new(strings.Builder)
		err = templ.Execute(builder, activity)
		if err != nil {
			return err
		}
		key := builder.String()
		tr.Set(s.GroupingFeedActivitiesSubspace.Sub(feedID, key, activity.StringId), []byte(activity.StringId))
		tr.Set(s.ActivityGroupingFeedsSubspace.Sub(activity.StringId, feedID), []byte(feedID))
	}
	return nil

}

func (s *FeedsStorageFDB) AddReaction(feedID, reactionID string) error {
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		reactionBytes, err := tr.Get(s.ActivitiesSubspace.Sub(reactionID)).Get()
		if err != nil {
			return nil, err
		}
		if reactionBytes == nil {
			return nil, nerrors.BadRequest.New("no such reaction")
		}
		reaction := &proto.Activity{}
		err = protobuf.Unmarshal(reactionBytes, reaction)
		if err != nil {
			return nil, err
		}
		return nil, s.addToFeed(tr, feedID, reaction)
	})
	return err
}

func (s *FeedsStorageFDB) GetReaction(reactionID string) (*proto.Activity, error) {
	reaction := &proto.Activity{}
	_, err := s.DB.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		bytes, err := tr.Get(s.ActivitiesSubspace.Sub(reactionID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such reaction")
		}
		err = protobuf.Unmarshal(bytes, reaction)
		if err != nil {
			return nil, err
		}
		if reaction.LinkedActivityId == "" {
			return nil, nerrors.BadRequest.New("activity is not a reaction")
		}
		return nil, nil
	})
	return reaction, err
}

func (s *FeedsStorageFDB) GetActivityReactions(activityID string, limit int64, offsetID string) ([]*proto.Activity, error) {
	reactions := make([]*proto.Activity, 0)
	_, err := s.DB.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		begin, end := s.ActivityReactionsSubspace.Sub(activityID).FDBRangeKeySelectors()
		if offsetID != "" {
			end = fdb.FirstGreaterOrEqual(s.ActivityReactionsSubspace.Sub(activityID, offsetID))
		}
		iter := tr.GetRange(fdb.SelectorRange{Begin: begin, End: end}, fdb.RangeOptions{Limit: int(limit), Reverse: true}).Iterator()
		for iter.Advance() {
			kv, err := iter.Get()
			if err != nil {
				return nil, err
			}
			reactionID := string(kv.Value)
			bytes, err := tr.Get(s.ActivitiesSubspace.Sub(reactionID)).Get()
			if err != nil {
				return nil, err
			}
			if bytes == nil {
				return nil, nerrors.Internal.New("reaction doesn't exist")
			}

			reaction := &proto.Activity{}
			err = protobuf.Unmarshal(bytes, reaction)
			if err != nil {
				return nil, err
			}
			reactions = append(reactions, reaction)
		}
		return nil, nil
	})
	return reactions, err
}

func (s *FeedsStorageFDB) UpdateReaction(reactionID string, extraData *anypb.Any) (*proto.Activity, error) {
	reaction := &proto.Activity{}
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := tr.Get(s.ActivitiesSubspace.Sub(reactionID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such activity")
		}
		err = protobuf.Unmarshal(bytes, reaction)
		if err != nil {
			return nil, err
		}
		if reaction.LinkedActivityId == "" {
			return nil, nerrors.BadRequest.New("activity is not a reaction")
		}
		reaction.ExtraData = extraData

		bytes, err = protobuf.Marshal(reaction)
		if err != nil {
			return nil, err
		}
		tr.Set(s.ActivitiesSubspace.Sub(reaction.StringId), bytes)
		return nil, nil
	})
	return reaction, err
}

func (s *FeedsStorageFDB) DeleteReaction(reactionID string) error {
	reaction := &proto.Activity{}
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := tr.Get(s.ActivitiesSubspace.Sub(reactionID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("1) no such activity " + reactionID)
		}
		if err := protobuf.Unmarshal(bytes, reaction); err != nil {
			return nil, err
		}
		if reaction.LinkedActivityId == "" {
			return nil, nerrors.BadRequest.New("2) activity is not a reaction " + reaction.String())
		}
		if err := s.deleteActivity(tr, reaction); err != nil {
			return nil, err
		}
		// reactions clearing
		tr.Clear(s.ActivityReactionsSubspace.Sub(reaction.LinkedActivityId, reactionID))
		tr.Clear(s.ObjectIDActivitiesSubspace.Sub(reaction.ObjectId, reactionID))
		begin, end := s.ActivityReactionsSubspace.Sub(reactionID).FDBRangeKeys()
		iter := tr.GetRange(fdb.KeyRange{Begin: begin, End: end}, fdb.RangeOptions{}).Iterator()
		for iter.Advance() {
			kv, err := iter.Get()
			if err != nil {
				return nil, err
			}
			bytes, err := tr.Get(s.ActivitiesSubspace.Sub(string(kv.Value))).Get()
			if err != nil {
				return nil, err
			}
			if bytes == nil {
				return nil, nerrors.Internal.New("3) no such activity " + string(kv.Value))
			}
			activity := &proto.Activity{}
			err = protobuf.Unmarshal(bytes, activity)
			if err != nil {
				return nil, err
			}
			activity.LinkedActivityId = "deleted"
			bytes, err = protobuf.Marshal(activity)
			if err != nil {
				return nil, err
			}
			tr.Set(s.ActivitiesSubspace.Sub(activity.StringId), bytes)
		}
		tr.ClearRange(s.ActivityReactionsSubspace.Sub(reactionID))

		return nil, nil
	})
	return err
}

func (s *FeedsStorageFDB) CreateGroupingFeed(userID, keyFormat string, extraData *anypb.Any) (*proto.GroupingFeed, error) {
	id, err := UUIDFromTimestamp(GetUnixTimestampNow())
	if err != nil {
		return nil, err
	}
	feed := &proto.GroupingFeed{
		Id:        id,
		UserId:    userID,
		KeyFormat: keyFormat,
		ExtraData: extraData,
	}
	_, err = s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := protobuf.Marshal(feed)
		if err != nil {
			return nil, err
		}
		tr.Set(s.GroupingFeedsSubspace.Sub(feed.Id), bytes)
		return nil, nil
	})
	return feed, err
}

func (s *FeedsStorageFDB) GetGroupingFeed(groupingFeedID string) (*proto.GroupingFeed, error) {
	feed := &proto.GroupingFeed{}
	_, err := s.DB.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		bytes, err := tr.Get(s.GroupingFeedsSubspace.Sub(groupingFeedID)).Get()
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

func (s *FeedsStorageFDB) GetGroupingFeedActivities(groupingFeedID string, limit int64, offsetID string) ([]*proto.ActivityGroup, error) {
	activityGroups := make(map[string][]*proto.Activity)
	_, err := s.DB.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		bytes, err := tr.Get(s.GroupingFeedsSubspace.Sub(groupingFeedID)).Get()
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, nerrors.BadRequest.New("no such feed")
		}

		begin, end := s.GroupingFeedActivitiesSubspace.FDBRangeKeySelectors()
		if offsetID != "" {
			end = fdb.FirstGreaterOrEqual(s.GroupingFeedActivitiesSubspace.Sub(groupingFeedID, offsetID))
		}
		iter := tr.GetRange(fdb.SelectorRange{Begin: begin, End: end}, fdb.RangeOptions{Limit: int(limit), Reverse: true}).Iterator()
		for iter.Advance() {
			kv, err := iter.Get()
			if err != nil {
				return nil, err
			}
			tuple, err := s.GroupingFeedActivitiesSubspace.Unpack(kv.Key)
			if err != nil {
				return nil, err
			}
			if len(tuple) != 3 {
				return nil, nerrors.Internal.New("wrong key format in GroupingFeedActivitiesSubspace")
			}
			key, ok := tuple[1].(string)
			if !ok {
				return nil, nerrors.Internal.New("wrong key format in GroupingFeedActivitiesSubspace")
			}
			activityID, ok := tuple[2].(string)
			if !ok {
				return nil, nerrors.Internal.New("wrong key format in GroupingFeedActivitiesSubspace")
			}

			bytes, err := tr.Get(s.ActivitiesSubspace.Sub(activityID)).Get()
			if err != nil {
				return nil, err
			}
			if bytes == nil {
				return nil, nerrors.Internal.New("activity doesn't exist")
			}
			activity := &proto.Activity{}
			err = protobuf.Unmarshal(bytes, activity)
			if err != nil {
				return nil, err
			}
			if group, ok := activityGroups[key]; !ok {
				activityGroups[key] = []*proto.Activity{activity}
			} else {
				activityGroups[key] = append(group, activity)
			}
		}
		return nil, nil
	})
	result := make([]*proto.ActivityGroup, 0)
	for key, activities := range activityGroups {
		result = append(result, &proto.ActivityGroup{
			GroupKey:   key,
			Activities: activities,
		})
	}
	return result, err
}

func (s *FeedsStorageFDB) UpdateGroupingFeed(groupingFeedID string, extraData *anypb.Any) (*proto.GroupingFeed, error) {
	feed := &proto.GroupingFeed{}
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		bytes, err := tr.Get(s.GroupingFeedsSubspace.Sub(groupingFeedID)).Get()
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
		tr.Set(s.GroupingFeedsSubspace.Sub(feed.Id), bytes)
		return nil, nil
	})
	return feed, err
}

func (s *FeedsStorageFDB) DeleteGroupingFeed(groupingFeedID string) error {
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		tr.Clear(s.GroupingFeedsSubspace.Sub(groupingFeedID))
		begin, end := s.GroupingFeedActivitiesSubspace.FDBRangeKeys()
		iter := tr.GetRange(fdb.KeyRange{Begin: begin, End: end}, fdb.RangeOptions{}).Iterator()
		for iter.Advance() {
			kv, err := iter.Get()
			if err != nil {
				return nil, err
			}
			tr.Clear(s.ActivityGroupingFeedsSubspace.Sub(kv.Key, groupingFeedID))
		}
		tr.ClearRange(s.GroupingFeedActivitiesSubspace.Sub(groupingFeedID))
		return nil, nil
	})
	return err
}
