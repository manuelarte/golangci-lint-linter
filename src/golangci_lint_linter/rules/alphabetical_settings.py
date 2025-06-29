from typing import Any

from ruamel.yaml.comments import CommentedMap

from golangci_lint_linter.rules import (
    Rule,
    Report,
    _is_alphabetical,
    validate_commented_map,
)


def _get_settings(file: CommentedMap) -> Any | None:
    linters: Any = file.get("linters", default=[])
    if not isinstance(linters, CommentedMap):
        return None
    return linters.get("settings")


class AlphabeticalSettings(object):
    """Rule to check that linters' settings are sorted alphabetically."""

    # TODO(manuelarte): Think about not sorting alphabetically,
    #  but as it's defined in enabled, or if `default: all`, then alphabetically.

    rule: Rule = Rule.GCI002

    def __init__(self) -> None:
        pass

    def lint(self, file: CommentedMap) -> list[Report]:
        validate_commented_map(file)
        reports: list[Report] = []
        settings: Any = _get_settings(file)
        if not isinstance(settings, CommentedMap):
            return []
        keys: list[str] = list(settings.keys())
        if keys and not _is_alphabetical(keys):
            reports.append(
                Report(self.rule, f"[{settings.lc}]linters.settings not sorted alphabetically.")
            )

        return reports

    def fix(self, file: CommentedMap) -> None:
        settings: Any = _get_settings(file)
        if isinstance(settings, CommentedMap):
            for key in sorted(settings, reverse=True):
                value = settings.pop(key)
                settings.insert(0, key, value)
