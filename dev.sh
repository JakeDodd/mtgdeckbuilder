#!/bin/bash

if [[ ! -f .air.toml ]]; then
	air init
fi
cd react && npm install --quiet
npm run dev & cd ../ && air && fg
