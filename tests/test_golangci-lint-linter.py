import io

from click.testing import CliRunner
from ruamel.yaml.comments import CommentedMap
from ruamel.yaml.tokens import CommentToken

from golangci_lint_linter import read_yaml_file, main


def test_comments_are_available():
    f = io.StringIO(
        """
        # version 2 of golangci-lint
        version: 2
        # now the formatter block, with the enable formatters
        formatters: # here are defined the formatters
          # enable the most used formatters  
          enable:
            - gofmt
            - gofumpt
            - gci
    """.lstrip()
    )

    commented_map: CommentedMap = read_yaml_file(f)
    formatters_comments: [CommentToken | None] = commented_map.ca.items["formatters"]
    formatters_comment_token: CommentToken = formatters_comments[2]

    assert isinstance(formatters_comment_token, CommentToken)
    assert (
        formatters_comment_token.value
        == """# here are defined the formatters
          # enable the most used formatters  \n"""
    )


def test_main_valid_golangci():
    runner = CliRunner()
    result = runner.invoke(main, ["../resources/examples/simple/.golangci.yml"])
    assert result.exit_code == 0


def test_main_invalid_golangci():
    runner = CliRunner()
    result = runner.invoke(main, ["../resources/examples/notvalid/.golangci.yml"])
    assert result.exit_code == 0
