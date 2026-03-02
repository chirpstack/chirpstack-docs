# Command-line interface

ChirpStack provides a command-line interface (CLI) for tasks such as resetting
user passwords, creating API keys, and managing device profiles. These commands
are useful for automated deployments, scripting, and administration.

## Usage

The CLI is invoked using the `chirpstack` command:

```bash
chirpstack --config <DIRECTORY> <COMMAND> [OPTIONS]
```

Where `--config` points to the ChirpStack configuration directory (containing
`chirpstack.toml` and region configuration files).

## Available commands

### reset-password

Reset a user's password without API access. This command is useful for:

- Initial setup of fresh installations
- Automated deployment scripts
- Recovery from forgotten passwords

**Usage:**

```bash
chirpstack --config /etc/chirpstack reset-password --email <EMAIL> [OPTIONS]
```

**Options:**

| Option | Description |
|--------|-------------|
| `-e, --email <EMAIL>` | User email address (required) |
| `-p, --password-file <FILE>` | Path to file containing new password |
| `--stdin` | Read password from stdin |

**Examples:**

```bash
# Interactive password reset (prompts for password twice)
chirpstack --config /etc/chirpstack reset-password -e admin@example.com

# Password from file (recommended for scripts)
read -rs PW && printf '%s' "$PW" > /tmp/pw.txt && unset PW
chmod 600 /tmp/pw.txt
chirpstack --config /etc/chirpstack reset-password -e admin@example.com -p /tmp/pw.txt
rm /tmp/pw.txt

# Password from stdin (pipeline use, e.g. from a secrets manager)
my-secrets-tool get admin-password | chirpstack --config /etc/chirpstack reset-password -e admin@example.com --stdin
```

**Security notes:**

- When using a file, set restrictive permissions (`chmod 600`) and delete it after use
- Avoid embedding a plaintext password in a command (e.g. `echo "pass" | ...`): it appears in shell history and is visible via process inspection
- For scripts, prefer the file method or pipe from a secrets manager — see the [Password handling](#password-handling) table below

**Password requirements:**

- Minimum 8 characters
- Maximum 128 characters
- NIST 800-63b compliant (no arbitrary complexity requirements)

### create-api-key

Create a global API key for administrative access.

**Usage:**

```bash
chirpstack --config /etc/chirpstack create-api-key --name <NAME>
```

**Example:**

```bash
chirpstack --config /etc/chirpstack create-api-key --name "automation-key"
```

Output:

```bash
id: <UUID>
token: <JWT_TOKEN>
```

### configfile

Print the configuration template to stdout.

**Usage:**

```bash
chirpstack --config /etc/chirpstack configfile
```

### import-device-profiles

Import LoRaWAN device profiles from a directory.

**Usage:**

```bash
chirpstack --config /etc/chirpstack import-device-profiles <DIRECTORY>
```

### migrate-device-profile-templates

Migrate device-profile templates to device profiles.

**Usage:**

```bash
chirpstack --config /etc/chirpstack migrate-device-profile-templates <DIRECTORY>
```

### print-ds

Print the device-session for debugging (requires --dev-eui).

**Usage:**

```bash
chirpstack --config /etc/chirpstack print-ds --dev-eui <DEV_EUI>
```

## Security considerations

### Password handling

The CLI provides multiple methods for password input, each with different
security implications:

| Method | Security | Use case |
|--------|----------|----------|
| Interactive prompt | Highest | Manual administration |
| File (0600 permissions) | Medium | Scripted deployments |
| Stdin | Low | Pipelines (ensure no history) |

For scripted deployments using password files:

```bash
# Read password without it appearing in shell history
read -rs PW && printf '%s' "$PW" > /tmp/pw.txt && unset PW
chmod 600 /tmp/pw.txt

# Use in command
chirpstack --config /etc/chirpstack reset-password -e admin@example.com -p /tmp/pw.txt

# Immediately delete
rm /tmp/pw.txt
```

### Default credentials

Upon first installation, ChirpStack creates a default admin user with the
credentials:

- Username: `admin`
- Password: `admin`

**It is strongly recommended to change this password immediately upon first
login** using either the web interface or the `reset-password` command.