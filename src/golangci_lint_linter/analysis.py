from golangci_lint_linter.rules import Rule


class Report(object):
    """Report a lint error"""

    rule: Rule
    msg: str

    def __init__(self, rule: Rule, msg: str) -> None:
        self.rule = rule
        self.msg = msg
