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

package e2e

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/icmd"
)

func TestLocalComposeBuild(t *testing.T) {
	c := NewParallelE2eCLI(t, binDir)

	t.Run("build named and unnamed images", func(t *testing.T) {
		// ensure local test run does not reuse previously build image
		c.RunDockerOrExitError("rmi", "build-test_nginx")
		c.RunDockerOrExitError("rmi", "custom-nginx")

		res := c.RunDockerCmd("compose", "--project-directory", "fixtures/build-test", "build")

		res.Assert(t, icmd.Expected{Out: "COPY static /usr/share/nginx/html"})
		c.RunDockerCmd("image", "inspect", "build-test_nginx")
		c.RunDockerCmd("image", "inspect", "custom-nginx")
	})

	t.Run("build with build-arg", func(t *testing.T) {
		// ensure local test run does not reuse previously build image
		c.RunDockerOrExitError("rmi", "build-test_nginx")
		c.RunDockerOrExitError("rmi", "custom-nginx")

		c.RunDockerCmd("compose", "--project-directory", "fixtures/build-test", "build", "--build-arg", "FOO=BAR")

		res := c.RunDockerCmd("image", "inspect", "build-test_nginx")
		res.Assert(t, icmd.Expected{Out: `"FOO": "BAR"`})
	})

	t.Run("build with build-arg set by env", func(t *testing.T) {
		// ensure local test run does not reuse previously build image
		c.RunDockerOrExitError("rmi", "build-test_nginx")
		c.RunDockerOrExitError("rmi", "custom-nginx")

		icmd.RunCmd(c.NewDockerCmd("compose", "--project-directory", "fixtures/build-test", "build", "--build-arg", "FOO"),
			func(cmd *icmd.Cmd) {
				cmd.Env = append(cmd.Env, "FOO=BAR")
			})

		res := c.RunDockerCmd("image", "inspect", "build-test_nginx")
		res.Assert(t, icmd.Expected{Out: `"FOO": "BAR"`})
	})

	t.Run("build as part of up", func(t *testing.T) {
		c.RunDockerOrExitError("rmi", "build-test_nginx")
		c.RunDockerOrExitError("rmi", "custom-nginx")

		res := c.RunDockerCmd("compose", "--project-directory", "fixtures/build-test", "up", "-d")
		t.Cleanup(func() {
			c.RunDockerCmd("compose", "--project-directory", "fixtures/build-test", "down")
		})

		res.Assert(t, icmd.Expected{Out: "COPY static /usr/share/nginx/html"})
		res.Assert(t, icmd.Expected{Out: "COPY static2 /usr/share/nginx/html"})

		output := HTTPGetWithRetry(t, "http://localhost:8070", http.StatusOK, 2*time.Second, 20*time.Second)
		assert.Assert(t, strings.Contains(output, "Hello from Nginx container"))

		c.RunDockerCmd("image", "inspect", "build-test_nginx")
		c.RunDockerCmd("image", "inspect", "custom-nginx")
	})

	t.Run("no rebuild when up again", func(t *testing.T) {
		res := c.RunDockerCmd("compose", "--project-directory", "fixtures/build-test", "up", "-d")

		assert.Assert(t, !strings.Contains(res.Stdout(), "COPY static"), res.Stdout())
	})

	t.Run("rebuild when up --build", func(t *testing.T) {
		res := c.RunDockerCmd("compose", "--workdir", "fixtures/build-test", "up", "-d", "--build")

		res.Assert(t, icmd.Expected{Out: "COPY static /usr/share/nginx/html"})
		res.Assert(t, icmd.Expected{Out: "COPY static2 /usr/share/nginx/html"})
	})

	t.Run("cleanup build project", func(t *testing.T) {
		c.RunDockerCmd("compose", "--project-directory", "fixtures/build-test", "down")
		c.RunDockerCmd("rmi", "build-test_nginx")
		c.RunDockerCmd("rmi", "custom-nginx")
	})
}

func TestLocalComposeBuildStaticDockerfilePath(t *testing.T) {
	c := NewParallelE2eCLI(t, binDir)

	t.Run("build ddev-style compose files", func(t *testing.T) {
		dir, err := ioutil.TempDir("", "ddev")
		assert.NilError(t, err)
		defer os.RemoveAll(dir) //nolint:errcheck

		assert.NilError(t, ioutil.WriteFile(filepath.Join(dir, "compose.yaml"), []byte(`services:
  service1:
    build:
      context: `+dir+`/service1
      dockerfile: Dockerfile
  service2:
    build:
      context: `+dir+`/service2
      dockerfile: `+dir+`/service2/Dockerfile
  `), 0644))

		assert.NilError(t, os.Mkdir(filepath.Join(dir, "service1"), 0700))
		assert.NilError(t, ioutil.WriteFile(filepath.Join(dir, "service1", "Dockerfile"), []byte(`FROM alpine
		RUN echo "hello"
		`), 0644))

		assert.NilError(t, os.Mkdir(filepath.Join(dir, "service2"), 0700))
		assert.NilError(t, ioutil.WriteFile(filepath.Join(dir, "service2", "Dockerfile"), []byte(`FROM alpine
		RUN echo "world"
		`), 0644))

		res := c.RunDockerCmd("compose", "-f", filepath.Join(dir, "compose.yaml"), "build")

		res.Assert(t, icmd.Expected{Out: `RUN echo "hello"`})
		res.Assert(t, icmd.Expected{Out: `RUN echo "world"`})

		c.RunDockerCmd("compose", "-f", filepath.Join(dir, "compose.yaml"), "down", "--rmi", "all")
	})
}
