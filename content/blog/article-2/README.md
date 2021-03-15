# example-development-k3d-skaffold
This is a demo of a developing at local environment using k3s and skaffold.

[![Open in Cloud Shell](https://gstatic.com/cloudssh/images/open-btn.svg)](https://ssh.cloud.google.com/cloudshell/editor?cloudshell_git_repo=https://github.com/cloudandbuild/example-development-k3d-skaffold.git&cloudshell_tutorial=README.md)
## prerequisite

- OS
  - ubuntu 20.04 or cloudshell
- Tools
  - docker
  - kubectl
  - make
  - curl

## install k3d
```
curl -s https://raw.githubusercontent.com/rancher/k3d/main/install.sh | bash
```

## install skaffold
```
curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64 && \
sudo install skaffold /usr/local/bin/
```

## create cluster 
```
make create
```

## develop using skaffold
```
make dev
```

