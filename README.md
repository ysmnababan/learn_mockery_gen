SIMPLE DEMO TO GENERATE MOCK USING MOCKERY + TESTIFY


Requirements:
    - Docker
Command:
    1. Directly using cmd line (specify the interface name and its path)
        - docker pull vektra/mockery
        - docker run --rm -v "%cd%:/src" -w /src vektra/mockery:latest --name AnyRepository --dir=repository
    2. Using yaml file
        - docker pull vektra/mockery
        - edit the .mockery.yaml file (see the file)
        - docker run --rm -v "%cd%:/src" -w /src vektra/mockery:latest