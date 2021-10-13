#!/bin/bash

set -x
curl -X POST http://localhost:7071/api/clctranspilerfunction \
    -H "Content-Type: application/yaml" \
    --data-binary @- <<EOF
storage:
  filesystems:
    - name: "OEM"
      mount:
        device: "/dev/disk/by-label/OEM"
        format: "btrfs"
  files:
    - filesystem: "OEM"
      path: "/grub.cfg"
      mode: 0644
      append: true
      contents:
        inline: |
          set linux_console="console=ttyAMA0,115200n8 console=tty1"
          set linux_append="flatcar.autologin usbcore.autosuspend=-1"
passwd:
  users:
    - name: core
      ssh_authorized_keys:
        - ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAICMNbha1AXiaK2YwkeXHcYXPMgSq636zsV2ZJ6FP5BUi azureuser@jeremi-spot-dev1
EOF
