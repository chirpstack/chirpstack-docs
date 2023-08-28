# Ansible and Vagrant

[Ansible](https://docs.ansible.com/) is an open-source tool for automating
deployment and server-management related steps. ChirpStack provides a so-called
Ansible Playbook which can be used to deploy and configure ChirpStack and
optionally setup TLS certificates. The source of this Playbook, including
additional documentation, can be found at [https://github.com/chirpstack/chirpstack-ansible-playbook](https://github.com/chirpstack/chirpstack-ansible-playbook).

## Local VM deployment using Vagrant

[Vagrant](https://www.vagrantup.com/) is a tool for automating the creation
of virtual machines. It can integrate with Ansible so that it not only create
the VM for you, but also will install and configure ChirpStack.

After following the instructions documented in the [chirpstack-ansible-playbook](https://github.com/chirpstack/chirpstack-ansible-playbook)
repository, you can create a local test environment running inside a VM using
the following command:

```bash
vagrant up
```

As this is using exactly the same Ansible Playbook as for remote deployments,
this can also be used for testing before doing a remote deployment, e.g.
when making modifications to the playbook.

## Remote deployment

Ansible can also be used to perform remote deployments. After following the
steps documented in [chirpstack-ansible-playbook](https://github.com/chirpstack/chirpstack-ansible-playbook),
the following will perform a remote deployment:

```bash
ansible-playbook -i inventory deploy.yml
```
