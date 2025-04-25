#!/bin/bash
cd "client"
npm run build
cd ".."
docker-compose up --build