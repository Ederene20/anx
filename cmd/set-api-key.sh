#!/bin/bash

if grep -qF "export ANX_OPENAI_API_KEY" ~/.bashrc; then
    sed -i '/^export ANX_OPENAI_API_KEY/d' ~/.bashrc
fi

echo "export ANX_OPENAI_API_KEY=$1" >> ~/.bashrc