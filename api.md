# VMs

Response Types:

- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMLogsResponse">VMLogsResponse</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmsshResponse">VmsshResponse</a>

Methods:

- <code title="get /v0/vms/logs2">client.VMs.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMService.Logs">Logs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMLogsParams">VMLogsParams</a>) (\*<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMLogsResponse">VMLogsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/vms/ssh">client.VMs.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMService.SSH">SSH</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMSSHParams">VMSSHParams</a>) (\*<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmsshResponse">VmsshResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Script

Params Types:

- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#UserDataUnionParam">UserDataUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#UserDataUnion">UserDataUnion</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMScriptNewResponse">VMScriptNewResponse</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMScriptGetResponse">VMScriptGetResponse</a>

Methods:

- <code title="post /v0/vms/script">client.VMs.Script.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMScriptService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMScriptNewParams">VMScriptNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMScriptNewResponse">VMScriptNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/vms/script">client.VMs.Script.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMScriptService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMScriptGetResponse">VMScriptGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Images

Response Types:

- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMImageListResponse">VMImageListResponse</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMImageGetResponse">VMImageGetResponse</a>

Methods:

- <code title="get /v1/vms/images">client.VMs.Images.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMImageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMImageListResponse">VMImageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vms/images/{image_id}">client.VMs.Images.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMImageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, imageID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VMImageGetResponse">VMImageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Nodes

Params Types:

- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#CreateNodesRequestParam">CreateNodesRequestParam</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#ExtendNodeRequestParam">ExtendNodeRequestParam</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeType">NodeType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#AcceleratorType">AcceleratorType</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#ListResponseNode">ListResponseNode</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#Node">Node</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeType">NodeType</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#Status">Status</a>

Methods:

- <code title="post /v1/nodes">client.Nodes.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeNewParams">NodeNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#ListResponseNode">ListResponseNode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/nodes">client.Nodes.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeListParams">NodeListParams</a>) (\*<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#ListResponseNode">ListResponseNode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/nodes/{id}">client.Nodes.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="patch /v1/nodes/{id}/extend">client.Nodes.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeService.Extend">Extend</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeExtendParams">NodeExtendParams</a>) (\*<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#Node">Node</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/nodes/{id}">client.Nodes.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#Node">Node</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/nodes/{id}/redeploy">client.Nodes.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeService.Redeploy">Redeploy</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeRedeployParams">NodeRedeployParams</a>) (\*<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#Node">Node</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/nodes/{id}/release">client.Nodes.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeService.Release">Release</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#Node">Node</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Zones

Response Types:

- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#ZoneListResponse">ZoneListResponse</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#ZoneGetResponse">ZoneGetResponse</a>

Methods:

- <code title="get /v0/zones">client.Zones.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#ZoneService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#ZoneListResponse">ZoneListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/zones/{id}">client.Zones.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#ZoneService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#ZoneGetResponse">ZoneGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
