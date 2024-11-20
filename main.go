/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2024 HereweTech Co.LTD
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

/**
 * @file main.go
 * @package main
 * @author Dr.NP <np@herewe.tech>
 * @since 11/20/2024
 */

package main

import (
	brkNats "github.com/go-sicky/sicky/broker/nats"
	"github.com/go-sicky/sicky/driver"
	"github.com/go-sicky/sicky/logger"
	rgConsul "github.com/go-sicky/sicky/registry/consul"
	"github.com/go-sicky/sicky/runtime"
	"github.com/go-sicky/sicky/server"
	srvGRPC "github.com/go-sicky/sicky/server/grpc"
	"github.com/go-sicky/sicky/service"
	"github.com/go-sicky/sicky/service/sicky"
	"svc.biz.room/handler"
)

const (
	AppName = "svc.biz.room"
	Version = "latest"
)

func main() {
	// Runtime
	runtime.Init(AppName)
	runtime.Config.Unmarshal(&config)

	// Logger
	logger.Logger.Level(logger.DebugLevel)

	// Drivers
	driver.InitDB(config.Driver.DB)
	driver.InitRedis(config.Driver.Redis)

	// GRPC server
	grpcSrv := srvGRPC.New(&server.Options{Name: AppName + "@grpc"}, config.Server.GRPC)
	grpcSrv.Handle(handler.NewGRPCRoom())

	// Broker
	brkNats := brkNats.New(nil, config.Broker.Nats)

	// Registry
	rgConsul := rgConsul.New(nil, config.Registry.Consul)

	// Service
	svc := sicky.New(&service.Options{Name: AppName}, config.Service)
	svc.Servers(grpcSrv)
	svc.Brokers(brkNats)
	svc.Registries(rgConsul)

	service.Run()
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
