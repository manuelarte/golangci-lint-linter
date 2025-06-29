import io

from ruamel.yaml.comments import CommentedMap
from golangci_lint_linter import read_yaml_file
from golangci_lint_linter.rules import Report
from golangci_lint_linter.rules.alphabetical_linters import AlphabeticalLinters


def test_enable_present_disable_absent_not_sorted_lint():
    f = io.StringIO(
        """
version: 2
linters:
  enable:
    - tagliatelle
    - funcorder
    """.lstrip()
    )

    commented_map: CommentedMap = read_yaml_file(f)
    rule: AlphabeticalLinters = AlphabeticalLinters()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert len(reports) == 1


def test_enable_sorted_lint():
    f = io.StringIO(
        """
version: 2
linters:
  enable:
    - funcorder
    - tagliatelle
    """.lstrip()
    )

    commented_map: CommentedMap = read_yaml_file(f)
    rule: AlphabeticalLinters = AlphabeticalLinters()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert not reports


def test_disable_present_enable_absent_not_sorted_lint():
    f = io.StringIO(
        """
version: 2
linters:
  default: all
  disable:
    - tagliatelle
    - funcorder
    """.lstrip()
    )

    commented_map: CommentedMap = read_yaml_file(f)
    rule: AlphabeticalLinters = AlphabeticalLinters()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert len(reports) == 1


def test_not_valid_golangci_but_valid_yaml_lint():
    f = io.StringIO(
        """
{ "not_valid": "not valid golangci.yml file" }
    """.lstrip()
    )

    commented_map: CommentedMap = read_yaml_file(f)
    rule: AlphabeticalLinters = AlphabeticalLinters()
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert len(reports) == 0


def test_enable_present_disable_absent_not_sorted_fix():
    f = io.StringIO(
        """
version: 2
linters:
  enable:
    - tagliatelle
    - funcorder
    """.lstrip()
    )

    commented_map: CommentedMap = read_yaml_file(f)
    rule: AlphabeticalLinters = AlphabeticalLinters()
    rule.fix(commented_map)
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert not reports

def test_disable_present_enable_absent_not_sorted_lint():
    f = io.StringIO(
        """
version: 2
linters:
  default: all
  disable:
    - tagliatelle
    - funcorder
    """.lstrip()
    )

    commented_map: CommentedMap = read_yaml_file(f)
    rule: AlphabeticalLinters = AlphabeticalLinters()
    rule.fix(commented_map)
    reports: list[Report] = rule.lint(commented_map)
    assert reports is not None
    assert not reports

