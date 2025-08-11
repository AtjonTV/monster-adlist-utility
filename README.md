# Monster Adlist Utility
This is a tool that can create gigantic adlists for DNS-based blocking.

It was initially written as a [Bash script](https://gist.github.com/AtjonTV/4b41e123806ffb7faa9f2158f9625c5c), but was rewritten in Golang to add more complex features.

For the creation of the "monster.list", a `sources.yaml` is used to source the lists that are used.

To generate the list, you have to run:
```bash
# Compile
./build.sh

# Run
./monster-update
```

You can optionally specify the location of the sources.yaml:
```bash
./monster-update -source ./sources.yaml
```

A "monster_base.list" symlink can be created when adding the "-relink" option.

Additional options can be viewed using:
```bash
./monster-update --help
```

## License
Monster Adlist Utility is licensed under the [EUPL v1.2 or later](./LICENSE.txt).
