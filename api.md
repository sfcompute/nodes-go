# Vms

Response Types:

- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmLogsResponse">VmLogsResponse</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmSSHResponse">VmSSHResponse</a>

Methods:

- <code title="get /v0/vms/logs2">client.Vms.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmService.Logs">Logs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmLogsParams">VmLogsParams</a>) (<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmLogsResponse">VmLogsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/vms/ssh">client.Vms.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmService.SSH">SSH</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmSSHParams">VmSSHParams</a>) (<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmSSHResponse">VmSSHResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Script

Params Types:

- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#UserDataUnionParam">UserDataUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#UserDataUnion">UserDataUnion</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmScriptNewResponse">VmScriptNewResponse</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmScriptGetResponse">VmScriptGetResponse</a>

Methods:

- <code title="post /v0/vms/script">client.Vms.Script.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmScriptService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmScriptNewParams">VmScriptNewParams</a>) (<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmScriptNewResponse">VmScriptNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/vms/script">client.Vms.Script.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmScriptService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#VmScriptGetResponse">VmScriptGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#ProcurementStatus">ProcurementStatus</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#Status">Status</a>
- <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#Zone">Zone</a>

Methods:

- <code title="post /v1/nodes">client.Nodes.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeNewParams">NodeNewParams</a>) (<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#ListResponseNode">ListResponseNode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/nodes">client.Nodes.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#ListResponseNode">ListResponseNode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/nodes/{id}/extend">client.Nodes.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeService.Extend">Extend</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeExtendParams">NodeExtendParams</a>) (<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#Node">Node</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/nodes/{id}/release">client.Nodes.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#NodeService.Release">Release</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go">sfcnodes</a>.<a href="https://pkg.go.dev/github.com/sfcompute/nodes-go#Node">Node</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
