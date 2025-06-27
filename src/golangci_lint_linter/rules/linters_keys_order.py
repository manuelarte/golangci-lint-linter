from typing import Any

from ruamel.yaml.comments import CommentedMap

from golangci_lint_linter.rules import Rule, Report, validate_commented_map


class LintersKeyOrder(object):
    """Rule to check that the linter fields are ordered by default/enable/disable."""

    rule: Rule = Rule.GCI003

    def __init__(self) -> None:
        pass

    def lint(self, file: CommentedMap) -> list[Report]:
        validate_commented_map(file)
        all_linter_keys: list[str] = ["default", "enable", "disable", "settings", "exclusions"]
        reports: list[Report] = []
        linters: Any = file.get("linters", default=[])
        if not isinstance(linters, CommentedMap):
            return []
        keys: list[str] = list(linters.keys())
        expected_keys_order: list[str] = list(
            filter(lambda x: x in keys.copy(), all_linter_keys)
        )
        if keys != expected_keys_order:
            reports.append(
                Report(
                    self.rule,
                    "linter keys not in expected order, default, enable, disable",
                )
            )

        return reports
