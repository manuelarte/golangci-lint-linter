import io

from ruamel.yaml.comments import CommentedMap
from golangci_lint_linter import read_yaml_file
from golangci_lint_linter.analysis import Report
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
