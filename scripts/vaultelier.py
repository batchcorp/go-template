#!/usr/bin/env python
#
# Vaultelier -- script for finding/replacing vault secrets in your k8s deploy cfg
#
# Usage: cat config.yaml | python vaultelier.py | kubectl apply -f -
#

import os
import re
import sys
import json
import subprocess

vault_binary = ""


def validate_env():
    # Verify we've got the vault binary
    try:
        output = subprocess.check_output(['which vault'], shell=True, stderr=subprocess.STDOUT)
    except subprocess.CalledProcessError as exec_err:
        return "unable to determine 'vault' existence: %s (return code: %d)" % (exec_err.output, exec_err.returncode)

    vault_path = output.rstrip()

    if not os.path.exists(vault_path):
        return "'%s' does not exist" % vault_path

    # Verify that we get a valid status from vault
    try:
        subprocess.check_output([vault_path, 'status'], stderr=subprocess.STDOUT)
    except subprocess.CalledProcessError as exec_err:
        return "'vault status' returned a non-0 result: %d" % exec_err.returncode

    global vault_binary

    vault_binary = vault_path

    return None


def generate_want_secrets(input_str):
    # {"object": "key"}
    want_secrets = {}

    for line in input_str.splitlines():
        # If line contains secret syntax, add to map
        # Looking for {{vault:object:key}} in line
        match_obj = re.match(".+{{(vault):(.+):(.+)}}.+", line)
        if match_obj is None:
            continue

        secret_type = match_obj.group(1)
        secret_obj  = match_obj.group(2)
        secret_key  = match_obj.group(3)

        if secret_type not in want_secrets:
            want_secrets[secret_type] = {
                secret_obj: secret_key
            }
        else:
            want_secrets[secret_type][secret_obj] = secret_key

    return want_secrets


def fetch_vault_secrets(want_secrets):
    have_secrets = {}

    if 'vault' not in want_secrets:
        return None, "want_secrets does not contain any 'vault' secrets"

    for obj, key in want_secrets['vault'].items():
        if "%s:%s" % (obj, key) in have_secrets:
            # Already have secret
            continue

        # Fetch secret from vault
        cmd = [vault_binary, 'kv', 'get', '-format', 'json', obj]

        try:
            output = subprocess.check_output(cmd, stderr=subprocess.STDOUT)
        except subprocess.CalledProcessError as exec_err:
            return None, "vault lookup failed: %s (return code: %d)" % (exec_err.output, exec_err.returncode)

        # Parse JSON
        try:
            secrets = json.loads(output)
        except Exception as exc:
            return None, "unable to decode vault JSON output: %s" % exc.message

        if 'data' not in secrets:
            return None, "unexpected data structure in vault response"

        if 'data' not in secrets['data']:
            return None, "unexpected data structure in vault response"

        for key_name, secret in secrets['data']['data'].items():
            have_secrets["%s:%s" % (obj, key_name)] = secret

    return have_secrets, None


def replace_secrets(input_file, have_secrets):
    lines = ""

    for line in input_file.splitlines():
        match_obj = re.match(".+{{vault:(.+:.+)}}.+", line)

        if match_obj is None:
            lines = lines + line + "\n"
            continue

        # Match, if we don't have a secret -> error!
        want = match_obj.group(1)

        if want not in have_secrets:
            return None, "unable to lookup secret for '%s'" % want

        # Replace
        line = re.sub("{{vault:(.+:.+)}}", have_secrets[want], line)
        lines = lines + line + "\n"

    return lines, None


def main():
    # Ensure 'vault' CLI tool exists
    err = validate_env()
    if err is not None:
        print("ERROR: Unable to validate env: %s" % err)
        sys.exit(1)

    # Read input
    input_str = sys.stdin.read()

    # # Build map of all secrets we need
    want_secrets = generate_want_secrets(input_str)

    if len(want_secrets) < 1:
        print("NOTICE: No secrets found in provided input file")
        sys.exit(0)

    # Fetch secrets
    have_secrets, err = fetch_vault_secrets(want_secrets)
    if err is not None:
        print("ERROR: Unable to fetch secrets from vault: %s" % err)
        sys.exit(1)

    # Replace, provide modified output
    output_file, err = replace_secrets(input_str, have_secrets)
    if err is not None:
        print("ERROR: Unable to replace secrets: %s" % err)
        sys.exit(1)

    print(output_file)

    sys.exit(0)


if __name__ == '__main__':
    main()
