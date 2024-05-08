package service

import (
	"context"
	"fmt"
	"github.com/daffarg/distributed-cascading-cb/util"
)

func (s *service) initSubscribe(ctx context.Context) error {
	keys, err := s.repository.Scan(ctx, fmt.Sprintf("%s*", util.RequiringsEndpointKeyPrefix), 15)
	if err != nil {
		return err
	}

	for _, key := range keys {
		topic := util.GetEndpointFromRequiringsKey(key)
		encodedTopic := util.EncodeTopic(topic)
		go s.broker.SubscribeAsync(ctx, encodedTopic, s.repository.SetWithExp)
	}

	return nil
}
