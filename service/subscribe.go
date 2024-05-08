package service

import (
	"context"
	"fmt"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/go-kit/log/level"
)

func (s *service) initSubscribe(ctx context.Context) error {
	keys, err := s.repository.Scan(ctx, fmt.Sprintf("%s*", util.RequiringsEndpointKeyPrefix), 15)
	if err != nil {
		return err
	}

	for _, key := range keys {
		endpoints, err := s.repository.GetMemberOfSet(ctx, key)
		if err != nil {
			level.Error(s.log).Log(
				util.LogMessage, "failed to get members of set",
				util.LogError, err,
				util.LogKey, key,
			)
		}
		endpoint := util.GetEndpointFromRequiringsKey(key)
		for _, ep := range endpoints {
			if ep != endpoint {
				s.subscribeMap[ep] = true
				encodedTopic := util.EncodeTopic(ep)
				go s.broker.SubscribeAsync(ctx, encodedTopic, s.repository.SetWithExp)
			}
		}
	}

	return nil
}
