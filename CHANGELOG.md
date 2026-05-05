# Changelog

## 0.1.1 (2026-05-05)

Full Changelog: [v0.1.0...v0.1.1](https://github.com/lumalabs/luma-agents-cli/compare/v0.1.0...v0.1.1)

### Chores

* update SDK settings ([96d539b](https://github.com/lumalabs/luma-agents-cli/commit/96d539bfd7bdff40a1db3d46ea5f2da68d5d9201))
* update SDK settings ([1923f49](https://github.com/lumalabs/luma-agents-cli/commit/1923f49ef118129a9a04a47c6d5bbc4e5616f9d2))


### Documentation

* update API documentation URL in README ([9535717](https://github.com/lumalabs/luma-agents-cli/commit/9535717afa9e2e3b6a98e1fe3969c84828c29f61))

## 0.1.0 (2026-05-04)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/lumalabs/luma-agents-cli/compare/v0.0.1...v0.1.0)

### Features

* **api:** add user_id parameter to generations create ([90d828d](https://github.com/lumalabs/luma-agents-cli/commit/90d828d88a28a91350fa1935cce55859a369bf49))
* **api:** manual updates ([a1a9c39](https://github.com/lumalabs/luma-agents-cli/commit/a1a9c39369ed30636a23871e9b720c45c18e48fb))
* **api:** manual updates ([4f23255](https://github.com/lumalabs/luma-agents-cli/commit/4f232559ecc91c2e34c4022f5c1c53b2f26b50b4))
* **api:** manual updates ([a7197bc](https://github.com/lumalabs/luma-agents-cli/commit/a7197bc6de2e08e1500be35088dd309b1a72b6c2))
* **api:** manual updates ([51538cf](https://github.com/lumalabs/luma-agents-cli/commit/51538cf43c26afa95a26fc02648a643521f8dc89))
* **cli:** add `--raw-output`/`-r` option to print raw (non-JSON) strings ([8c740c5](https://github.com/lumalabs/luma-agents-cli/commit/8c740c50398569aa81839263bc771e5b59f8ba8f))
* **cli:** alias parameters in data with `x-stainless-cli-data-alias` ([7216fbb](https://github.com/lumalabs/luma-agents-cli/commit/7216fbbe937a88b5e0c4806c02eabe76aef9e791))
* **cli:** send filename and content type when reading input from files ([f6c2587](https://github.com/lumalabs/luma-agents-cli/commit/f6c258784c765dfb28e840341f1a0cf7bc2fcdd6))
* support passing path and query params over stdin ([f06b836](https://github.com/lumalabs/luma-agents-cli/commit/f06b83678c1a33113110bca972186e14104e3f8b))


### Bug Fixes

* **api:** remove default value from model parameter in generations create ([75064aa](https://github.com/lumalabs/luma-agents-cli/commit/75064aa75cffa033f9b90c4fc2e82594c9991bf9))
* **cli:** correctly load zsh autocompletion ([ace56fe](https://github.com/lumalabs/luma-agents-cli/commit/ace56fe10fa691b3a04172d9d9c8ca661bb92f5e))
* **cli:** fix incompatible Go types for flag generated as array of maps ([f415a44](https://github.com/lumalabs/luma-agents-cli/commit/f415a44a3f952ca9da5b0a034da1efbfcde4dd69))
* fix for failing to drop invalid module replace in link script ([c3946ff](https://github.com/lumalabs/luma-agents-cli/commit/c3946ffe88ff93e053d81ce91985134c54ee6095))
* flags for nullable body scalar fields are strictly typed ([0472624](https://github.com/lumalabs/luma-agents-cli/commit/0472624a94ae7de2b1038216780d47b74ddf07d3))


### Chores

* add documentation for ./scripts/link ([0658309](https://github.com/lumalabs/luma-agents-cli/commit/065830944f43fb5a314bef3582c63b7f865c7b9d))
* **ci:** support manually triggering release workflow ([36bc04b](https://github.com/lumalabs/luma-agents-cli/commit/36bc04b6a827ccb860ec0f925391824dc53690e3))
* **cli:** additional test cases for `ShowJSONIterator` ([79054c0](https://github.com/lumalabs/luma-agents-cli/commit/79054c0bedf7501d5f6b10e5613e90cbb59b401a))
* **cli:** fall back to JSON when using default "explore" with non-TTY ([014d0c7](https://github.com/lumalabs/luma-agents-cli/commit/014d0c7c3e7f0f1c4d66f8d04728f5f41946b88e))
* **cli:** let `--format raw` be used in conjunction with `--transform` ([4654615](https://github.com/lumalabs/luma-agents-cli/commit/4654615a811ba7321cd290c57a69b225e48d59cb))
* **cli:** switch long lists of positional args over to param structs ([323ea97](https://github.com/lumalabs/luma-agents-cli/commit/323ea9707c7feb3691ff659b69b85c46392f530f))
* **cli:** use `ShowJSONOpts` as argument to `formatJSON` instead of many positionals ([a460710](https://github.com/lumalabs/luma-agents-cli/commit/a460710765b204ee8f17d2e629668d13e568ed7c))
* configure new SDK language ([4ad4c0f](https://github.com/lumalabs/luma-agents-cli/commit/4ad4c0fce4fb2d6190c30a964d8f98db1d9c1345))
* **internal:** more robust bootstrap script ([be0a924](https://github.com/lumalabs/luma-agents-cli/commit/be0a924e4e90994210e01f9114cc7fdbf4776e6f))
* **tests:** bump steady to v0.22.1 ([f20417d](https://github.com/lumalabs/luma-agents-cli/commit/f20417d8a30d005a7bd7b8366130b559840411ca))
* update SDK settings ([878a229](https://github.com/lumalabs/luma-agents-cli/commit/878a22904535906bc1e5e0b4f7657e570493c31e))
* update SDK settings ([2e66f1b](https://github.com/lumalabs/luma-agents-cli/commit/2e66f1b660280d2c3e255c502dc6595722620bf8))


### Documentation

* **api:** clarify image-ref parameter limits in generations create ([275e391](https://github.com/lumalabs/luma-agents-cli/commit/275e39178003727ccf9d43459dc4cf1076b1b7c3))
