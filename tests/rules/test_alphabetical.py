import io

from ruamel.yaml.comments import CommentedMap
from golangci_lint_linter import read_yaml_file
from golangci_lint_linter.rules import Report
from golangci_lint_linter.rules.alphabetical import Alphabetical


def test_enable_present_disable_absent_not_sorted():
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
    rule: Alphabetical = Alphabetical()
    reports: [Report] = rule.lint(commented_map)
    assert reports is not None
    assert len(reports) == 1


def test_disable_present_enable_absent_not_sorted():
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
    rule: Alphabetical = Alphabetical()
    reports: [Report] = rule.lint(commented_map)
    assert reports is not None
    assert len(reports) == 1
