from typing import Any

from ruamel.yaml.comments import CommentedMap

from golangci_lint_linter.rules import (
    Rule,
    Report,
    _is_alphabetical,
    validate_commented_map,
)


class AlphabeticalLinters(object):
    """Rule to check that enable/disable linters are sorted alphabetically."""

    rule: Rule = Rule.GCI001

    def __init__(self) -> None:
        pass

    def lint(self, file: CommentedMap) -> list[Report]:
        validate_commented_map(file)
        reports: list[Report] = []
        linters: Any = file.get("linters", default=[])
        if not isinstance(linters, CommentedMap):
            return []
        enable: [str] = linters.get("enable", default=[])
        if enable and not _is_alphabetical(enable):
            reports.append(
                Report(self.rule, "linters.enable not sorted alphabetically.")
            )
        disable: list[str] = linters.get("disable", default=[])
        if disable and not _is_alphabetical(disable):
            reports.append(
                Report(self.rule, "linters.disable not sorted alphabetically.")
            )

        return reports
