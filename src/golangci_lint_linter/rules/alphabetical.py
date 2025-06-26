from golangci_lint_linter import ProgramError
from ruamel.yaml.comments import CommentedMap

from golangci_lint_linter.rules import Rule, Report


def _is_alphabetical(original: [str]) -> bool:
    sorted_list: [str] = original.copy()
    sorted_list.sort()
    return sorted_list == original


class Alphabetical(object):
    """Rule to check that enable/disable linters are sorted alphabetically."""

    rule: Rule = Rule.GCI001

    def __init__(self) -> None:
        pass

    def lint(self, file: CommentedMap) -> [Report]:
        """Lint the .golangci file."""
        if not file:
            raise ProgramError("file not found")
        if not isinstance(file, CommentedMap):
            raise ProgramError("file is not CommentedMap")
        reports: [Report] = []
        linters: CommentedMap = file["linters"]
        enable: [str] = linters.get("enable", default=[])
        if enable and not _is_alphabetical(enable):
            reports.append(
                Report(self.rule, "linters.enable not sorted alphabetically.")
            )
        disable: [str] = linters.get("disable", default=[])
        if disable and not _is_alphabetical(disable):
            reports.append(
                Report(self.rule, "linters.disable not sorted alphabetically.")
            )

        return reports
