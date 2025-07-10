# Mock Generation Demo using Mockery + Testify

This is a simple demo to generate Go interface mocks using [Mockery](https://github.com/vektra/mockery) and [Testify](https://github.com/stretchr/testify).

## 🧰 Requirements

- [Docker](https://www.docker.com/)

## 🚀 Commands

You can generate mocks in two ways:

### 1. Directly via Command Line

Use this method by specifying the interface name and its directory:

```bash
docker pull vektra/mockery:3.5
docker run --rm -v "%cd%:/src" -w /src vektra/mockery --name AnyRepository --dir=repository (for windows)
```

> Replace `AnyRepository` with your interface name and `repository` with the relative path to it.

---

### 2. Using `.mockery.yaml` Configuration File

Steps:

1. Pull the Docker image:

```bash
docker pull vektra/mockery:3.5
```

2. Edit the `.mockery.yaml` file (see example below).

3. Run mock generation:

```bash
docker run --rm -v "%cd%:/src" -w /src vektra/mockery (for windows)
```

#### Example `.mockery.yaml`

```yaml
name: AnyRepository
dir: repository
output: mocks
```

This configuration generates a mock for `AnyRepository` located in the `repository` directory and places the mock in the `mocks` directory.

---

## 📦 Output

Generated mocks will be placed in the directory specified by the `output` field (e.g., `mocks/`).

---

## ✅ Testing

You can now use these generated mocks with [Testify](https://github.com/stretchr/testify) in your unit tests.

---