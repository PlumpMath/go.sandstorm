# Sandstorm - Personal Cloud Sandbox
# Copyright (c) 2016 Sandstorm Development Group, Inc. and contributors
# All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

@0xacf10e4407376d3f;

$import "/capnp/c++.capnp".namespace("sandstorm");

using Identity = import "identity.capnp";
using SystemPersistent = import "supervisor.capnp".SystemPersistent;

interface PersistentIdentity extends (Identity.Identity, SystemPersistent) {}
using Go = import "/go.capnp";
$Go.package("identityimpl");
$Go.import("zenhack.net/go/sandstorm/capnp/identityimpl");
