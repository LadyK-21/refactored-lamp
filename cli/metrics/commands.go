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

package metrics

var commandFlags = []string{
	//added to catch scan details
	"--version", "--login",
	// added for build
	"--builder", "--platforms",
}

// Generated with generatecommands/main.go
var managementCommands = []string{
	"help",
	"app",
	"builder",
	"buildx",
	"checkpoint",
	"compose",
	"config",
	"container",
	"context",
	"create",
	"dev",
	"ecs",
	"extension",
	"image",
	"imagetools",
	"key",
	"login",
	"logout",
	"manifest",
	"network",
	"node",
	"plugin",
	"sbom",
	"scan",
	"secret",
	"service",
	"signer",
	"stack",
	"swarm",
	"system",
	"trust",
	"volume",
}

var commands = []string{
	"aci",
	"add",
	"annotate",
	"attach",
	"azure",
	"b",
	"bake",
	"build",
	"bundle",
	"ca",
	"commit",
	"completion",
	"config",
	"connect",
	"convert",
	"cp",
	"create",
	"debug",
	"demote",
	"deploy",
	"df",
	"diff",
	"disable",
	"disconnect",
	"down",
	"du",
	"ecs",
	"enable",
	"events",
	"exec",
	"export",
	"f",
	"generate",
	"history",
	"images",
	"import",
	"info",
	"init",
	"inspect",
	"install",
	"join",
	"join-token",
	"kill",
	"leave",
	"list",
	"load",
	"login",
	"logout",
	"logs",
	"ls",
	"merge",
	"pause",
	"port",
	"promote",
	"prune",
	"ps",
	"pull",
	"push",
	"remove",
	"rename",
	"render",
	"reset",
	"restart",
	"revoke",
	"rm",
	"rmi",
	"rollback",
	"run",
	"save",
	"scale",
	"search",
	"services",
	"set",
	"show",
	"sign",
	"split",
	"start",
	"stats",
	"status",
	"stop",
	"tag",
	"top",
	"ui-source",
	"uninstall",
	"unlock",
	"unlock-key",
	"unpause",
	"up",
	"update",
	"upgrade",
	"use",
	"validate",
	"version",
	"wait",
}
