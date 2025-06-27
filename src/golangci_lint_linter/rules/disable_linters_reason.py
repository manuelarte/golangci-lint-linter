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
        for idx, linter in enumerate(disable):
            if idx > len(disable.ca.items) - 1 or not disable.ca.items[idx]:
                reports.append(
                    Report(
                        self.rule,
                        f"linters.disable.{linter} has no reason.",
                    )
                )
                break
            if idx < len(disable) and disable.ca.items[idx]:
                comment: CommentToken = disable.ca.items[idx]
                if len(comment) > 0 and len(comment[0].value) > 2:
                    continue

        return reports
