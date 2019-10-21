# ecflow-checker

An ecflow tool to check node status.

## Installing

Use `Makefile` to build `ecflow_checker` command.

## Getting started

Start `ecflow_grpc_server` from [perillaroc/ecflow-client-cpp](https://github.com/perillaroc/ecflow-client-cpp).

Run `ecflow_checker check` command to check nodes listed in config files.

```bash
ecflow_checker check --config-dir=some/path/to/dir
```

Results are as follows: 

```
-- config file: nwpc_op/gda_grapes_gfs.yml
gda_grapes_gfs_v2.4 00H:
[✔] Checking for start
[✔] Checking for complete
gda_grapes_gfs_v2.4 06H:
[✔] Checking for start
[━] Checking for complete
gda_grapes_gfs_v2.4 12H: Ignore
gda_grapes_gfs_v2.4 18H: Ignore
```

## Configure

See [conf/example.yml](conf/example.yml) for more information.

## License

Copyright 2019, perillaroc.

`ecflow-checker` is licensed under [MIT License](./LICENSE.md).