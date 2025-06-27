from typing import Any

from ruamel.yaml import CommentToken, CommentedSeq
from ruamel.yaml.comments import CommentedMap

from golangci_lint_linter.rules import Rule, Report, validate_commented_map


class DisableLintersReason(object):
    """Rule to check that the disable linters have a reason."""

    rule: Rule = Rule.GCI003

    def __init__(self) -> None:
        pass

    def lint(self, file: CommentedMap) -> list[Report]:
        validate_commented_map(file)
        reports: list[Report] = []
        linters: Any = file.get("linters", default=[])
        if not isinstance(linters, CommentedMap):
            return []
        disable: Any = linters.get("disable", default=[])
        if not isinstance(disable, CommentedSeq):
            return []
        starting_idx: int = 0
        for index_comment in disable.ca.items:
            if index_comment != starting_idx:
                reports.append(
                    Report(
                        self.rule,
                        f"linters.disable.{disable[starting_idx]} has no reason.",
                    )
                )
                break
            starting_idx = starting_idx + 1

        return reports
