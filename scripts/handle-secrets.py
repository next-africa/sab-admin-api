#!/usr/bin/env python

try:
    import yaml
except ImportError:
    print("You need to have yaml install. Make sure you have run make install to install all dependencies for this project")
    exit(1)

import json
import sys

SECRET_VARIABLE_FILENAME = 'secret-env-variables.json'
SECRET_VALUES_FILENAME = 'secrets.json'
REPLACE_PREFIX = "Replace With"


def restore_secrets():
    secret_values_file = file(SECRET_VALUES_FILENAME)

    secrets_map = json.load(secret_values_file)

    data = yaml.load(sys.stdin)

    env_variables = data.get('env_variables', [])
    for key in env_variables:
        if key in secrets_map:
            current_value = env_variables[key]
            if current_value.startswith(REPLACE_PREFIX):
                env_variables[key] = secrets_map[key]

    yaml.safe_dump(data, sys.stdout, default_flow_style=False)


def remove_secrets():
    secret_variables_file = file(SECRET_VARIABLE_FILENAME)

    secrets = json.load(secret_variables_file)

    data = yaml.load(sys.stdin)

    key_map = {}

    env_variables = data.get('env_variables', [])
    for key in env_variables:
        if key in secrets:
            current_value = env_variables[key]
            if not current_value.startswith(REPLACE_PREFIX):
                key_map[key] = current_value
                env_variables[key] = "{} {}".format(REPLACE_PREFIX, key)

    secret_variables_file.close()

    yaml.safe_dump(data, sys.stdout, default_flow_style=False)

    json.dump(key_map, file(SECRET_VALUES_FILENAME, 'w'))

    exit(0)

if __name__ == "__main__":
    action = sys.argv[1]

    if action == "remove":
        remove_secrets()
    else:
        restore_secrets()
