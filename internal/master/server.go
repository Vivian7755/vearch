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

package master

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/vearch/vearch/v3/internal/client"
	"github.com/vearch/vearch/v3/internal/config"
	"github.com/vearch/vearch/v3/internal/entity"
	"github.com/vearch/vearch/v3/internal/pkg/errutil"
	"github.com/vearch/vearch/v3/internal/pkg/log"
	"github.com/vearch/vearch/v3/internal/proto/vearchpb"
	"go.etcd.io/etcd/server/v3/embed"
	"go.etcd.io/etcd/server/v3/etcdserver"
)

type Server struct {
	etcCfg     *embed.Config
	client     *client.Client
	etcdServer *embed.Etcd
	ctx        context.Context
}

func NewServer(ctx context.Context) (*Server, error) {
	// log.Regist(vearchlog.NewVearchLog(config.Conf().GetLogDir(config.Master), "Master", config.Conf().GetLevel(config.Master), false))
	//Logically, this code should not be executed, because if the local master is not found, it will panic
	if config.Conf().Masters.Self() == nil {
		return nil, vearchpb.NewError(vearchpb.ErrorEnum_CONFIG_ERROR, fmt.Errorf("master config is null"))
	}

	var server *Server
	// manage etcd by yourself
	if config.Conf().Global.SelfManageEtcd {
		// no vearch etcd cfg
		server = &Server{ctx: ctx}
	} else {
		// manage etcd by vearch
		cfg, err := config.Conf().GetEmbed()
		if err != nil {
			return nil, err
		}
		if err := os.MkdirAll(cfg.Dir, os.ModePerm); err != nil {
			return nil, err
		}
		server = &Server{etcCfg: cfg, ctx: ctx}
	}
	return server, nil
}

func (s *Server) Start() (err error) {
	//process panic
	defer errutil.CatchError(&err)
	//start api server
	log.Debug("master start ...")

	// if vearch manage etcd then start it
	if !config.Conf().Global.SelfManageEtcd {
		//start etcd server
		s.etcdServer, err = embed.StartEtcd(s.etcCfg)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		defer s.etcdServer.Close()
		select {
		case <-s.etcdServer.Server.ReadyNotify():
			log.Info("Server is ready!")
		case <-time.After(60 * time.Second):
			s.etcdServer.Server.Stop() // trigger a shutdown
			log.Error("Server took too long to start!")
			return vearchpb.NewError(vearchpb.ErrorEnum_INTERNAL_ERROR, fmt.Errorf("etcd start timeout"))
		}
	}

	s.client, err = client.NewClient(config.Conf())
	if err != nil {
		return err
	}
	service, err := newMasterService(s.client)
	if err != nil {
		return err
	}

	monitorService := &monitorService{}
	if config.Conf().Global.SelfManageEtcd {
		monitorService = newMonitorService(service, &etcdserver.EtcdServer{})
	} else {
		monitorService = newMonitorService(service, s.etcdServer.Server)
	}

	if !log.IsDebugEnabled() {
		gin.SetMode(gin.ReleaseMode)
	}

	// start http server

	engine := gin.New()

	ExportToClusterHandler(engine, service, s)

	ExportToMonitorHandler(engine, monitorService)

	//register monitor

	go func() {
		if err := engine.Run(":" + cast.ToString(config.Conf().Masters.Self().ApiPort)); err != nil {
			panic(err)
		}
	}()

	// add root user
	root := entity.RootName
	userInfo := &entity.User{
		Name:     root,
		Password: &config.Conf().Global.Signkey,
		RoleName: &root,
	}
	if _, err := service.queryUserService(s.ctx, userInfo.Name, true); err != nil {
		log.Debug("query root user : %s", err.Error())
		if err := service.createUserService(s.ctx, userInfo, false); err != nil {
			log.Debug("create root user : %s", err.Error())
			// check again
			_, err := service.queryUserService(s.ctx, userInfo.Name, true)
			if err != nil {
				log.Error("query root user : %s", err.Error())
				panic(err)
			}
		} else {
			log.Info("root user create success")
		}
	} else {
		log.Info("root user already exist")
	}

	// start watch server
	err = s.WatchServerJob(s.ctx, s.client)
	errutil.ThrowError(err)
	log.Debug("start WatchServerJob success!")
	if !config.Conf().Global.SelfManageEtcd {
		return <-s.etcdServer.Err()
	}

	return nil
}

func (s *Server) Stop() {
	log.Info("master shutdown... start")
	s.etcdServer.Server.Stop()
	log.Info("master shutdown... end")
}
