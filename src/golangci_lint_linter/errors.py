class ProgramError(Exception):
    """Exception raised for the tool golangci-lint-linter.

    Attributes:
        message -- explanation of the error
    """

    message: str
    original_error: Exception | None

    def __init__(self, message: str, original_error: Exception | None = None) -> None:
        self.message = message
        self.original_error = original_error
