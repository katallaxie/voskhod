// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package agent

import (
	"context"
	"fmt"

	"github.com/katallaxie/voskhod/config"
	"github.com/katallaxie/voskhod/docker/dockerapi"

	log "github.com/sirupsen/logrus"
)

var _ Agent = (*agent)(nil)

// New is returning a new agent
func New(ctx context.Context, cfg *config.Config) Agent {
	return &agent{
		cfg: cfg,
		ctx: ctx,
	}
}

// Start is starting the agent
func (a *agent) Start() func() error {
	// set custom logger for the agent itself
	a.logger = log.WithFields(log.Fields{})

	return func() error {
		var err error

		// connect docker client
		dc, err := dockerclient.New()
		if err != nil {
			return err
		}

		a.dc = dc // assign the client to the agent

		// init hearbeat for docker client,
		// because we do not know if the client still live

		a.logger.Infof(fmt.Sprintf("Agent succesfully started ..."))

		// just have one channel to end it,\
		// so not using something differne
		<-a.ctx.Done()

		// cleanup
		err = a.Stop()

		// noop
		return err
	}
}

// Stop is actually stopping the agent and tearing down everything.
// Cleaning up the mess.
func (a *agent) Stop() error {
	return a.dc.Stop() // nothing more right now
}
