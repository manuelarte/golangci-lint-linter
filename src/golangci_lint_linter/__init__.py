from typing import TextIO

import click

from ruamel.yaml import YAML
from ruamel.yaml.comments import CommentedMap

from golangci_lint_linter.errors import ProgramError


def read_yaml_file(f: TextIO) -> CommentedMap:
    yaml = YAML()
    yaml.preserve_quotes = True
    parsed = yaml.load(f)
    if not isinstance(parsed, CommentedMap):
        raise ProgramError("Error parsing the yaml file")
    return parsed


@click.command()
@click.argument("file", type=click.File("r"))
def main(file: TextIO) -> int:
    """Linter for the golanci-lint configuration file."""
    if not file.readable():
        raise ValueError(f"Cannot read file {file}")
    read_yaml_file(file)

    return 0
