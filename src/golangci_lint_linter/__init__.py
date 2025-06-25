from typing import TextIO
from ruamel.yaml import YAML
from ruamel.yaml.comments import CommentedMap

from golangci_lint_linter.errors import ProgramError


def _read_yaml_file(f: TextIO) -> None:
    yaml = YAML()
    yaml.preserve_quotes = True
    parsed = yaml.load(f)
    if not isinstance(parsed, CommentedMap):
        raise ProgramError(f"Error parsing the yaml file")


def main() -> int:
    print("Hello from golangci-lint-linter!")
    f = open("resources/examples/simple/.golangci.yml")
    if not f.readable():
        raise ValueError(f"Cannot read file {f}")
    _read_yaml_file(f)

    return 0
