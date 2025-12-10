#!/bin/bash

if [[ ! -f .air.toml ]]; then
	air init
fi

cd react && npm run dev & air && fg
