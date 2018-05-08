## Go Cloud VPN

A pure self hosted OpenVPN levaging the cloud.

### Motivation

Because VPN configuration and setup are too complicated, this project gathers
some ideas, tools I used to manage my own server.

Key ideas:

  - No user "password"
  - Keep your private certificate keys encrypted locally with you
  - Low cost: use cloud providers
  - Manage your own CA: deals with renewing, adding, revoquing for your
  - Generate the client configuration for easy plug and play
  - Easily ditch and recreate VPN server using cloud provider

### Features

  - TOPT Authentication
  - Run && Ditch cloud VMs
  - Handle local CA and certificate in local encrypted storage
  - Generate configuration for clients

### TODO:

  - Cleanup Ansible code
  - Embed certificates into ovpn configuration for easier deployments
