#!/bin/bash

# Cleanup function
cleanup() {
    echo "Cleaning up processes..."
    pkill -P $$ # Kill all child processes
    killall -9 air main templ tailwindcss 2>/dev/null || true
    exit 0
}

# Trap Ctrl+C and run cleanup
trap cleanup SIGINT SIGTERM

# Start all watchers in background
task watch:templ-local &
task watch:templ-shared &
task watch:tailwind &
task run:air &

# Wait for all background jobs
wait
