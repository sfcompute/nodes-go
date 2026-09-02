# Changelog

## 0.1.0-alpha.6 (2026-09-02)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/sfcompute/nodes-go/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* **api:** api update ([70b1ef5](https://github.com/sfcompute/nodes-go/commit/70b1ef549bbcd7ce7509af3651213d4c3ae1173c))
* **api:** api update ([d5661f4](https://github.com/sfcompute/nodes-go/commit/d5661f4c88ecfaeb14c02bbd4875dd124727317f))
* **api:** api update ([ae026a2](https://github.com/sfcompute/nodes-go/commit/ae026a2223a052bd1ee4c1b82cba772f6c5fb465))
* **api:** api update ([e3e81a8](https://github.com/sfcompute/nodes-go/commit/e3e81a871faa2a26d367ed8c90ffc18993bf1b60))
* **api:** api update ([a05fc6b](https://github.com/sfcompute/nodes-go/commit/a05fc6bf9a8bd6f5d84d1fada717f3eadb63c2c3))
* **api:** api update ([510c275](https://github.com/sfcompute/nodes-go/commit/510c275c7fc09a93c3725a62d1c8a8bc7b03247a))
* **api:** api update ([8d84967](https://github.com/sfcompute/nodes-go/commit/8d84967e055d00f63918ba7aa6507d508c5195cd))
* **api:** api update ([3a3ece1](https://github.com/sfcompute/nodes-go/commit/3a3ece13628cf493ef152ce10ddd108307d3bfdc))
* **api:** api update ([afd2178](https://github.com/sfcompute/nodes-go/commit/afd21788395f75b15c2c58dda4b61266c6c9e5a5))
* **api:** api update ([72ee288](https://github.com/sfcompute/nodes-go/commit/72ee288f4d79e0508f8b01065e4bee4e4ee2d9a7))
* **api:** api update ([5c53052](https://github.com/sfcompute/nodes-go/commit/5c530522281441ba93e17b182dcde27e30fb43cc))
* **api:** api update ([8594097](https://github.com/sfcompute/nodes-go/commit/85940976fa7a1fad13d0ee62e4325f7b17a2bfb2))
* **api:** api update ([eb93c6d](https://github.com/sfcompute/nodes-go/commit/eb93c6dcb89945517c8d3e506d80bec1531a8fbd))
* **api:** api update ([0ae17e1](https://github.com/sfcompute/nodes-go/commit/0ae17e1a7c4894100cf2bc6c3a8098eba4e17e35))
* **api:** update config to account for breaking changes ([bca26e2](https://github.com/sfcompute/nodes-go/commit/bca26e229a3a6120f77df90f693989b54d26a7b1))
* **client:** optimize json encoder for internal types ([77a27ee](https://github.com/sfcompute/nodes-go/commit/77a27ee3287988b67c53a7a08111b91a5e08d2de))
* **go:** add default http client with timeout ([e8e7fff](https://github.com/sfcompute/nodes-go/commit/e8e7fff59aeb8bcd5b44703b55c59665e47dde46))
* **internal:** support comma format in multipart form encoding ([ecfc0ce](https://github.com/sfcompute/nodes-go/commit/ecfc0ceb5a993acb87fc73731ba19c0e24349ea9))
* **stlc:** configurable CI runner and private-production-repo support in workflow templates ([db22310](https://github.com/sfcompute/nodes-go/commit/db22310b9a93c44dcb4142a7d44c248b503569d1))
* support setting headers via env ([c14c4ea](https://github.com/sfcompute/nodes-go/commit/c14c4ead498de87c9cc80bba147fbc6f285d64a4))


### Bug Fixes

* better respect format tags from the spec ([0b4c0ec](https://github.com/sfcompute/nodes-go/commit/0b4c0ecae518042d6ceecb8525718fb141e3e1c8))
* fix issue with unmarshaling in some cases ([fa94f5f](https://github.com/sfcompute/nodes-go/commit/fa94f5f06356b76992a572891c762c3fe3559dc0))
* **go:** avoid panic when http.DefaultTransport is wrapped ([57cc619](https://github.com/sfcompute/nodes-go/commit/57cc619fd65be59bbf0d85a3f25a3a9c1a07843a))
* prevent duplicate ? in query params ([179abd8](https://github.com/sfcompute/nodes-go/commit/179abd873e3776a78b5f465b2878981f684c0001))


### Chores

* avoid embedding reflect.Type for dead code elimination ([851e07b](https://github.com/sfcompute/nodes-go/commit/851e07b60b7b90623fb34c9d8e947677806c55ae))
* **ci:** skip lint on metadata-only changes ([59072fa](https://github.com/sfcompute/nodes-go/commit/59072fa7b6eda168b36e5702c27e85cd99ab84cc))
* **ci:** skip uploading artifacts on stainless-internal branches ([d5f51e8](https://github.com/sfcompute/nodes-go/commit/d5f51e8e5330f29e7c05a58fea234a90d3c28a86))
* **ci:** support opting out of skipping builds on metadata-only commits ([046b6e1](https://github.com/sfcompute/nodes-go/commit/046b6e1a7cb29da1c78c90944644f50e5629f304))
* **client:** fix multipart serialisation of Default() fields ([68d8f8e](https://github.com/sfcompute/nodes-go/commit/68d8f8e87b235485f65c079a52cd3c381d9936be))
* **internal:** codegen related update ([0efeaca](https://github.com/sfcompute/nodes-go/commit/0efeacad5ecdbea83c326a17bb7a2f593afdcdcb))
* **internal:** codegen related update ([6905ec5](https://github.com/sfcompute/nodes-go/commit/6905ec5c3e4361acdbb1eda50ebcb869330c0504))
* **internal:** minor cleanup ([d4332e4](https://github.com/sfcompute/nodes-go/commit/d4332e4aeae3e5a569fc586ab430182cb705de17))
* **internal:** more robust bootstrap script ([1aad79a](https://github.com/sfcompute/nodes-go/commit/1aad79aad4eb6f35fc6a72077c7f4f4ac2299ba4))
* **internal:** support default value struct tag ([9be56ed](https://github.com/sfcompute/nodes-go/commit/9be56edd00c8e78ee7d3cd56604eeebe63df376b))
* **internal:** tweak CI branches ([f1ab274](https://github.com/sfcompute/nodes-go/commit/f1ab27477b86808c8942b6ce899803ac150c8b16))
* **internal:** update gitignore ([c997311](https://github.com/sfcompute/nodes-go/commit/c9973115ba8dcfb4eb10f7cab9c438c7729779fb))
* **internal:** use explicit returns ([f740ed8](https://github.com/sfcompute/nodes-go/commit/f740ed8beb5a3e52d0c879e31a9ed47f85af233d))
* **internal:** use explicit returns in more places ([f8f3570](https://github.com/sfcompute/nodes-go/commit/f8f3570dcb58aaecc9139b4bf4e93c1a77f4a711))
* redact api-key headers in debug logs ([ce03df6](https://github.com/sfcompute/nodes-go/commit/ce03df6b12d63a0711fd15049397c4e6f69c4071))
* remove unnecessary error check for url parsing ([9a31c22](https://github.com/sfcompute/nodes-go/commit/9a31c22ba410290e150a66b0f7bec72c9ae131be))
* update docs for api:"required" ([afd6c27](https://github.com/sfcompute/nodes-go/commit/afd6c27ac6b3ae51794c4098293b9f1a676cf2a2))

## 0.1.0-alpha.5 (2026-02-25)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/sfcompute/nodes-go/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* **api:** api update ([a0a7fa2](https://github.com/sfcompute/nodes-go/commit/a0a7fa24667407359b24deb9a6391599873cfc4e))
* **api:** api update ([531d77e](https://github.com/sfcompute/nodes-go/commit/531d77efc30775a747cf3a96320e1a85c172fdb2))
* **api:** api update ([dff5f7f](https://github.com/sfcompute/nodes-go/commit/dff5f7f8daf4d7f353962442bf41e2d55ff028a1))
* **api:** api update ([ca18f27](https://github.com/sfcompute/nodes-go/commit/ca18f27daf6688518636b18a4851308dcdae0db6))
* **api:** api update ([894995d](https://github.com/sfcompute/nodes-go/commit/894995db7b4a6347ef91fed5db242fbb14df34ef))
* **api:** api update ([afa8e35](https://github.com/sfcompute/nodes-go/commit/afa8e35b48706ac111a768c9a7a81c59aa1602b0))
* **api:** api update ([f77e6d0](https://github.com/sfcompute/nodes-go/commit/f77e6d07e261819f07b68b9241218f42a0ab36a1))
* **api:** api update ([3068ae6](https://github.com/sfcompute/nodes-go/commit/3068ae6937f3d6f6f5b36f8ac7c7e318866c3ce1))
* **api:** api update ([a4e5181](https://github.com/sfcompute/nodes-go/commit/a4e5181e44815a66b891ecd41790c9dab76c0375))
* **api:** api update ([6a064fd](https://github.com/sfcompute/nodes-go/commit/6a064fdc947dd668159e8643e7cab0689254d4c6))
* **api:** api update ([dd70400](https://github.com/sfcompute/nodes-go/commit/dd70400abff7e7332cf7a7f3dee59e9e2e4b61b5))
* **client:** add a convenient param.SetJSON helper ([a7293f3](https://github.com/sfcompute/nodes-go/commit/a7293f32d7d9f676d53e6200ca7bd82f603029bb))
* **encoder:** support bracket encoding form-data object members ([bf75f99](https://github.com/sfcompute/nodes-go/commit/bf75f99a26c353742229dabb561684545d0835c7))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([b16eec0](https://github.com/sfcompute/nodes-go/commit/b16eec05809da7c51496804b92be7f79a664860a))
* **docs:** add missing pointer prefix to api.md return types ([6dd1ead](https://github.com/sfcompute/nodes-go/commit/6dd1eadf389b5a9adf98195f624f0bba54f0b686))
* **encoder:** correctly serialize NullStruct ([ed7b95d](https://github.com/sfcompute/nodes-go/commit/ed7b95dc7003020ff3f2738321d023ae3599e9e0))
* **mcp:** correct code tool API endpoint ([2eeb64b](https://github.com/sfcompute/nodes-go/commit/2eeb64be8909979b8cb3524a077cb764c85a1f21))
* rename param to avoid collision ([f221c75](https://github.com/sfcompute/nodes-go/commit/f221c7569b746c43bd46c4ab3c1615b30fc0c05c))
* skip usage tests that don't work with Prism ([61d03bd](https://github.com/sfcompute/nodes-go/commit/61d03bd3c1f7772a74bef48676020df27c371ab1))


### Chores

* add float64 to valid types for RegisterFieldValidator ([499e663](https://github.com/sfcompute/nodes-go/commit/499e663b659fa00c013ebe3db1e0622b5e2a6a51))
* elide duplicate aliases ([9e83189](https://github.com/sfcompute/nodes-go/commit/9e83189fd37d851e8441c313c738656a72760483))
* **internal:** codegen related update ([44eb5d0](https://github.com/sfcompute/nodes-go/commit/44eb5d0a3f745631aebb3b0d2986a049302131fe))
* **internal:** codegen related update ([e40a3de](https://github.com/sfcompute/nodes-go/commit/e40a3debee22efa7015e5230d14374245af79b1f))
* **internal:** move custom custom `json` tags to `api` ([d68a748](https://github.com/sfcompute/nodes-go/commit/d68a7480dd7d4eda8fefa348d728cabf53ad671e))
* **internal:** remove mock server code ([f28644a](https://github.com/sfcompute/nodes-go/commit/f28644aaee0119dccf2223264f06703f00567327))
* **internal:** update `actions/checkout` version ([c853455](https://github.com/sfcompute/nodes-go/commit/c853455bd8e0c977606bc29049cfc84b2caea449))
* update mock server docs ([588eb6a](https://github.com/sfcompute/nodes-go/commit/588eb6aca1879b14a61976f013b0155140dc1971))

## 0.1.0-alpha.4 (2025-12-01)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/sfcompute/nodes-go/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** add .zones SDK methods ([7d0d7e8](https://github.com/sfcompute/nodes-go/commit/7d0d7e8e8c2902b64ac32afc36c9ab3c2d759c3f))
* **api:** api update ([9f94ea5](https://github.com/sfcompute/nodes-go/commit/9f94ea57dcaabbec647d257634ee5c550bfe823b))
* **api:** api update ([1342c39](https://github.com/sfcompute/nodes-go/commit/1342c3989b487bd7725ab95ba720c7da200c6899))
* **api:** api update ([02e6349](https://github.com/sfcompute/nodes-go/commit/02e6349b1823ed96fecf3353231e32b00ef3989b))
* **api:** api update ([26ba165](https://github.com/sfcompute/nodes-go/commit/26ba165c4e22a00f69c89f53edf7b0445057594f))


### Bug Fixes

* **client:** correctly specify Accept header with */* instead of empty ([02ed3d4](https://github.com/sfcompute/nodes-go/commit/02ed3d4181ab9b9c97e85fb3dd0b9ded24489f40))


### Chores

* bump gjson version ([4190bc4](https://github.com/sfcompute/nodes-go/commit/4190bc4b5cf17ed7fda70ff9e8ccab34eb33d477))
* **internal:** codegen related update ([766d0aa](https://github.com/sfcompute/nodes-go/commit/766d0aaf277f7c0733d15ba33111d5e9519660a8))
* **internal:** grammar fix (it's -&gt; its) ([1a198ba](https://github.com/sfcompute/nodes-go/commit/1a198bad6aa4689cecfaccec1953310a60a1a47c))

## 0.1.0-alpha.3 (2025-10-13)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/sfcompute/nodes-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** add vm images resources and update formatting ([273c45f](https://github.com/sfcompute/nodes-go/commit/273c45f75e75104ad10eb1eab1bdb7acc3596cce))
* **api:** api update ([97607e0](https://github.com/sfcompute/nodes-go/commit/97607e02d439b4fe1463a483d362ab5bf65fd708))
* **api:** api update ([5788fb0](https://github.com/sfcompute/nodes-go/commit/5788fb0d297e3c4f79e76e1f2369bd086dad610b))
* **api:** api update ([59b16bc](https://github.com/sfcompute/nodes-go/commit/59b16bcc748ffea703644d6332f33a79196aeeee))
* **api:** api update ([d9c2daa](https://github.com/sfcompute/nodes-go/commit/d9c2daa0151d664a5bbff0609d0a6e1e6db3b8c8))
* **api:** disable retries ([8c19a86](https://github.com/sfcompute/nodes-go/commit/8c19a865b584b8dc40feecd1a26e3d8d9a013230))


### Bug Fixes

* **api:** remove undocumented endpoints, add list endpoint ([ee795ee](https://github.com/sfcompute/nodes-go/commit/ee795ee93f88b5c379d9ce834356c59fae43e162))
* **internal:** unmarshal correctly when there are multiple discriminators ([2836b68](https://github.com/sfcompute/nodes-go/commit/2836b6828c5f60e8c0c4e8892bbefa54e54b64b8))
* use slices.Concat instead of sometimes modifying r.Options ([6adf0a1](https://github.com/sfcompute/nodes-go/commit/6adf0a14ecd1b5d20924b1b26816af8c5ed270ad))


### Chores

* bump minimum go version to 1.22 ([90c5664](https://github.com/sfcompute/nodes-go/commit/90c566457cad4e6964d6217feeb21eb699d2c385))
* configure new SDK language ([42f31df](https://github.com/sfcompute/nodes-go/commit/42f31df1c143e56ca36cea4b5de2b937ac000c7c))
* do not install brew dependencies in ./scripts/bootstrap by default ([5983889](https://github.com/sfcompute/nodes-go/commit/5983889bfd8681434b39372a0d58fe60115ff267))
* update more docs for 1.22 ([82733b2](https://github.com/sfcompute/nodes-go/commit/82733b2eff62084fa0b0e415e8a3ad6e1f8f546c))

## 0.1.0-alpha.2 (2025-09-05)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/sfcompute/nodes-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** api update ([ed6f2e9](https://github.com/sfcompute/nodes-go/commit/ed6f2e905fabfd0fb526e5bb5a69565a76d86244))

## 0.1.0-alpha.1 (2025-09-04)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/sfcompute/nodes-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** add `get` method ([d6f4ab4](https://github.com/sfcompute/nodes-go/commit/d6f4ab49e0ca0454aaeb03a638b15b5df5e2262d))
* **api:** api update ([99f2770](https://github.com/sfcompute/nodes-go/commit/99f2770c5b4fd68074f1958219e1ccc186e94f63))
* **api:** api update ([7878508](https://github.com/sfcompute/nodes-go/commit/7878508378486e556461a3b27ff266977e102039))
* **api:** api update ([05632b4](https://github.com/sfcompute/nodes-go/commit/05632b40d2e7138a0b936bdd7e89ebe0ce6b1d26))
* **api:** api update ([923efee](https://github.com/sfcompute/nodes-go/commit/923efee077f1eb7e99c269be72979b34d540ad4c))
* **api:** api update ([93ad1a5](https://github.com/sfcompute/nodes-go/commit/93ad1a595748fd9f5cceca985c86b6f5d3d36d92))
* **api:** api update ([c30069f](https://github.com/sfcompute/nodes-go/commit/c30069f47d688ef8357ab9be7df3a3dc662b5063))
* **api:** api update ([00fffab](https://github.com/sfcompute/nodes-go/commit/00fffab84d93344a46ca156fb947881757e431ca))
* **api:** api update ([f419ced](https://github.com/sfcompute/nodes-go/commit/f419ced0cd6952e529b1bcc031f37d2db0bad646))
* **api:** api update ([834cc67](https://github.com/sfcompute/nodes-go/commit/834cc67601b75f3312c1c7b039f60540b145d6b8))
* **api:** manual updates ([02c41ab](https://github.com/sfcompute/nodes-go/commit/02c41ab561a1ebb5bc54dee1ccc7b7df1433ccbd))
* **api:** manually add new Nodes API models ([2e782e3](https://github.com/sfcompute/nodes-go/commit/2e782e38c064a41bbddbe8b32c077979e1312020))
* **api:** re-add removed operations ([a7829d3](https://github.com/sfcompute/nodes-go/commit/a7829d30b907b2242db5da128c1ee8181564ab24))
* **api:** revert custom ErrorConfig ([2d1f7c8](https://github.com/sfcompute/nodes-go/commit/2d1f7c84ea22ee322a08b9a1911e202ab16659c4))
* **api:** update via SDK Studio ([5a3ea36](https://github.com/sfcompute/nodes-go/commit/5a3ea36fb89e3c4cfcbe3b612644f98a5ae33e31))
* **client:** support optional json html escaping ([4e27fc0](https://github.com/sfcompute/nodes-go/commit/4e27fc0cbfd8829f304a4836fd8c681c8aca82b5))


### Bug Fixes

* **client:** process custom base url ahead of time ([59732a1](https://github.com/sfcompute/nodes-go/commit/59732a1b52f5adf2d65e59f781528be6d0f797f0))
* close body before retrying ([a2c02f0](https://github.com/sfcompute/nodes-go/commit/a2c02f0cd347e3d9bb01eedb6fa99f03dd527be0))
* query param arrays are repeated ([29c2a31](https://github.com/sfcompute/nodes-go/commit/29c2a31a047f5a3fec6f3dc66da6748057969625))


### Chores

* configure new SDK language ([6e68d64](https://github.com/sfcompute/nodes-go/commit/6e68d64dc949d10023baf84985a7aa35be588a0e))
* **internal:** codegen related update ([dc39de2](https://github.com/sfcompute/nodes-go/commit/dc39de2d48b4c32b3e2d6499b64ff9114395cc1f))
* **internal:** codegen related update ([8bc1bd2](https://github.com/sfcompute/nodes-go/commit/8bc1bd2ba42962466ab9b90728fb4daa43c6573c))
* **internal:** update comment in script ([765fd0e](https://github.com/sfcompute/nodes-go/commit/765fd0e2969527765ab268ca4f33be42bb08deea))
* lint tests in subpackages ([4ec488f](https://github.com/sfcompute/nodes-go/commit/4ec488f47f3be8686f911403dc5f28159e0fa055))
* update @stainless-api/prism-cli to v5.15.0 ([625d237](https://github.com/sfcompute/nodes-go/commit/625d2373fa5bcd39da1074b7df386a89118e0a9f))
* update SDK settings ([ffa4444](https://github.com/sfcompute/nodes-go/commit/ffa4444199a4a95811ff66907dc56d71c356f921))
