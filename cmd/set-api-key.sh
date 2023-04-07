#!/bin/bash

if ! grep -qxF "export ANX_OPENAI_API_KEY=$1" ~/.bashrc; then
    echo "export ANX_OPENAI_API_KEY=$1" >> ~/.bashrc
fi
