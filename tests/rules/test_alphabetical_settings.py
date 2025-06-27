import io

from ruamel.yaml.comments import CommentedMap
from golangci_lint_linter import read_yaml_file
from golangci_lint_linter.rules import Report
from golangci_lint_linter.rules.alphabetical_settings import AlphabeticalSettings


def test_settings_not_sorted():
    f = io.StringIO(
        """
version: 2
linters:
  enable:
    - funcorder
    - tagliatelle
  settings:
    tagliatelle:
        prop: value
    funcorder:
        prop: value
    """.lstrip()
    )

    commented_map: CommentedMap = read_yaml_file(f)
    rule: AlphabeticalSettings = AlphabeticalSettings()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert len(reports) == 1


def test_settings_sorted():
    f = io.StringIO(
        """
version: 2
linters:
  enable:
    - funcorder
    - tagliatelle
  settings:
    funcorder:
        prop: value
    tagliatelle:
        prop: value
    """.lstrip()
    )

    commented_map: CommentedMap = read_yaml_file(f)
    rule: AlphabeticalSettings = AlphabeticalSettings()
    reports: [Report] = rule.lint(commented_map)
    assert reports is not None
    assert not reports


def test_not_valid_golangci_but_valid_yaml():
    f = io.StringIO(
        """
{ "not_valid": "not valid golangci.yml file" }
    """.lstrip()
    )

    commented_map: CommentedMap = read_yaml_file(f)
    rule: AlphabeticalSettings = AlphabeticalSettings()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert len(reports) == 0
