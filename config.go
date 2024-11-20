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
 * @file config.go
 * @package main
 * @author Dr.NP <np@herewe.tech>
 * @since 10/17/2024
 */

package main

import (
	brkNats "github.com/go-sicky/sicky/broker/nats"
	brkNsq "github.com/go-sicky/sicky/broker/nsq"
	"github.com/go-sicky/sicky/driver"
	rgConsul "github.com/go-sicky/sicky/registry/consul"
	rgMdns "github.com/go-sicky/sicky/registry/mdns"
	srvGRPC "github.com/go-sicky/sicky/server/grpc"
	srvHTTP "github.com/go-sicky/sicky/server/http"
	srvWebsocket "github.com/go-sicky/sicky/server/websocket"
	"github.com/go-sicky/sicky/service/sicky"
)

type ConfigDef struct {
	Server struct {
		GRPC      *srvGRPC.Config      `json:"grpc" yaml:"grpc" mapstructure:"grpc"`
		HTTP      *srvHTTP.Config      `json:"http" yaml:"http" mapstructure:"http"`
		Websocket *srvWebsocket.Config `json:"websocket" yaml:"websocket" mapstructure:"websocket"`
	} `json:"server" yaml:"server" mapstructure:"server"`
	Broker struct {
		Nats *brkNats.Config `json:"nats" yaml:"nats" mapstructure:"nats"`
		Nsq  *brkNsq.Config  `json:"nsq" yaml:"nsq" mapstructure:"nsq"`
	} `json:"broker" yaml:"broker" mapstructure:"broker"`
	Registry struct {
		Consul *rgConsul.Config `json:"consul" yaml:"consul" mapstructure:"consul"`
		Mdns   *rgMdns.Config   `json:"mdns" yaml:"mdns" mapstructure:"mdns"`
	} `json:"registry" yaml:"registry" mapstructure:"registry"`
	Service *sicky.Config `json:"service" yaml:"service" mapstructure:"service"`
	Driver  struct {
		DB    *driver.DBConfig    `json:"db" yaml:"db" mapstructure:"db"`
		Redis *driver.RedisConfig `json:"redis" yaml:"redis" mapstructure:"redis"`
		Nats  *driver.NatsConfig  `json:"nats" yaml:"nats" mapstructure:"nats"`
	} `json:"driver" yaml:"driver" mapstructure:"driver"`
}

var (
	config ConfigDef
)

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
