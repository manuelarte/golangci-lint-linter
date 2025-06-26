from enum import Enum
from typing import Protocol
from ruamel.yaml.comments import CommentedMap


class Rule(Enum):
    GCI001 = "GCI001"


class Report(object):
    """Report a lint error"""

    rule: Rule
    msg: str

    def __init__(self, rule: Rule, msg: str) -> None:
        self.rule = rule
        self.msg = msg


class Ruler(Protocol):
    rule: Rule

    def lint(self, file: CommentedMap) -> [Report]:
        """Lint the .golangci file."""
        ...
