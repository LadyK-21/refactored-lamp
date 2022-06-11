/*
   Copyright 2020 Docker Compose CLI authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package ecs

import (
	"context"

	"github.com/docker/compose/v2/pkg/api"

	"github.com/docker/compose-cli/utils"
)

func (b *ecsAPIService) Logs(ctx context.Context, projectName string, consumer api.LogConsumer, options api.LogOptions) error {
	if err := checkUnsupportedLogOptions(ctx, options); err != nil {
		return err
	}
	if len(options.Services) > 0 {
		consumer = utils.FilteredLogConsumer(consumer, options.Services)
	}
	err := b.aws.GetLogs(ctx, projectName, consumer.Log, options.Follow)
	return err
}

func checkUnsupportedLogOptions(ctx context.Context, o api.LogOptions) error {
	var errs error
	checks := []struct {
		toCheck, expected interface{}
		option            string
	}{
		{o.Since, "", "since"},
		{o.Tail, "all", "tail"},
		{o.Timestamps, false, "timestamps"},
		{o.Until, "", "until"},
	}
	for _, c := range checks {
		errs = utils.CheckUnsupported(ctx, errs, c.toCheck, c.expected, "logs", c.option)
	}
	return errs
}
