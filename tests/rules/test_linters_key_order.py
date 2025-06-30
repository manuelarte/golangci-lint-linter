import io

from ruamel.yaml import YAML
from ruamel.yaml.comments import CommentedMap
from golangci_lint_linter import read_yaml_file, create_yaml_parser
from golangci_lint_linter.rules import Report
from golangci_lint_linter.rules.linters_keys_order import LintersKeyOrder


def test_linter_keys_not_sorted_lint():
    f = io.StringIO(
        """
version: 2
linters:
  disable:
    - funcorder
    - tagliatelle
  default: all  
  settings:
    tagliatelle:
        prop: value
    funcorder:
        prop: value
    """.lstrip()
    )

    yaml: YAML = create_yaml_parser()
    commented_map: CommentedMap = read_yaml_file(yaml, f)
    rule: LintersKeyOrder = LintersKeyOrder()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert len(reports) == 1


def test_linter_keys_sorted():
    f = io.StringIO(
        """
version: 2
linters:
  enable:
    - funcorder
    - tagliatelle
  disable:
    - errcheck  
  settings:
    funcorder:
        prop: value
    tagliatelle:
        prop: value
    """.lstrip()
    )

    yaml: YAML = create_yaml_parser()
    commented_map: CommentedMap = read_yaml_file(yaml, f)
    rule: LintersKeyOrder = LintersKeyOrder()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert not reports


def test_not_valid_golangci_but_valid_yaml():
    f = io.StringIO(
        """
{ "not_valid": "not valid golangci.yml file" }
    """.lstrip()
    )

    yaml: YAML = create_yaml_parser()
    commented_map: CommentedMap = read_yaml_file(yaml, f)
    rule: LintersKeyOrder = LintersKeyOrder()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert len(reports) == 0


def test_linter_keys_not_sorted_fix():
    f = io.StringIO(
        """
version: 2
linters:
  disable:
    - funcorder
    - tagliatelle
  default: all  
  settings:
    tagliatelle:
        prop: value
    funcorder:
        prop: value
    """.lstrip()
    )

    yaml: YAML = create_yaml_parser()
    commented_map: CommentedMap = read_yaml_file(yaml, f)
    rule: LintersKeyOrder = LintersKeyOrder()
    rule.fix(commented_map)
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert not reports
