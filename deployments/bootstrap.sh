#!/usr/bin/env bash

ansible-playbook -u iot -i hosts playbook.yml $@