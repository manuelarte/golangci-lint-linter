from ruamel.yaml.comments import CommentedMap

from golangci_lint_linter.rules import Rule, Report, validate_commented_map


def _is_alphabetical(original: [str]) -> bool:
    sorted_list: [str] = original.copy()
    sorted_list.sort()
    return sorted_list == original


class AlphabeticalSettings(object):
    """Rule to check that linters' settings are sorted alphabetically."""

    # TODO(manuelarte): Think about not sorting alphabetically,
    #  but as it's defined in enabled, or if `default: all`, then alphabetically.

    rule: Rule = Rule.GCI002

    def __init__(self) -> None:
        pass

    def lint(self, file: CommentedMap) -> [Report]:
        validate_commented_map(file)
        reports: [Report] = []
        linters: CommentedMap = file["linters"]
        settings: CommentedMap = linters.get("settings", default=CommentedMap())
        keys: list[str] = list(settings.keys())
        if keys and not _is_alphabetical(keys):
            reports.append(
                Report(self.rule, "linters.settings not sorted alphabetically.")
            )

        return reports
