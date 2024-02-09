[![CircleCI](https://circleci.com/gh/giantswarm/helm-chart-docs-generator/tree/main.svg?style=svg&circle-token=2847f4b99edcb9776cbd8ee622b294eb96bfd55f)](https://circleci.com/gh/giantswarm/helm-chart-docs-generator/tree/main)

# helm-chart-docs-generator

Generates configuration template for Cluster App documentation.

This tool is built to generate our [Cluster App configuration reference](https://docs.giantswarm.io/ui-api/management-api/cluster-apps/).

The generated output consists of Markdown files which includes automatically the content of the Helm chart READMEs.

## Usage

### Docker

The generator can be executed in Docker using a command like this:

```nohighlight
docker run \
    -v $PWD/path/to/output-folder:/opt/helm-chart-docs-generator/output \
    -v $PWD:/opt/helm-chart-docs-generator/config \
    quay.io/giantswarm/helm-chart-docs-generator:0.1.0 \
      --config /opt/helm-chart-docs-generator/config/config.example.yaml
```

Here, the tag `0.1.0` is the version number of the helm-chart-docs-generator release you're going to use. Check the [image repository](https://quay.io/repository/giantswarm/helm-chart-docs-generator?tab=tags) for available tags.

The volume mapping defines where the generated output will land.

### Development

With Go installed and this repository cloned, you can exetute the program like this:

```nohighlight
go run main.go --config config.example.yaml
```

See the `config.example.yaml` file for an idea how to configure your source repositories.
