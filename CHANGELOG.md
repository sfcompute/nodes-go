# Changelog

## 0.1.0-alpha.6 (2026-03-07)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/sfcompute/nodes-go/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([d5f51e8](https://github.com/sfcompute/nodes-go/commit/d5f51e8e5330f29e7c05a58fea234a90d3c28a86))
* **internal:** codegen related update ([0efeaca](https://github.com/sfcompute/nodes-go/commit/0efeacad5ecdbea83c326a17bb7a2f593afdcdcb))
* **internal:** codegen related update ([6905ec5](https://github.com/sfcompute/nodes-go/commit/6905ec5c3e4361acdbb1eda50ebcb869330c0504))

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
