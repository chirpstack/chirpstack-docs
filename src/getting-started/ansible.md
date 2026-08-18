# Ansible playbook

[Ansible](https://docs.ansible.com/) is an open-source tool for automating
deployment and server-management related steps. ChirpStack provides a so-called
Ansible Playbook which can be used to deploy and configure ChirpStack and
optionally setup TLS certificates. The source of this Playbook, including
additional documentation, can be found at
[https://github.com/chirpstack/chirpstack-ansible-playbook](https://github.com/chirpstack/chirpstack-ansible-playbook).

## Deploying ChirpStack

To deploy ChirpStack, including its dependencies to a remote VM, follow the
steps documented in
[chirpstack-ansible-playbook](https://github.com/chirpstack/chirpstack-ansible-playbook).
In short, when Ansible is installed and configuration parameters in the playbook
have been updated, ChirpStack can be deployed using the following command:

```bash
ansible-playbook -i inventory deploy.yml
```
