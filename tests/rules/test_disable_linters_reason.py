import io

from ruamel.yaml.comments import CommentedMap
from golangci_lint_linter import read_yaml_file
from golangci_lint_linter.rules import Report
from golangci_lint_linter.rules.disable_linters_reason import DisableLintersReason


def test_no_disable_linters():
    f = io.StringIO(
        """
version: 2
linters:
  default: all  
  settings:
    tagliatelle:
        prop: value
    funcorder:
        prop: value
    """.lstrip()
    )

    commented_map: CommentedMap = read_yaml_file(f)
    rule: DisableLintersReason = DisableLintersReason()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert not reports


def test_disable_linters_no_reason():
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

    commented_map: CommentedMap = read_yaml_file(f)
    rule: DisableLintersReason = DisableLintersReason()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert len(reports) == 1
    assert (
        str(reports[0]) == "GCI003: [LineCol(6, 4)]linters.disable have no reason(s)."
    )


def test_disable_linters_with_reason():
    f = io.StringIO(
        """
version: 2
linters:
  enable:
    - funcorder
    - tagliatelle
  disable:
    - errcheck  # not applicable
  settings:
    funcorder:
        prop: value
    tagliatelle:
        prop: value
    """.lstrip()
    )

    commented_map: CommentedMap = read_yaml_file(f)
    rule: DisableLintersReason = DisableLintersReason()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert not reports


def test_not_valid_golangci_but_valid_yaml():
    f = io.StringIO(
        """
{ "not_valid": "not valid golangci.yml file" }
    """.lstrip()
    )

    commented_map: CommentedMap = read_yaml_file(f)
    rule: DisableLintersReason = DisableLintersReason()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert not reports
