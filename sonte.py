#! /usr/bin/env python3

import json
import os
import subprocess
import sys
import time

# Define globals.
JSON_DATA = {}
JSON_PATH = "./sonte.json"
TEMP_PATH = "./sonte.txt"
EDITOR = os.environ["EDITOR"] or os.environ["VISUAL"]

# Load data file if it exists.
if os.path.isfile(JSON_PATH):
    data = open(JSON_PATH, "r").read()
    JSON_DATA = json.loads(data)

# Run zero-argument command.
if len(sys.argv) <= 1:
    # Erase scratch file and open in editor.
    open(TEMP_PATH, "w").write("")
    subprocess.run([EDITOR, TEMP_PATH])

    # Add non-empty scratch to data object.
    if data := open(TEMP_PATH, "r").read():
        unix = str(int(time.time()))
        JSON_DATA[unix] = data.strip() + "\n"

    # Save modified data to file.
    data = json.dumps(JSON_DATA, indent=2, sort_keys=True)
    open(JSON_PATH, "w").write(data)

# Run multi-argument command.
else:
    # Clean tag arguments.
    tags = ["#" + tag.strip().lower() for tag in sys.argv[1:]]

    # Iterate and print matching entries.
    for unix, body in JSON_DATA.items():
        if any(tag in body for tag in tags):
            tobj = time.localtime(int(unix))
            strf = time.strftime("%c", tobj)
            print(f"# {strf}")
            print(body.strip())
            print("-----")
