import io

import pytest
from ruamel.yaml.comments import CommentedMap

from golangci_lint_linter import _read_yaml_file


def test_comments_are_available():
    f = io.StringIO(
        """
        version: 2
        formatters:
          enable:
            - gofmt
            - gofumpt
            - gci
    """)

    commented_map: CommentedMap = _read_yaml_file(f)
    assert commented_map['version'] == 2
