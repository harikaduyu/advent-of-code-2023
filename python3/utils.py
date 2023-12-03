import os


def readlines(day_string: str):
    input_file_path = os.path.abspath(
        os.path.join(
            os.path.abspath(__file__), "..", "..", "inputs", f"{day_string}.txt"
        )
    )
    with open(input_file_path) as f:
        line = f.readline()
        while line:
            yield line
            line = f.readline()
