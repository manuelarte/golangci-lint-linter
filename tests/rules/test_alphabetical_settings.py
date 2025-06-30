import io

from ruamel.yaml import YAML
from ruamel.yaml.comments import CommentedMap
from golangci_lint_linter import read_yaml_file, create_yaml_parser
from golangci_lint_linter.rules import Report
from golangci_lint_linter.rules.alphabetical_settings import AlphabeticalSettings


def test_settings_not_sorted_lint():
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

    yaml: YAML = create_yaml_parser()
    commented_map: CommentedMap = read_yaml_file(yaml, f)
    rule: AlphabeticalSettings = AlphabeticalSettings()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert len(reports) == 1
    assert (
        str(reports[0])
        == "GCI002: [LineCol(6, 4)]linters.settings not sorted alphabetically."
    )


def test_settings_sorted_lint():
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

    yaml: YAML = create_yaml_parser()
    commented_map: CommentedMap = read_yaml_file(yaml, f)
    rule: AlphabeticalSettings = AlphabeticalSettings()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert not reports


def test_not_valid_golangci_but_valid_yaml_lint():
    f = io.StringIO(
        """
{ "not_valid": "not valid golangci.yml file" }
    """.lstrip()
    )

    yaml: YAML = create_yaml_parser()
    commented_map: CommentedMap = read_yaml_file(yaml, f)
    rule: AlphabeticalSettings = AlphabeticalSettings()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert not reports


def test_settings_not_sorted_fix():
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

    yaml: YAML = create_yaml_parser()
    commented_map: CommentedMap = read_yaml_file(yaml, f)
    rule: AlphabeticalSettings = AlphabeticalSettings()
    rule.fix(commented_map)
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert not reports
