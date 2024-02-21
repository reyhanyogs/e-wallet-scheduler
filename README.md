# E-Wallet-Scheduler

An e-wallet scheduler for scheduling weekly summaries of spending and earnings to the Redis queue so it can be processed by the e-wallet-queue.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Docker

## How to Build the image
1. Build ```Dockerfile``` to generate docker image using:
   ```bash
   make docker_build
   ```
