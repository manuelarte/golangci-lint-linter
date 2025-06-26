from enum import Enum
from typing import Protocol
from ruamel.yaml.comments import CommentedMap

from golangci_lint_linter.errors.errors import ProgramError


class Rule(Enum):
    GCI001 = "GCI001"
    GCI002 = "GCI002"


class Report(object):
    """Report a lint error"""

    rule: Rule
    msg: str

    def __init__(self, rule: Rule, msg: str) -> None:
        self.rule = rule
        self.msg = msg

    def __repr__(self) -> str:
        return f"{self.rule.name}: {self.msg}"


class Ruler(Protocol):
    rule: Rule

    def lint(self, file: CommentedMap) -> [Report]:
        """Lint the .golangci file."""
        ...


def validate_commented_map(file: CommentedMap) -> None:
    if not file:
        raise ProgramError("file not found")
    if not isinstance(file, CommentedMap):
        raise ProgramError("file is not CommentedMap")
    return
