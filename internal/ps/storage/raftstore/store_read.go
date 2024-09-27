// Copyright 2019 The Vearch Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package raftstore

import (
	"context"

	"github.com/vearch/vearch/v3/internal/entity"
	"github.com/vearch/vearch/v3/internal/pkg/log"
	"github.com/vearch/vearch/v3/internal/pkg/vearchlog"
	"github.com/vearch/vearch/v3/internal/proto/vearchpb"
)

func (s *Store) GetDocument(ctx context.Context, readLeader bool, doc *vearchpb.Document, getByDocId bool, next bool) (err error) {
	if err = s.checkReadable(readLeader); err != nil {
		return err
	}
	return s.Engine.Reader().GetDoc(ctx, doc, getByDocId, next)
}

// check this store can read
func (s *Store) checkReadable(readLeader bool) error {
	status := s.Partition.GetStatus()

	if status == entity.PA_CLOSED {
		return vearchlog.LogErrAndReturn(vearchpb.NewError(vearchpb.ErrorEnum_PARTITION_IS_CLOSED, nil))
	}

	if status == entity.PA_INVALID {
		return vearchpb.NewError(vearchpb.ErrorEnum_PARTITION_IS_INVALID, nil)
	}

	if readLeader && status != entity.PA_READWRITE {
		log.Error("checkReadable status: %d , err: %v", status, vearchpb.NewError(vearchpb.ErrorEnum_PARTITION_NOT_LEADER, nil).Error())
		return vearchpb.NewError(vearchpb.ErrorEnum_PARTITION_NOT_LEADER, nil)
	}

	return nil
}

func (s *Store) Search(ctx context.Context, request *vearchpb.SearchRequest, response *vearchpb.SearchResponse) (err error) {
	leader := false
	if request.Head.ClientType == "leader" {
		leader = true
	}
	if err = s.checkReadable(leader); err != nil {
		return err
	}
	err = s.Engine.Reader().Search(ctx, request, response)
	return err
}

func (s *Store) Query(ctx context.Context, request *vearchpb.QueryRequest, response *vearchpb.SearchResponse) (err error) {
	leader := false
	if request.Head.ClientType == "leader" {
		leader = true
	}
	if err = s.checkReadable(leader); err != nil {
		return err
	}
	err = s.Engine.Reader().Query(ctx, request, response)
	return err
}
