This repository contains two Docker examples demonstrating different build approaches:

1. **Multi-Stage Build**: Demonstrates how to create a smaller Docker image by separating the build environment from the runtime environment.

2. **Single-Stage Build**: Illustrates a straightforward approach where the build and runtime are combined into a single Docker image.



***

## Usage

### Multi-Stage Build
1. Navigate to the `docker-multi-stage/` directory.
2. Build the Docker image:
   ```bash
   docker build -t my-multi-stage-app .
3. Run the Docker container:
   ```bash
   docker run -it my-multi-stage-app .
   
### Single-Stage Build
1. Navigate to the single-stage/ directory.
2. Build the Docker image:
   ```bash
   docker build -t my-single-stage-app .
3. Run the Docker container:
   ```bash
   docker run -it my-single-stage-app

# Conclusion 
This repository helps you understand the differences between multi-stage and single-stage Docker builds. Each approach has its own benefits, and you can explore them through the provided examples.


