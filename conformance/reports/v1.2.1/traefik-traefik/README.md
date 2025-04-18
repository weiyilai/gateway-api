# Traefik Proxy

## Table of Contents

| API channel  | Implementation version                                           | Mode    | Report                                            |
|--------------|------------------------------------------------------------------|---------|---------------------------------------------------|
| experimental | [v3.2.2](https://github.com/traefik/traefik/releases/tag/v3.2.2) | default | [v3.2.2 report](./experimental-v3.2.2-default-report.yaml) |

## Reproduce

To reproduce the results, clone the Traefik Proxy repository:

```shell
git clone https://github.com/traefik/traefik.git && cd traefik
```

Check out the desired version:

```shell
git checkout vX.Y
```

Run the conformance tests with:

```shell
make test-gateway-api-conformance
```

Check the produced report in the `./integration/conformance-reports` folder.
