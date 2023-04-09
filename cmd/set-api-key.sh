#!/bin/bash

if [ "$SHELL" = "/bin/bash" ]; then

    if grep -qF "export ANX_OPENAI_API_KEY" ~/.bashrc; then
        sed -i '/^export ANX_OPENAI_API_KEY/d' ~/.bashrc
    fi

    echo "export ANX_OPENAI_API_KEY=$1" >> ~/.bashrc
else
    case $SHELL in
    /bin/zsh) shell_profile=".zshrc" ;;
    *) shell_profile=".bash_profile" ;;
    esac

    echo "Manually add the API Key to your $HOME/$shell_profile (or similar)"
    echo " export ANX_OPENAI_API_KEY=$1 "
fi