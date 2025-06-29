from typing import Any

from ruamel.yaml.comments import CommentedMap

from golangci_lint_linter.rules import Rule, Report, validate_commented_map


def _get_linters(file: CommentedMap) -> Any:
    return file.get("linters")


class LintersKeyOrder(object):
    """Rule to check that the linter fields are ordered by default/enable/disable."""

    rule: Rule = Rule.GCI003
    all_sorted_linter_keys: list[str] = [
        "default",
        "enable",
        "disable",
        "settings",
        "exclusions",
    ]

    def __init__(self) -> None:
        pass

    def lint(self, file: CommentedMap) -> list[Report]:
        validate_commented_map(file)

        reports: list[Report] = []
        linters: Any = _get_linters(file)
        if not isinstance(linters, CommentedMap):
            return []
        keys: list[str] = list(linters.keys())
        expected_keys_order: list[str] = list(
            filter(lambda x: x in keys.copy(), self.all_sorted_linter_keys)
        )
        if keys != expected_keys_order:
            reports.append(
                Report(
                    self.rule,
                    "linter keys not in expected order, default, enable, disable",
                )
            )

        return reports

    def fix(self, file: CommentedMap) -> None:
        order_dict = {item: idx for idx, item in enumerate(self.all_sorted_linter_keys)}
        linters: Any = _get_linters(file)
        if not isinstance(linters, CommentedMap):
            return
        for key in sorted(linters, key=lambda x: order_dict[x], reverse=True):
            value = linters.pop(key)
            linters.insert(0, key, value)
