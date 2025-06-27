from typing import TextIO

import click

from ruamel.yaml import YAML
from ruamel.yaml.comments import CommentedMap

from golangci_lint_linter.errors.errors import ProgramError
from golangci_lint_linter.rules import Ruler, Report
from golangci_lint_linter.rules.alphabetical_linters import AlphabeticalLinters
from golangci_lint_linter.rules.alphabetical_settings import AlphabeticalSettings
from golangci_lint_linter.rules.linters_keys_order import LintersKeyOrder
from golangci_lint_linter.rules.disable_linters_reason import DisableLintersReason


def get_all_rules() -> list[Ruler]:
    return [
        AlphabeticalLinters(),
        AlphabeticalSettings(),
        LintersKeyOrder(),
        DisableLintersReason()
    ]


def read_yaml_file(f: TextIO) -> CommentedMap:
    yaml = YAML()
    yaml.preserve_quotes = True
    parsed = yaml.load(f)
    if not isinstance(parsed, CommentedMap):
        raise ProgramError("Error parsing the yaml file")
    return parsed


@click.command()
@click.argument("file", type=click.File("r"), default="./golangci-lint.yml")
def main(file: TextIO) -> int:
    """Linter for the golanci-lint configuration file."""
    if not file.readable():
        click.echo(message=f"Cannot read: {file.name}.", err=True)
    click.echo(message=f"Linting: {file.name}.")
    reports: list[Report] = []
    try:
        commented_map: CommentedMap = read_yaml_file(file)
        lint_rules: list[Ruler] = get_all_rules()
        for rule in lint_rules:
            try:
                rule_reports: list[Report] = rule.lint(commented_map)
                for rule_report in rule_reports:
                    reports.append(rule_report)
            except Exception as e:
                click.secho(
                    message=f"Internal error applying rule: {rule.rule}: {e}", err=True
                )
        click.echo(f"Summary: {len(reports)} error(s).")
        for report in reports:
            click.secho(message=str(report), fg="red", err=True)
    except ProgramError as e:
        click.echo(message=str(e), err=True)

    return 0 if len(reports) == 0 else 1
