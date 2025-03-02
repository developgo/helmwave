# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html),
and is generated by [Changie](https://github.com/miniscruff/changie).

# v0.24.4

### New feature

* [#466](https://github.com/helmwave/helmwave/issues/466) Add experimental support for showing 3-way merge diff (@r3nic1e)

### Fixed

* [#518](https://github.com/helmwave/helmwave/issues/518) Added missing `tmpl` and `tpl` functions to gomplate renderer (@r3nic1e)




# v0.24.3

### Fixed

* [#515](https://github.com/helmwave/helmwave/issues/515) Make diff more consistent with semantical changes (@r3nic1e)

### Dependencies

* [#515](https://github.com/helmwave/helmwave/issues/515) Upgrade helm-diff to latest version (@r3nic1e)

* Upgrade helm to 3.10.2 (@r3nic1e)




# v0.24.2

### Fixed

* [#502](https://github.com/helmwave/helmwave/issues/502) Reverted to prior YAML parser due to numerous parsing issues (@r3nic1e)

* [#502](https://github.com/helmwave/helmwave/issues/502) Fix jsonschema to allow additional properties at root level and incorrect version pattern (@r3nic1e)

### CI

* Fix telegram template (@r3nic1e)




# v0.24.1

### Fixed

* Reverted strict helmwave.yml parsing to support anchors (@r3nic1e)

### Other

* Update minor dependencies (@r3nic1e)




# v0.24.0

### **Breaking change!**

* [#487](https://github.com/helmwave/helmwave/issues/487) YAML parsing helmwave.yml is now more strict: duplicate fields and unknown fields in map are now prohibited and return parsing error (@r3nic1e)

### New feature

* Create human-readable changelog (@r3nic1e)

* [#441](https://github.com/helmwave/helmwave/issues/441) Add support for postrenderer (@r3nic1e)

* [#453](https://github.com/helmwave/helmwave/issues/453) Allow to set `delimiter_left` and `delimiter_right` for each values file (@r3nic1e)

* [#462](https://github.com/helmwave/helmwave/issues/462) Allow to set kubecontext for each release via `context` field. Kubedog requires only one context in plan used. (@r3nic1e)

* [#422](https://github.com/helmwave/helmwave/issues/422) Store remote charts in plandir and use them for apply (@r3nic1e)

* [#482](https://github.com/helmwave/helmwave/issues/482) Allow to skip nonexisting dependencies via `optional` flag (@r3nic1e)

* [#484](https://github.com/helmwave/helmwave/issues/484) Allow to use tag as dependency. All releases with this tag will be added as dependencies (@r3nic1e)

* [#488](https://github.com/helmwave/helmwave/issues/488) Add descriptions, enum, defaults, patterns to jsonschema (@r3nic1e)

### Fixed

* Fix `client rate limiter Wait returned an error: context deadline exceeded` with setting default release timeout to 5 minutes (@r3nic1e)

* [#447](https://github.com/helmwave/helmwave/issues/447) Disable client-only for dry-run to handle `.Capabilities.APIVersions` correctly (@r3nic1e)

* [#459](https://github.com/helmwave/helmwave/issues/459) Do not fail if plan doesn't contain any releases (@r3nic1e)

* [#488](https://github.com/helmwave/helmwave/issues/488) Tune jsonschema to better generate `required` fields (@r3nic1e)

### CI

* Simplify release workflow (@r3nic1e)

* Use changie to generate changelog (@r3nic1e)

* Force all release PRs to have changelog (@r3nic1e)

* Fix generating schema.json (@r3nic1e)

### Dependencies

* Update helm to 3.10.1 (@r3nic1e)

* Update other dependencies (@r3nic1e)

### Other

* [#487](https://github.com/helmwave/helmwave/issues/487) Switch to github.com/goccy/go-yaml YAML library (@r3nic1e)




# v0.23.1

### **Breaking change!**

* Set kubedog `AllowFailuresCount` to 0 by default (@r3nic1e)

### Fixed

* Fix deadlock when dependency is failed (@r3nic1e)



# v0.23.0

### New feature

* Generate json schema for helmwave.yml (@zhilyaev)



# v0.22.2

### Fixed

* Increase kubernetes client ratelimit (@r3nic1e)

### Dependencies

* Update helm to 3.9.4 (@r3nic1e)



# v0.22.1

### New feature

* [#417](https://github.com/helmwave/helmwave/issues/417) Show diff for charts metadata (@r3nic1e)

### Fixed

* [#425](https://github.com/helmwave/helmwave/issues/425) Add warn if build flags are unused (@r3nic1e)



# v0.22.0

### Fixed

* [#413](https://github.com/helmwave/helmwave/issues/413) Warn if kubedog is enabled, but `wait` helm option is disabled (@r3nic1e)

* [#415](https://github.com/helmwave/helmwave/issues/415) Generate full `depends_on` in planfile (@r3nic1e)



# v0.21.1

### New feature

* [#405](https://github.com/helmwave/helmwave/issues/405) Make namespace in `depends_on` optional (@r3nic1e)

### Fixed

* [#410](https://github.com/helmwave/helmwave/issues/410) Fix logging in public OCI registry (@zhilyaev)

### CI

* Add goreleaser testing config (@r3nic1e)



# v0.21.0

### New feature

* [#358](https://github.com/helmwave/helmwave/issues/358) Add `pending_release_strategy` option to handle pending releases (@r3nic1e)

* Build apk, deb and rpm packages (@r3nic1e)

* Build arm64 packages (@r3nic1e)

* Make docker images multiarch (@r3nic1e)

### Fix

* [#391](https://github.com/helmwave/helmwave/issues/391) Make existing repository error more user-friendly (@r3nic1e)

### Other

* Refactor kubedog (@r3nic1e)



# v0.20.3

### Fixed

* Stop kubedog after release is installed (@r3nic1e)



# v0.20.2

### Fixed

* Added validation for values during plan (@zhilyaev)



# v0.20.1

### CI

* Fix workflows permissions (@r3nic1e)



# v0.20.0

### New feature

* [#354](https://github.com/helmwave/helmwave/issues/354) Use new dependency graph implementation (@r3nic1e)

* Allow to use short syntax to reference chart (@zhilyaev)

* [#354](https://github.com/helmwave/helmwave/issues/354) Add support for limited parallel releases (@r3nic1e)

* [#377](https://github.com/helmwave/helmwave/issues/377) Add support for ignoring nonexisting values via `strict` option (@r3nic1e)

* [#378](https://github.com/helmwave/helmwave/issues/378) Add support for skipping templating values via `render` option (@r3nic1e)

### Fixed

* Silence some klog messages (@zhilyaev)

### Other

* Add some generics (@r3nic1e)



# v0.19.7

### Dependencies

* Update helm to 3.9.2 (@zhilyaev)



# v0.19.6

### Dependencies

* Update helm to 3.9.1 (@zhilyaev)



# v0.19.5

### Dependencies

* Update helm to 3.9.0 (@zhilyaev)



# v0.19.4

### Fixed

* Revert updating kubernetes client to 0.24.0 (@r-mironov)

* [#330](https://github.com/helmwave/helmwave/issues/330) Fix ignoring `${HELM_KUBECONTEXT}` (@zhilyaev)

* [#329](https://github.com/helmwave/helmwave/issues/329) Add missing custom functions for gomplate templater (@r-mironov)



# v0.19.3

### Fixed

* [#116](https://github.com/helmwave/helmwave/issues/116) Fail if kubedog is failed (@zhilyaev)

### Dependencies

* Update helm to 3.8.2 (@zhilyaev)

* Update minor dependencies (@zhilyaev)

### Other

* [#309](https://github.com/helmwave/helmwave/issues/309) Set `${ASSUME_NO_MOVING_GC_UNSAFE_RISK_IT_WITH}` (@zhilyaev)



# v0.19.2

### New feature

* Drop some kubernetes logs (@zhilyaev)

### Dependencies

* Update golang to 1.18 (@zhilyaev)

* Update base alpine image to 3.15 (@zhilyaev)



# v0.19.0

### New feature

* [#115](https://github.com/helmwave/helmwave/issues/115) Add support for OCI helm registries (@zhilyaev)

### Dependencies

* Update golangci-lint to 1.45.0 (@zhilyaev)

* Update helm to 3.8.1 (@zhilyaev)



# v0.18.0

### **Breaking change!**

* [#264](https://github.com/helmwave/helmwave/issues/264) Helm options are now named with snake_case (@r-mironov)

### New feature

* [#257](https://github.com/helmwave/helmwave/issues/257) Add annotation to skip resource in diff (@r3nic1e)

* [#257](https://github.com/helmwave/helmwave/issues/257) Hide helm hooks in diff (@r3nic1e)

### Fixed

* Disable helm diff colors if colors are disabled (@zhilyaev)

* Disable emojis when colors are disabled (@zhilyaev)

### Dependencies

* Update minor dependencies (@r3nic1e)

* Update golangci-lint to 1.44 (@r3nic1e)



# v0.17.2

### Fixed

* [#256](https://github.com/helmwave/helmwave/issues/256) Fix upgrade diff (@r3nic1e)



# v0.17.1

### Fixed

* Fix nil pointer exception in logger (@r3nic1e)



# v0.17.0

### **Breaking change!**

* [#218](https://github.com/helmwave/helmwave/issues/218) Gomplate datasources via config are no longer supported (@r3nic1e)

### New feature

* [#218](https://github.com/helmwave/helmwave/issues/218) Allow to select templater with `--templater` flag (@r3nic1e)

* Add more fields to lots of logs (@r3nic1e)

* [#248](https://github.com/helmwave/helmwave/issues/248) Add support for log timestamps via `log-timestamps` flag (@r3nic1e)

* [#229](https://github.com/helmwave/helmwave/issues/229) Move helm progress info to INFO loglevel via `progress` flag (@r-mironov)

### Fixed

* [#236](https://github.com/helmwave/helmwave/issues/236) Make kubedog respect `${KUBECONFIG}` (@r-mironov)

### CI

* Use trivy to analyze images (@zhilyaev)

* Add gifs to closed PRs (@zhilyaev)

### Other

* Fix lots of linter issues (@r3nic1e)

* Add more tests (@r3nic1e)



# v0.16.7

### New feature

* Add support for `wait-for-jobs` helm option (@r3nic1e)

### Fixed

* Fix incorrect `reset-values` helm option for upgrade (@r3nic1e)

### Dependencies

* Update gopkg.in/yaml to v3 (@r3nic1e)



# v0.16.6

### Dependencies

* [#211](https://github.com/helmwave/helmwave/issues/211) Update minor dependencies (@r3nic1e)

### Other

* Refactor to use more interfaces (@r3nic1e)

* Add some tests with mocks (@r3nic1e)

* [#198](https://github.com/helmwave/helmwave/issues/198) Add tests for local plan diff (@r3nic1e)

* Enable more linters and fix linter issues (@r3nic1e)



# v0.16.5

### Fixed

* [#182](https://github.com/helmwave/helmwave/issues/182) Fix concurrent map write in plan (@zhilyaev)

### CI

* Update goreleaser config (`buildx`) (@r-mironov)



# v0.16.3

### Fixed

* Add missing dependencies to plan (@r3nic1e)



# v0.16.2

### New feature

* Add validation for duplicate releases (@zhilyaev)

### Dependencies

* Update golangci-lint to 1.43.0 (@zhilyaev)

### Other

* Fix some linter errors (@zhilyaev)



# v0.16.1

### Fixed

* Validate plan before using it (@zhilyaev)



# v0.16.0

### New feature

* Generate helm diff with live release (@zhilyaev)

### Fixed

* Fix some kubedog errors (@zhilyaev)



# v0.15.1

### CI

* Add building scratch docker images (@zhilyaev)



# v0.15.0

### New feature

* [#168](https://github.com/helmwave/helmwave/issues/168) Add `allow_failure` to skip failed dependencies (@r3nic1e)

### CI

* Update golangci-lint to 1.42.1 (@r3nic1e)



# v0.14.2

### Dependencies

* [#166](https://github.com/helmwave/helmwave/issues/166) Update helm to 3.7.1 (@zhilyaev)

* Update minor dependencies (@zhilyaev)



# v0.14.1

### New feature

* Add support for setting gomplate datasources (@r3nic1e)

### Other

* Add new tests (@r3nic1e)



# v0.14.0

### New feature

* [#103](https://github.com/helmwave/helmwave/issues/103) [#159](https://github.com/helmwave/helmwave/issues/159) Add support for gomplate (@r3nic1e)



# v0.13.0

### New feature

* Add `--yml` flag to `build` command to automatically run `yml` command (@zhilyaev)

* Add `--build` flag to `up` command to automatically run `build` command (@zhilyaev)

### Dependencies

* Update helm to 3.7.0 (@r-mironov)

* Update minor dependencies (@zhilyaev)



# v0.12.8

### Fixed

* [#150](https://github.com/helmwave/helmwave/issues/150) Fix checking if release is installed (@zhilyaev)



# v0.12.7

### Fixed

* [#146](https://github.com/helmwave/helmwave/issues/146) Allow to render manifests when kubernetes is unaccessible (@zhilyaev)



# v0.12.6

### Fixed

* Handle different filesystems for plandir and tmpdir (@zhilyaev)



# v0.12.5

### Other

* Refactor some errors (@zhilyaev)



# v0.12.4

### New feature

* Speed up exporting plan with parallelization (@zhilyaev)

### Other

* CLI refactoring (@zhilyaev)



# v0.12.3

### Dependencies

* Update golang to 1.16 (@zhilyaev)

* Update base alpine image to 3.14 (@zhilyaev)

* Update helm to 3.6.3 (@zhilyaev)



# v0.12.2

### New feature

* Add values to plandir (@zhilyaev)

### Fixed

* Make logs more informative (@zhilyaev)



# v0.12.1

### Fixed

* Revert removed `version` command and command-not-found (@zhilyaev)



# v0.12.0

### **Breaking change!**

* `install` and `dry-run` helm options are no more available to set (@zhilyaev)

* Parallel releases is now stable and enabled always (@zhilyaev)

### New feature

* [#101](https://github.com/helmwave/helmwave/issues/101) Add new `uninstall` command to delete deployed releases (@zhilyaev)

* [#101](https://github.com/helmwave/helmwave/issues/101) Add new `diff` command to show release diffs (@zhilyaev)

* [#101](https://github.com/helmwave/helmwave/issues/101) Add new `validate` command to check plan (@zhilyaev)

* Save manifests in plandir (@zhilyaev)

* Add `createnamespace` helm option (@zhilyaev)

* Add command to generate shell completion scripts (@zhilyaev)

* Repositories are now installed in parallel (@zhilyaev)

* Visualization graph for depends_on is now shown (@zhilyaev)

### Fixed

* [#124](https://github.com/helmwave/helmwave/issues/124) Fixed success when repository does not exist (@zhilyaev)

### CI

* Update golangci-lint, tune linters and fix linter issues (@r3nic1e)

### Other

* Refactor CLI commands (@zhilyaev)

* Add more tests (@zhilyaev)

* [#106](https://github.com/helmwave/helmwave/issues/106) Refactor with interfaces (@zhilyaev)



# v0.11.0

### New feature

* [#101](https://github.com/helmwave/helmwave/issues/101) Add new `list` command to list deployed releases (@r3nic1e)

* [#118](https://github.com/helmwave/helmwave/issues/118) Add support for remote values via go-getter (@r3nic1e)

* [#101](https://github.com/helmwave/helmwave/issues/101) Add new `status` command release statuses (@r3nic1e)

### CI

* Do not use deprecated docker.pkg.github.com (@r3nic1e)



# v0.10.1

### New feature

* Move v0.10.0 breaking change under `match-all-tags` flag (@r3nic1e)

### Other

* Ignore CLI code in code coverage (@r3nic1e)



# v0.10.0

### **Breaking change!**

* Releases must match all provided tags, not any of them (@r3nic1e)

### New feature

* [#105](https://github.com/helmwave/helmwave/issues/105) Add dependencies to plan even if they don't match tags (behind `plan-dependencies` flag) (@r3nic1e)

* Add warning if dependency was not found (@r3nic1e)

### Other

* Add a lot of tests (@r3nic1e)



# v0.9.6

### Fixed

* Fix failed releases were considered as successful (@r3nic1e)



# v0.9.4

### Fixed

* Add flag `kubedog-log-width` to set kubedog output width (@r3nic1e)



# v0.9.3

### **Breaking change!**

* [#90](https://github.com/helmwave/helmwave/issues/90) `depends_on` now requires release uniqname (`${release_name}@${namespace}`) (@r3nic1e)

### New feature

* Stop kubedog tracking after release end (@r3nic1e)

### Fixed

* [#90](https://github.com/helmwave/helmwave/issues/90) Fix `depends_on` for same release name in different projects (@r3nic1e)

### Dependencies

* Update some minor dependencies (@zhilyaev)

### Other

* Move log emoji formatter to separate project (@zhilyaev)



# v0.9.2

### CI

* Use goreleaser to build docker images (@zhilyaev)

* Use goreleaser to release to homebrew tap (@zhilyaev)

### Other

* Move the project to `helmwave` organization (@zhilyaev)



# v0.9.1

### CI

* Add release binaries for windows and darwin (@r3nic1e)



# v0.9.0

### New feature

* [#28](https://github.com/helmwave/helmwave/issues/28) Make new `plan` command with subcommands to generate plan (@r3nic1e)

* [#60](https://github.com/helmwave/helmwave/issues/60) Add support for release dependencies (@r3nic1e)

### Fixed

* [#70](https://github.com/helmwave/helmwave/issues/70) Fixed behavior when helm config dir doesn't exist (@r3nic1e)

### CI

* Add integration and unit tests (@r3nic1e)

* Add collecting code coverage (@r3nic1e)

* Use goreleaser to create new releases (@r3nic1e)



# v0.8.4

### Fixed

* [#59](https://github.com/helmwave/helmwave/issues/59) Fix panic when kubedog stops after helm upgrade (@r3nic1e)

### CI

* Fix release workflow for draft releases (@r3nic1e)

### Other

* Handle kubernetes client errors as debug log (@r3nic1e)

* Use permissions 0755 for created directories (@r3nic1e)



# v0.8.3

### New feature

* Add flag `kubedog-start-delay` to start kubedog a bit later. This mitigates potential kubedog failure for no release exists (@r3nic1e)



# v0.8.2

### New feature

* Add flag `kubedog-timeout` for setting kubedog timeout (@zhilyaev)

### Fixed

* Fix null pointer exception during unlocking repositories flock (@r3nic1e)



# v0.8.1

### New feature

* Add `kubedog-status-interval` flag to set kubedog interval between status updates (@r3nic1e)

### CI

* Add workflow for golangci-lint (@r3nic1e)



# v0.8.0

### New feature

* [#36](https://github.com/helmwave/helmwave/issues/36) Add support for kubedog to show realtime release upgrade (@zhilyaev)

* Add `manifest` command to show rendered manifests (@zhilyaev)



# v0.7.2

### Dependencies

* Update base alpine image to 3.13.1 (@zhilyaev)

### Other

* Change rendered manifest loglevel to `trace` (@zhilyaev)



# v0.7.1

### Fixed

* [#34](https://github.com/helmwave/helmwave/issues/34) Fix checking if repository is in plan (@zhilyaev)

### Dependencies

* [#33](https://github.com/helmwave/helmwave/issues/33) Update helm to 3.5.0 (@zhilyaev)



# v0.7.0

### New feature

* Use colors for logs (@zhilyaev)

### Fixed

* [#31](https://github.com/helmwave/helmwave/issues/31) Fix panic when there are no releases (@zhilyaev)



# v0.6.1

### Other

* Optimize plan repositories (@zhilyaev)



# v0.6.0

### New feature

* [#26](https://github.com/helmwave/helmwave/issues/26) Render releases manifests (@zhilyaev)



# v0.5.0

### New feature

* [#15](https://github.com/helmwave/helmwave/issues/15) Use planfile to apply changes (@zhilyaev)

* [#16](https://github.com/helmwave/helmwave/issues/16) Make logs prettier with emojis (@zhilyaev)

### Fixed

* [#20](https://github.com/helmwave/helmwave/issues/20) Exit with error if template rendering failed (@r3nic1e)

* Override sprig's hasKey function (@r3nic1e)


