# Changelog

## [0.14.1](https://github.com/canonical/identity-platform-login-ui/compare/v0.14.0...v0.14.1) (2024-06-17)


### Bug Fixes

* group -&gt; parent/child relationship + add spaces for readability ([da553c4](https://github.com/canonical/identity-platform-login-ui/commit/da553c46155e078dd241a8366e547f349285447f))
* update docker compose ([533a70c](https://github.com/canonical/identity-platform-login-ui/commit/533a70caee68fd0028f29510a510fcce9d306584))

## [0.14.0](https://github.com/canonical/identity-platform-login-ui/compare/v0.13.0...v0.14.0) (2024-04-30)


### Features

* add error handling for login flow WD-10252 ([f96567e](https://github.com/canonical/identity-platform-login-ui/commit/f96567e7901b8b41c9b5cabff09bf49139a8ddc8))
* return common errors to ui ([5828a6a](https://github.com/canonical/identity-platform-login-ui/commit/5828a6a28e6b36dd45ae53a5bdee7ec8d895848f))
* support password method ([f8842f0](https://github.com/canonical/identity-platform-login-ui/commit/f8842f0d7edc386c8ac7a29064c48487facb65e9))


### Bug Fixes

* fix path ([e7ffc03](https://github.com/canonical/identity-platform-login-ui/commit/e7ffc03d718fb7507de3c162009b54ea2ed8b595))
* password login method WD-10252 ([1e98e09](https://github.com/canonical/identity-platform-login-ui/commit/1e98e094dd37e1e2865848ed3794e09f227db3ce))
* use relative path for callback ([4aa6c83](https://github.com/canonical/identity-platform-login-ui/commit/4aa6c83ea98ce0905f23e60626258dafade874cc))

## [0.13.0](https://github.com/canonical/identity-platform-login-ui/compare/v0.12.0...v0.13.0) (2024-04-16)


### Features

* add success screen for device flow WD-10251 ([e9b77db](https://github.com/canonical/identity-platform-login-ui/commit/e9b77dba63dfdb622460bf1c6a6add1aff923aab))


### Bug Fixes

* add device flow logic to hydra client ([ca4a1f9](https://github.com/canonical/identity-platform-login-ui/commit/ca4a1f99c999ef57a42bb4b5c3c95ff11a891487))
* add error handling to the device_code page ([666f95b](https://github.com/canonical/identity-platform-login-ui/commit/666f95b6ad4e38bfc0c07879f0ccc626c6edcbf7))
* **CheckAllowedProviders:** fix OAuthKeeper case ([24212f4](https://github.com/canonical/identity-platform-login-ui/commit/24212f4decc0c5e76e74a21ffc21eab8eae08feb))
* implement device_code backend logic ([06bb0b2](https://github.com/canonical/identity-platform-login-ui/commit/06bb0b29105aab10e87ada456d932899f7c8ef3e))
* pass error response from hydra to UI ([b165f43](https://github.com/canonical/identity-platform-login-ui/commit/b165f43d8e3e83d6bfcf595088c94497203b3bd4))
* typo in Oathkeeper name ([5f7634d](https://github.com/canonical/identity-platform-login-ui/commit/5f7634dca31afb854be4ff0254659c5c51859d2d))
* update device_code page ([4ca0d13](https://github.com/canonical/identity-platform-login-ui/commit/4ca0d138e4eb850a5e10cf5ae1ea384035b105c8))
* update hydra config ([cb1e049](https://github.com/canonical/identity-platform-login-ui/commit/cb1e0493d7c95c0b1b980c5a18bc3bdcede96110))

## [0.12.0](https://github.com/canonical/identity-platform-login-ui/compare/v0.11.3...v0.12.0) (2024-01-31)


### Features

* Add authorization logic for allowed providers ([5da85b7](https://github.com/canonical/identity-platform-login-ui/commit/5da85b785a49e11600ce8f342c80ecd20bbd38bf))
* Add authorizer struct ([c6170a2](https://github.com/canonical/identity-platform-login-ui/commit/c6170a209c49003dbf19ff277a0d2b3e1b268a53))
* add create-fga-model CLI command ([42f47df](https://github.com/canonical/identity-platform-login-ui/commit/42f47df508312f4a28a3a0d4a71859d40c053ed4))
* add fallback logo for identity providers ([ca957ee](https://github.com/canonical/identity-platform-login-ui/commit/ca957ee6db4f20f7b0c89fe172b2b2f9d446cdb8))
* Add openfga client ([8628fa6](https://github.com/canonical/identity-platform-login-ui/commit/8628fa6a41805519cc99b4aaea94d3f857d7079e))
* remove okta logo from login providers list ([88bb621](https://github.com/canonical/identity-platform-login-ui/commit/88bb621c8cd82334d91219c8f5c8d983fe81fc26))
* use cobra for CLI ([50963a3](https://github.com/canonical/identity-platform-login-ui/commit/50963a396ad4c64023c4e5d28b53b4a456f71eb4))
* use new design and add user flows as dummy pages WD-8469 ([2ac0e1a](https://github.com/canonical/identity-platform-login-ui/commit/2ac0e1addcad6ef899fcdb5d10aee653da7cb019))


### Bug Fixes

* add noop clients ([5b6f62a](https://github.com/canonical/identity-platform-login-ui/commit/5b6f62a9ce35cdd5c79463c34fe780a64aa3fbe1))
* bump otel/trace version ([d2f65fc](https://github.com/canonical/identity-platform-login-ui/commit/d2f65fc35bb53b938beeca70132851e3674ec217))
* readme typo ([80de6dd](https://github.com/canonical/identity-platform-login-ui/commit/80de6dd598c6a5b778ea43123bf9bc32b185aec1))
* switch to zap nop logger ([57b37f8](https://github.com/canonical/identity-platform-login-ui/commit/57b37f8b7102b4b62057f27c9648e6546e9a0e67))
* update rockcraft.yaml ([c595939](https://github.com/canonical/identity-platform-login-ui/commit/c595939e6ebffa562fa9597edfc9d677dda9280c))
* update ui ([e862475](https://github.com/canonical/identity-platform-login-ui/commit/e8624753928602ff5664d8bdc601705c2a2cdb76))
* Use github as test provider ([71eb388](https://github.com/canonical/identity-platform-login-ui/commit/71eb388594889c4707a1703fece37c41ea564b09))
* Use label to generate button text ([0494160](https://github.com/canonical/identity-platform-login-ui/commit/0494160446f5320989224e8b40b6aaa749eed7a1))

## [0.11.3](https://github.com/canonical/identity-platform-login-ui/compare/v0.11.2...v0.11.3) (2023-11-01)


### Bug Fixes

* **deps:** update dependency @canonical/react-components to v0.47.1 ([529c5d6](https://github.com/canonical/identity-platform-login-ui/commit/529c5d676039cb47b2e6f45404c59800e8326196))
* **deps:** update dependency vanilla-framework to v4.4.0 ([b9e3b03](https://github.com/canonical/identity-platform-login-ui/commit/b9e3b0395cbec8fde99705bdb4562979c1eff0a1))
* **deps:** update dependency vanilla-framework to v4.5.0 ([bd40882](https://github.com/canonical/identity-platform-login-ui/commit/bd40882423ec6fed74ddd9aa77d6aa236a7315d3))
* downgrade Kratos sdk ([b36de36](https://github.com/canonical/identity-platform-login-ui/commit/b36de3666f7732f26df52f07dafa9bda8090287f))

## [0.11.2](https://github.com/canonical/identity-platform-login-ui/compare/v0.11.1...v0.11.2) (2023-10-04)


### Bug Fixes

* add app version config ([78e93b5](https://github.com/canonical/identity-platform-login-ui/commit/78e93b5fd572c9a988bb1ed0f56989e51dd38045))
* add flag parsing logic ([39123a1](https://github.com/canonical/identity-platform-login-ui/commit/39123a13be19dfef2ddf6c88dc78d625ed925ec0))
* **deps:** update go deps ([71edbe7](https://github.com/canonical/identity-platform-login-ui/commit/71edbe7273671004547899189e6be097853c3939))
* **deps:** update go deps to v1.19.0 ([1a880ce](https://github.com/canonical/identity-platform-login-ui/commit/1a880cefbf2bf8dfe76f8e706281677425e4b097))
* **deps:** update module github.com/ory/kratos-client-go to v1 ([8a4894a](https://github.com/canonical/identity-platform-login-ui/commit/8a4894a5fd8d5838eb4121575e37015077554fba))
* **deps:** update module github.com/prometheus/client_golang to v1.17.0 ([1066d13](https://github.com/canonical/identity-platform-login-ui/commit/1066d13023735222f2e5505c36d4c2869c04dadf))
* **deps:** update module go.uber.org/zap to v1.26.0 ([9041a1a](https://github.com/canonical/identity-platform-login-ui/commit/9041a1aae951dc4c7c35a0fe6b382bab3fbf7c93))
* do not prefix version with v ([cdc9cf1](https://github.com/canonical/identity-platform-login-ui/commit/cdc9cf15a1c05afcd858bebbc2b80c7712f9392a))
* IAM-514 compare json payloads to verify object is the same ([fde33fa](https://github.com/canonical/identity-platform-login-ui/commit/fde33fa1557d325d9e725123652d5df74ca9f936))
* move version in a separate package ([6f3c03b](https://github.com/canonical/identity-platform-login-ui/commit/6f3c03b7d0e971324db040820b74d234f766122f))

## [0.11.1](https://github.com/canonical/identity-platform-login-ui/compare/v0.11.0...v0.11.1) (2023-09-14)


### Bug Fixes

* add debug flag ([e32b5c0](https://github.com/canonical/identity-platform-login-ui/commit/e32b5c059acd5d73a12786ce7f22555b3e940863)), closes [#155](https://github.com/canonical/identity-platform-login-ui/issues/155)
* **deps:** update dependency vanilla-framework to v4.3.0 ([e19304c](https://github.com/canonical/identity-platform-login-ui/commit/e19304c6ff15d4746659a1cc08fb0c34f29e75a6))
* **deps:** update go deps ([20cc702](https://github.com/canonical/identity-platform-login-ui/commit/20cc70269222d39e444465e451e0bbc64729dd17))
* **deps:** update go deps to v1.18.0 ([6bda56c](https://github.com/canonical/identity-platform-login-ui/commit/6bda56c2e7f27f94687fe9efd9d1204028b5ff68))

## [0.11.0](https://github.com/canonical/identity-platform-login-ui/compare/v0.10.0...v0.11.0) (2023-09-12)


### Features

* added ca-certificates package to stage-packages ([e09f07c](https://github.com/canonical/identity-platform-login-ui/commit/e09f07c1d1d932d345cc6fba330b664d57d696f1))
* removed index.tsx and registration.tsx from ui/pages ([0e15096](https://github.com/canonical/identity-platform-login-ui/commit/0e15096c8f942b18223da1cb178d7484de69d344))


### Bug Fixes

* **deps:** update dependency vanilla-framework to v4 ([27079f6](https://github.com/canonical/identity-platform-login-ui/commit/27079f606565911b94df0af6b69fc55bf1c86ad0))
* **deps:** update go deps ([6658d1a](https://github.com/canonical/identity-platform-login-ui/commit/6658d1a03809876becbe16aa37b30c6423d43616))
* **deps:** update internal ui dependencies ([72f9278](https://github.com/canonical/identity-platform-login-ui/commit/72f9278974523beac9e92310bf1ea747df881e2d))
* **deps:** update module github.com/go-chi/chi/v5 to v5.0.10 ([e5042c9](https://github.com/canonical/identity-platform-login-ui/commit/e5042c9cef0e1d56694a2605e5374ed804c5993b))
* remove renovate workflow ([991215f](https://github.com/canonical/identity-platform-login-ui/commit/991215f7c19534a2b20882d47572059acff51e72))
* tranform Kratos 422 HTTP resp to 200 ([3deebc4](https://github.com/canonical/identity-platform-login-ui/commit/3deebc469366bdddf05f58295fddfa39e9ce30fc))
* Use variables instead of ints for HTTP statuses ([9332221](https://github.com/canonical/identity-platform-login-ui/commit/93322212e82caaff61c6239d7ce188d249affa57))

## [0.10.0](https://github.com/canonical/identity-platform-login-ui/compare/v0.9.0...v0.10.0) (2023-08-23)


### Features

* add healthcheck package for background tasks ([4b3fe14](https://github.com/canonical/identity-platform-login-ui/commit/4b3fe148f36cc78e07aebdb8eea3f92e64956c5f))


### Bug Fixes

* fixed bug in consent page ([d1a8e7a](https://github.com/canonical/identity-platform-login-ui/commit/d1a8e7acf407ae806cb0f16eefde6566524447a2))
* move buildInfo into service ([a8a0115](https://github.com/canonical/identity-platform-login-ui/commit/a8a01150ed04908a7aedea2aa7589b27905d4252))
* use check statuses in the service layer ([1d89f0f](https://github.com/canonical/identity-platform-login-ui/commit/1d89f0f134e4cb5e16570f198558bf9b2861e4ca))

## [0.9.0](https://github.com/canonical/identity-platform-login-ui/compare/v0.8.2...v0.9.0) (2023-08-22)


### Features

* add instrumentation to deep health check ([f565c50](https://github.com/canonical/identity-platform-login-ui/commit/f565c50f20e15e580475cce5116a0cddff6f7199))


### Bug Fixes

* add set methods for each metric ([df62e32](https://github.com/canonical/identity-platform-login-ui/commit/df62e3272964cdc952d43b1c5fc02f978eaeb02c))
* adjust wiring of status pkg ([1d69a58](https://github.com/canonical/identity-platform-login-ui/commit/1d69a58f35e2173a35623d34abeb210322624d1c))
* drop MetricInterface and adjust interface methods ([50542f6](https://github.com/canonical/identity-platform-login-ui/commit/50542f6098eb52eb7c6811251523ac7f89ac07b6))
* fixed handlers for kratos api proxying ([2ac0764](https://github.com/canonical/identity-platform-login-ui/commit/2ac0764bcccdc68acb55dd5d88187bec71473a16))

## [0.8.2](https://github.com/canonical/identity-platform-login-ui/compare/v0.8.1...v0.8.2) (2023-08-17)


### Bug Fixes

* use the same import and drop an alias ([5af7944](https://github.com/canonical/identity-platform-login-ui/commit/5af79443803655e50e6bb392d2c8cceaf5688110))

## [0.8.1](https://github.com/canonical/identity-platform-login-ui/compare/v0.8.0...v0.8.1) (2023-08-15)


### Bug Fixes

* fixed timeout in deepcheck handler ([511f86f](https://github.com/canonical/identity-platform-login-ui/commit/511f86feadf3cde8898120ba1c47f98d46e400e1))

## [0.8.0](https://github.com/canonical/identity-platform-login-ui/compare/v0.7.0...v0.8.0) (2023-08-15)


### Features

* added new /api/v0/deepcheck route ([7ccc8eb](https://github.com/canonical/identity-platform-login-ui/commit/7ccc8eb41391fb87d96d40a09159d5986be1365f))
* added service layer to pkg/status ([aee2027](https://github.com/canonical/identity-platform-login-ui/commit/aee202779dc417a405013b2a7f50dd50a24e8a8b))
* adjust the wiring of the status pkg ([2b6b66c](https://github.com/canonical/identity-platform-login-ui/commit/2b6b66c92a83e23311bec21d0bad3808f4ae482a))


### Bug Fixes

* add MetadataAPI methods to clients objects ([c5e76e1](https://github.com/canonical/identity-platform-login-ui/commit/c5e76e1de728c38c2e93c958198fc670e8fc2425))

## [0.7.0](https://github.com/canonical/identity-platform-login-ui/compare/v0.6.3...v0.7.0) (2023-08-15)


### Features

* changed the route to the ui resources to /ui/* ([bff64e5](https://github.com/canonical/identity-platform-login-ui/commit/bff64e541ded6766c9feae668fc72468f326ec97))


### Bug Fixes

* fixed server treating all misc routes as fileserver requests ([b8b4ded](https://github.com/canonical/identity-platform-login-ui/commit/b8b4dedb5d69d50bab645e8aed34ff8b5ee92aa9))

## [0.6.3](https://github.com/canonical/identity-platform-login-ui/compare/v0.6.2...v0.6.3) (2023-07-27)


### Bug Fixes

* use otelhttp transport to propagate traces ([fe14b59](https://github.com/canonical/identity-platform-login-ui/commit/fe14b5970bbef93ec6809f9626912b0ad8f25194))

## [0.6.2](https://github.com/canonical/identity-platform-login-ui/compare/v0.6.1...v0.6.2) (2023-07-26)


### Bug Fixes

* add jaeger propagator as ory components support only these spans for now ([0b5f248](https://github.com/canonical/identity-platform-login-ui/commit/0b5f2483020b83cc69bea1fbdf6788b601c0005f))
* add otel grpc+http endpoint for tracing ([e1b1424](https://github.com/canonical/identity-platform-login-ui/commit/e1b14247c3ab02a44f3ca09b243bfd34c747c2c0))
* pass new context to clients to propagate trace ids ([7dfdf05](https://github.com/canonical/identity-platform-login-ui/commit/7dfdf0503320c9056b84ebe4763682097f213692))
* wire up new config needed for otel grpc+http ([17399ea](https://github.com/canonical/identity-platform-login-ui/commit/17399ea9219a06d522f213658bb8f3c88135ad32))

## [0.6.1](https://github.com/canonical/identity-platform-login-ui/compare/v0.6.0...v0.6.1) (2023-07-11)


### Bug Fixes

* Change response status to 200 ([#94](https://github.com/canonical/identity-platform-login-ui/issues/94)) ([89af73b](https://github.com/canonical/identity-platform-login-ui/commit/89af73baf5cfe384b48ce05f64fb435d36bdd3a0))
* Copy only cookies in proxied requests ([#95](https://github.com/canonical/identity-platform-login-ui/issues/95)) ([2d1fa6a](https://github.com/canonical/identity-platform-login-ui/commit/2d1fa6aad9d47127e6c914cc5389fc94beb35a1c))

## [0.6.0](https://github.com/canonical/identity-platform-login-ui/compare/v0.4.1...v0.6.0) (2023-07-07)


### Bug Fixes

* add cors middleware ([9d5cb04](https://github.com/canonical/identity-platform-login-ui/commit/9d5cb0412a31c77a377ee349219ccc5e5c9c3b91))
* Move logic out of misc package ([8ae6da5](https://github.com/canonical/identity-platform-login-ui/commit/8ae6da52532661af33ea86e06bc6a3a594c8f22a))


## [0.4.1](https://github.com/canonical/identity-platform-login-ui/compare/v0.4.0...v0.4.1) (2023-06-28)


### Bug Fixes

* add tracing enabling variable, defaulting to true ([f0858e3](https://github.com/canonical/identity-platform-login-ui/commit/f0858e3d98083e2e8b92ab4dc1f74e039976bb32))
* IAM-353 - allow tracing to be disabled ([1ff3186](https://github.com/canonical/identity-platform-login-ui/commit/1ff3186a5105a0b3c1fbdf854be53963a88e8a95))

## [0.4.0](https://github.com/canonical/identity-platform-login-ui/compare/v0.3.0...v0.4.0) (2023-06-27)


### Features

* IAM-326 - separate logic and http handling in extra package ([7f9549c](https://github.com/canonical/identity-platform-login-ui/commit/7f9549cab52c87c3f877a3206f6193cacad336ff))
* IAM-330 - introduce otel tracing ([198fb16](https://github.com/canonical/identity-platform-login-ui/commit/198fb16b0ed223433f297dae25e5012c53aece84))
* IAM-330 - use otelhttp middleware ([5efc2a0](https://github.com/canonical/identity-platform-login-ui/commit/5efc2a0dda0b6ffd7f8460826cd4ff8b98469816))
* wire up service in extra package ([9865de8](https://github.com/canonical/identity-platform-login-ui/commit/9865de8ce3e95ae752aaa1ea4b141e1f5d490361))
* wire up tracer inside status endpoint as a dummy example ([b51e91f](https://github.com/canonical/identity-platform-login-ui/commit/b51e91f8b6d1ea7922c91541689af7933da45215))

## [0.3.0](https://github.com/canonical/identity-platform-login-ui/compare/v0.2.0...v0.3.0) (2023-06-27)


### Features

* add log rotator for zap ([dc76368](https://github.com/canonical/identity-platform-login-ui/commit/dc763681a85c492121e163e1526b901a75ca3849))
* IAM-327 - introduce zap ([9e378e8](https://github.com/canonical/identity-platform-login-ui/commit/9e378e8f2cac263a54f67c16cd63eff5ecb2f0e5))
* IAM-327 - wire up zap.SugaredLogger ([6e22236](https://github.com/canonical/identity-platform-login-ui/commit/6e222368d8e58c184d4ca319200aed5ae1c685bb))
* IAM-328 - use chi http logger middleware ([79283d8](https://github.com/canonical/identity-platform-login-ui/commit/79283d8ad51cabcf2e716c8116af2d95670cbfa1))
* IAM-328 - use go-chi for router ([55c873d](https://github.com/canonical/identity-platform-login-ui/commit/55c873d08b633a20cba76d2b063e2f71de8f9125))
* introduce config management via envconfig ([55367ee](https://github.com/canonical/identity-platform-login-ui/commit/55367ee0d2694fe3634f1a0dd79270562c857b66))
* wire up new web pkg for routing ([0d67eda](https://github.com/canonical/identity-platform-login-ui/commit/0d67eda2569ce8637663cb6ceeacd1b16013eaed))


### Bug Fixes

* add * to UI route ([0c7d1d0](https://github.com/canonical/identity-platform-login-ui/commit/0c7d1d05b24c525241f9c196cd181150c5d9ef56))
* create prometheus middleware for chi ([1678f56](https://github.com/canonical/identity-platform-login-ui/commit/1678f56033ba6dc5d28f871d7158bb6abca32072))
* rename prometheus pkg to metrics and change endpoints ([ca99c3c](https://github.com/canonical/identity-platform-login-ui/commit/ca99c3c9c72132c338223df60387c3244c4063c3))

## [0.2.0](https://github.com/canonical/identity-platform-login-ui/compare/v0.1.0...v0.2.0) (2023-06-22)


### Features

* add git commit sha on health endpoint ([a09afce](https://github.com/canonical/identity-platform-login-ui/commit/a09afce63dedf08b6e38a348e9fdca05c0ceaaba))
* create an oidc pkg to be shared ([91e199f](https://github.com/canonical/identity-platform-login-ui/commit/91e199f760d83ba5899108d9d2a8c0473431990a))
* create hydra and kratos pkgs, useful for mocking ([e585da8](https://github.com/canonical/identity-platform-login-ui/commit/e585da8ba89f9f21d020446cbb84453d7e1c4981))
* move conset api in a separate pkgs ([895d392](https://github.com/canonical/identity-platform-login-ui/commit/895d392b6d1ec0b2d5a9d7b714457cce43579cff))
* move kratos api set in a separate pkgs ([0f4e552](https://github.com/canonical/identity-platform-login-ui/commit/0f4e552c05fde1eff5cd752d448a5bb17bf20e22))
* move ui api in a separate pkg ([7142e08](https://github.com/canonical/identity-platform-login-ui/commit/7142e08dfd8e26f86c643175e6083a8b67252c79))
* offload logic to packages ([3226beb](https://github.com/canonical/identity-platform-login-ui/commit/3226bebc9a84ceb51d8d3e5b5a4e08b5b277802c))


### Bug Fixes

* add graceful server shutdown ([d67989f](https://github.com/canonical/identity-platform-login-ui/commit/d67989f277409d9b1e93068be7e5b83c9e5933be))
* move helpers to a shared internal package, plan to remove ([c160874](https://github.com/canonical/identity-platform-login-ui/commit/c1608749c35bd1367cebd59cc5662b7121afb694))
* move logging pkg into internal ([23b32d4](https://github.com/canonical/identity-platform-login-ui/commit/23b32d4bbde66e9d2183ca3d05c95fa7c2e7580d))
* move tests into relative packages, rename health to status ([29d58e1](https://github.com/canonical/identity-platform-login-ui/commit/29d58e10a605ed6a12cefa686281239b001da092))
* use full repo url for go module name ([3c8bc22](https://github.com/canonical/identity-platform-login-ui/commit/3c8bc223db7dde172571d8a0879371394f83f4b8))
* wire up new pkgs handlers ([ed7e9be](https://github.com/canonical/identity-platform-login-ui/commit/ed7e9becc3ac981ded352b531c72ff9d95bcfdca))
