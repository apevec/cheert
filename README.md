# [Cheert](https://github.com/apevec/cheert)

## Run in OpenShift
```
$ export DOCKER_ID_USER="username"
$ oc new-app $DOCKER_ID_USER/s2i-cheert-centos7~https://github.com/$DOCKER_ID_USER/cheert.git
```

## Manually trigger build in OpenShift
```
$ oc start-build cheert
```

## How to run it?
```
$ npm install
$ npm run init
$ npm run watch
```

## How to update JavaScript dependencies (packages.json)?
```
$ sudo npm install -g npm-check-updates
$ npm-check-updates -u -a
```

## How to add configmap to openshift?
```
$ oc create -f .openshift/configmap.yaml
```

## How to update s2i builder image in openshift?
```
$ oc tag --scheduled --source=docker $DOCKER_ID_USER/s2i-cheert-centos7:latest s2i-cheert-centos7:latest
```
