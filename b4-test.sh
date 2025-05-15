#!/bin/bash

set -eou pipefail

test ! -e .b4.lock
touch ./.b4.lock
