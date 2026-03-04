#!/bin/bash

curl -X POST http://127.0.0.1:8080/chat \
  -H "Content-Type: application/json" \
  -d '{
    "message": "good evening",
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
