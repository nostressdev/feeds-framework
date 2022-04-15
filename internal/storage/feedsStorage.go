package storage

import (
	"github.com/nostressdev/feeds-framework/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type FeedsStorage interface {
	AddActivity(feedID, objectID, userID, activityType string, time int64, redirectTo []string, extraData *anypb.Any) (*proto.Activity, error)
	AddExistingActivity(feedID, activityID string) (*proto.Activity, error)
	GetActivity(activityID string) (*proto.Activity, error)
	GetActivityByObjectID(objectID string) (*proto.Activity, error)
	UpdateActivity(activityID string, extraData *anypb.Any) (*proto.Activity, error)
	DeleteActivity(activityID string) error
	CreateFeed(userID string, extraData *anypb.Any) (*proto.Feed, error)
	GetFeed(feedID string) (*proto.Feed, error)
	GetFeedActivities(feedID string, limit int64, offsetID string) ([]*proto.Activity, error)
	UpdateFeed(feedID string, extraData *anypb.Any) (*proto.Feed, error)
	DeleteFeed(feedID string) error
	CreateCollection(collectionID string, deletingType proto.DeletingType) (*proto.Collection, error)
	CreateObject(collectionID, objectID string, data *anypb.Any) (*proto.Object, error)
	GetObject(collectionID, objectID string) (*proto.Object, error)
	UpdateObject(collectionID, objectID string, data *anypb.Any) (*proto.Object, error)
	DeleteObject(collectionID, objectID string) error
	CreateReaction(activityType, userID, linkedActivityID string, time int64, extraData *anypb.Any) (*proto.Activity, error)
	AddReaction(feedID, reactionID string) error
	GetReaction(reactionID string) (*proto.Activity, error)
	GetActivityReactions(activityID string, limit int64, offsetID string) ([]*proto.Activity, error)
	UpdateReaction(reactionID string, extraData *anypb.Any) (*proto.Activity, error)
	DeleteReaction(reactionID string) error
	CreateGroupingFeed(userID, keyFormat string, extraData *anypb.Any) (*proto.GroupingFeed, error)
	GetGroupingFeed(groupingFeedID string) (*proto.GroupingFeed, error)
	GetGroupingFeedActivities(groupingFeedID string, limit int64, offsetID string) ([]*proto.ActivityGroup, error)
	UpdateGroupingFeed(groupingFeedID string, extraData *anypb.Any) (*proto.GroupingFeed, error)
	DeleteGroupingFeed(groupingFeedID string) error
}
