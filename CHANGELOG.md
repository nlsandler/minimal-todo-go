# Changelog

## 0.1.0-alpha.3 (2026-02-05)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/nlsandler/minimal-todo-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** add nullable delete response ([63715c7](https://github.com/nlsandler/minimal-todo-go/commit/63715c7b2b23c069fd60c5e1d979334a2945e9ec))
* **api:** additional properties in nested  request body... ([90002a7](https://github.com/nlsandler/minimal-todo-go/commit/90002a7e0096f1864396dda743e14726e9c42225))
* **api:** additionalProperties response ([7738773](https://github.com/nlsandler/minimal-todo-go/commit/7738773f5d033340893be56d1ed68441c431942e))
* **api:** get rid of irrational readOnly ([c90ca13](https://github.com/nlsandler/minimal-todo-go/commit/c90ca13dd3090a2cca5511cc9cd81ef568e77160))
* **api:** make a model shared ([a0a467d](https://github.com/nlsandler/minimal-todo-go/commit/a0a467da5718f3d5138b3d7b2d24f2925f84a449))
* **api:** make note non-nullable ([183fc90](https://github.com/nlsandler/minimal-todo-go/commit/183fc9014eb40af6a58fe4c2a02028d369a63525))
* **api:** make note required ([dbde10c](https://github.com/nlsandler/minimal-todo-go/commit/dbde10ceaa4ca502da632b86250f685a446f5083))
* **api:** make tag label nullable ([45381c7](https://github.com/nlsandler/minimal-todo-go/commit/45381c7e3a482a779139f88087390e559030df3a))
* **api:** make task name optional ([71c16ec](https://github.com/nlsandler/minimal-todo-go/commit/71c16ecc5906ee88571dc8df7bf0140bf131d598))
* **api:** re-use NESTED schema instead of top-level one in reqs/responses ([dbca9cd](https://github.com/nlsandler/minimal-todo-go/commit/dbca9cdb2cd7e4a8f66470ecfe0b173bbf097e31))
* **api:** readonly - fix typo ([69d43c8](https://github.com/nlsandler/minimal-todo-go/commit/69d43c85414751d3b0efc9e29b41aeebdabcee39))
* **api:** reuse schema in request take 2 ([482cdf0](https://github.com/nlsandler/minimal-todo-go/commit/482cdf08c2f2aa1cd1c709d9094abc3ee7e17842))
* **api:** string maxLength ([e647b54](https://github.com/nlsandler/minimal-todo-go/commit/e647b54ecfcde6fe2b487cc3799a7e2d4664c67c))
* **api:** use schema in request body take 3... ([849f919](https://github.com/nlsandler/minimal-todo-go/commit/849f91918e49b5d2210094c893abc31b557d0232))
* **api:** whoops ([3948ceb](https://github.com/nlsandler/minimal-todo-go/commit/3948cebe75b7fc1b9f09588a7f264e20bd9889a2))

## 0.1.0-alpha.2 (2026-02-03)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/nlsandler/minimal-todo-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** remove basic auth ([93e69b1](https://github.com/nlsandler/minimal-todo-go/commit/93e69b1b7976cf67ae6d2e1abef0b53d9abab18d))
* **codeflow:** enable breaking changes detection ([f282106](https://github.com/nlsandler/minimal-todo-go/commit/f28210615612448ca829641ae637d0cdfc96a2fa))

## 0.1.0-alpha.1 (2026-02-03)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/nlsandler/minimal-todo-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### ⚠ BREAKING CHANGES

* **client:** rename resp package
* **client:** improve core function names

### Features

* **api:** update via SDK Studio ([83c2d02](https://github.com/nlsandler/minimal-todo-go/commit/83c2d02122d1215415330299ffe2594622daff38))
* **api:** update via SDK Studio ([1715b05](https://github.com/nlsandler/minimal-todo-go/commit/1715b05f36fe2111232856667590521bae47a105))
* **api:** update via SDK Studio ([c4679fb](https://github.com/nlsandler/minimal-todo-go/commit/c4679fb6d12bd0e5bc2c2afc70e94a3511777697))
* **api:** update via SDK Studio ([27d3be0](https://github.com/nlsandler/minimal-todo-go/commit/27d3be021c762d0b30ec15c38a34df9d9d775d7c))
* **api:** update via SDK Studio ([a387e8b](https://github.com/nlsandler/minimal-todo-go/commit/a387e8ba207389298739f0fde0a30243bd705236))
* **client:** add debug log helper ([87e52eb](https://github.com/nlsandler/minimal-todo-go/commit/87e52eba7ea2c589be8940c27024f82037caae42))
* **client:** add escape hatch for null slice & maps ([822dd4f](https://github.com/nlsandler/minimal-todo-go/commit/822dd4fcbc015f6b3a46dee0bfeab57c912360eb))
* **client:** add support for endpoint-specific base URLs in python ([3607306](https://github.com/nlsandler/minimal-todo-go/commit/36073060a99a91f75925ed05885cd15ce6c25924))
* **client:** allow overriding unions ([f1dcd74](https://github.com/nlsandler/minimal-todo-go/commit/f1dcd747099b62578552b9b15518a9853bb758bd))
* **client:** experimental support for unmarshalling into param structs ([e120230](https://github.com/nlsandler/minimal-todo-go/commit/e12023086b647fbfd86aa7f3dee4b4090d1637a3))
* **client:** rename resp package ([4743615](https://github.com/nlsandler/minimal-todo-go/commit/47436150da3d4951f567e02bf6c87dd095f95755))


### Bug Fixes

* **client:** clean up reader resources ([a14f5ed](https://github.com/nlsandler/minimal-todo-go/commit/a14f5ed89b2296cd11649a33ed4dca7b20b62e45))
* **client:** correctly set stream key for multipart ([80047e0](https://github.com/nlsandler/minimal-todo-go/commit/80047e09cfb55cba57b6c4932800cebc939ff067))
* **client:** correctly update body in WithJSONSet ([d2771c2](https://github.com/nlsandler/minimal-todo-go/commit/d2771c2e879a0d91edc52c24ce4b510421e9a126))
* **client:** don't panic on marshal with extra null field ([e773040](https://github.com/nlsandler/minimal-todo-go/commit/e77304059ae7bcca8439aed18daba04bea2e290e))
* **client:** improve core function names ([b6c9d42](https://github.com/nlsandler/minimal-todo-go/commit/b6c9d4266703af301aa5bfd3879264cc978841c5))
* **client:** improve docs ([6871b72](https://github.com/nlsandler/minimal-todo-go/commit/6871b725d32c0a9e9b763f78083899a03dff8fc6))
* **client:** resolve issue with optional multipart files ([ff02592](https://github.com/nlsandler/minimal-todo-go/commit/ff02592ff8583ccd9ac53322b0e278705605c358))
* **client:** unmarshal responses properly ([534c71c](https://github.com/nlsandler/minimal-todo-go/commit/534c71c051c902c2067bcf9a9dbe4416f75ca375))
* don't try to deserialize as json when ResponseBodyInto is []byte ([228ac3b](https://github.com/nlsandler/minimal-todo-go/commit/228ac3b8e646d17f0b73447be43e5e67b850764d))
* fix error ([15b5730](https://github.com/nlsandler/minimal-todo-go/commit/15b5730b7d1fb2534b3fccb94e305ecef7218ae1))
* handle empty bodies in WithJSONSet ([c15f82e](https://github.com/nlsandler/minimal-todo-go/commit/c15f82ea2672a5803220f57ddb2e190691316760))


### Chores

* **ci:** enable for pull requests ([cc26751](https://github.com/nlsandler/minimal-todo-go/commit/cc2675141ec7c131763dc9cff2350702d6f2a712))
* **ci:** only run for pushes and fork pull requests ([19abdbd](https://github.com/nlsandler/minimal-todo-go/commit/19abdbd5f41d6d283f1001cded78c1ab1700adfe))
* **ci:** only use depot for staging repos ([b4fb4d4](https://github.com/nlsandler/minimal-todo-go/commit/b4fb4d467f89d213ef1776727e7c16f2fc407a16))
* configure new SDK language ([e323b97](https://github.com/nlsandler/minimal-todo-go/commit/e323b97819f9ae02d17fa1e1a150da0fe64873a8))
* **docs:** grammar improvements ([2cf2a4e](https://github.com/nlsandler/minimal-todo-go/commit/2cf2a4e1b233e97716ba3070095c04240d2d6042))
* **docs:** update respjson package name ([bc27ad1](https://github.com/nlsandler/minimal-todo-go/commit/bc27ad1edf2c7dbb53f1c5e6ad9f136891eecafb))
* fix documentation of null map ([a8aca65](https://github.com/nlsandler/minimal-todo-go/commit/a8aca65a3e5e07de57e24473d627664c35f83c58))
* improve devcontainer setup ([806604b](https://github.com/nlsandler/minimal-todo-go/commit/806604b9c932768080b97b26d37213f535e062d1))
* **internal:** codegen related update ([ff2cb32](https://github.com/nlsandler/minimal-todo-go/commit/ff2cb323addd27b05b64dcf7e8a83605ca885b51))
* make go mod tidy continue on error ([3fafbc5](https://github.com/nlsandler/minimal-todo-go/commit/3fafbc5e426a92bf3d0866f9ae2a475441a0ea7c))
* update SDK settings ([b7d9506](https://github.com/nlsandler/minimal-todo-go/commit/b7d950663f8030fffdb2c34084c2aa48f67df7f9))


### Documentation

* remove or fix invalid readme examples ([2b0a1b5](https://github.com/nlsandler/minimal-todo-go/commit/2b0a1b5df85c700ba860bb16ee3d246dd4380073))
