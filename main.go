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
	"context"

	"github.com/go-sicky/sicky"
	"github.com/go-sicky/sicky/logger"
	"github.com/go-sicky/sicky/server"
	srvGRPC "github.com/go-sicky/sicky/server/grpc"
	"github.com/go-sicky/sicky/service"
	"github.com/go-sicky/sicky/service/standard"
	"svc.biz.room/handler"
)

const (
	AppName   = "svc.biz.room"
	Version   = "latest"
	Branch    = "main"
	Commit    = ""
	BuildTime = ""
)

func main() {
	ctx := context.Background()
	sicky.Init(
		&sicky.Options{
			AppName:   AppName,
			Version:   Version,
			Branch:    Branch,
			Commit:    Commit,
			BuildTime: BuildTime,
			Context:   ctx,
		},
	)
	sicky.ConfigUnmarshal(&config)

	// Logger
	logger.Logger.Level(logger.DebugLevel)

	// GRPC server
	grpcSrv := srvGRPC.New(&server.Options{Name: AppName + "@grpc"}, config.Server.GRPC)
	grpcSrv.Handle(handler.NewGRPCRoom())

	// Service
	svc := standard.New(&service.Options{
		Name:    AppName,
		Version: Version,
		Branch:  Branch,
	}, config.Service)
	svc.Servers(grpcSrv)

	sicky.Run(config.Sicky)
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
