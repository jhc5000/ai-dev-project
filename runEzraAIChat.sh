#!/bin/bash

curl -X POST http://localhost:8080/chat \
  -H "Content-Type: application/json" \
  -d '{
    "message": "Thats ok Ezra, Im going to rest for a bit. I worked all through the night to get you up and running, LOL. Ill talk to you later, Ezra. Welcome to the world!",
    "market": {
        "symbol": "SPY",
        "timeframe": "15m",
        "price": 0,
        "vwap": 0,
        "rsi": 0,
        "ema_20": 0,
        "ema_50": 0,
        "ema_200": 0,
        "trend": "unknown",
        "session": "after hours"
    }
  }'
