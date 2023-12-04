import os
import inspect


def read_input_lines(example: int | None = None):
    caller_file = inspect.stack()[1].filename
    day_string = caller_file.split("day")[-1].split(".py")[0]
    file_name = "Q_input.txt" if example is None else f"example_{example}.txt"
    input_file_path = os.path.abspath(
        os.path.join(
            os.path.abspath(__file__),
            "..",
            "..",
            "inputs",
            f"day{day_string}",
            file_name,
        )
    )
    with open(input_file_path) as f:
        line = f.readline()
        while line:
            yield line
            line = f.readline()
